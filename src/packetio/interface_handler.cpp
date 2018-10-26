#include <zmq.h>

#include "api/api_route_handler.h"
#include "core/icp_core.h"
#include "packetio/interface_api.h"

namespace icp {
namespace packetio {
namespace interface {
namespace api {

using namespace Pistache;
using json = nlohmann::json;

json submit_request(void *socket, json& request)
{
    auto type = request["type"].get<request_type>();
    switch (type) {
    case request_type::GET_INTERFACE:
    case request_type::DELETE_INTERFACE:
        ICP_LOG(ICP_LOG_TRACE, "Sending %s request for interface %d\n",
                to_string(type).c_str(),
                request["id"].get<int>());
        break;
    default:
        ICP_LOG(ICP_LOG_TRACE, "Sending %s request\n", to_string(type).c_str());
    }

    std::vector<uint8_t> request_buffer = json::to_cbor(request);
    zmq_msg_t reply_msg;
    if (zmq_msg_init(&reply_msg) == -1
        || zmq_send(socket, request_buffer.data(), request_buffer.size(), 0) != static_cast<int>(request_buffer.size())
        || zmq_msg_recv(&reply_msg, socket, 0) == -1) {
        return {
            { "code", reply_code::ERROR },
            { "error", json_error(errno, zmq_strerror(errno)) }
        };
    }

    ICP_LOG(ICP_LOG_TRACE, "Received %s reply\n", to_string(type).c_str());

    std::vector<uint8_t> reply_buffer(static_cast<uint8_t *>(zmq_msg_data(&reply_msg)),
                                      static_cast<uint8_t *>(zmq_msg_data(&reply_msg)) + zmq_msg_size(&reply_msg));

    zmq_msg_close(&reply_msg);

    return json::from_cbor(reply_buffer);
}

class handler : public icp::api::route::handler::registrar<handler> {
public:
    handler(void *context, Rest::Router &router);

    void  list_interfaces(const Rest::Request &request, Http::ResponseWriter response);
    void create_interface(const Rest::Request &request, Http::ResponseWriter response);
    void    get_interface(const Rest::Request &request, Http::ResponseWriter response);
    void delete_interface(const Rest::Request &request, Http::ResponseWriter response);

    /* Bulk interface options */
    void bulk_create_interfaces(const Rest::Request &request, Http::ResponseWriter response);
    void bulk_delete_interfaces(const Rest::Request &request, Http::ResponseWriter response);

private:
    std::unique_ptr<void, icp_socket_deleter> socket;
};

handler::handler(void *context, Rest::Router &router)
    : socket(icp_socket_get_client(context, ZMQ_REQ, endpoint.c_str()))
{
    Rest::Routes::Get(router, "/interfaces",
                      Rest::Routes::bind(&handler::list_interfaces, this));
    Rest::Routes::Post(router, "/interfaces",
                       Rest::Routes::bind(&handler::create_interface, this));
    Rest::Routes::Get(router, "/interfaces/:id",
                      Rest::Routes::bind(&handler::get_interface, this));
    Rest::Routes::Delete(router, "/interfaces/:id",
                         Rest::Routes::bind(&handler::delete_interface, this));

    Rest::Routes::Post(router, "interfaces/x/bulk-create",
                       Rest::Routes::bind(&handler::bulk_create_interfaces, this));
    Rest::Routes::Post(router, "interfaces/x/bulk-delete",
                       Rest::Routes::bind(&handler::bulk_delete_interfaces, this));
}

void handler::list_interfaces(const Rest::Request &request,
                              Http::ResponseWriter response)
{
    json api_request = {
        {"type", request_type::LIST_INTERFACES }
    };

    /* Check for supported parameters */
    for (auto &s : { "port_id", "eth_mac_address", "ipv4_address" } ) {
        if (request.query().has(s)) {
            api_request[s] = request.query().get(s).get();
        }
    }

    json api_reply = submit_request(socket.get(), api_request);

    response.headers().add<Http::Header::ContentType>(MIME(Application, Json));
    if (api_reply["code"].get<reply_code>() == reply_code::OK) {
        response.send(Http::Code::Ok, api_reply["data"].get<std::string>());
    } else {
        response.send(Http::Code::Internal_Server_Error,
                      api_reply["error"].get<std::string>());
    }
}

void handler::create_interface(const Rest::Request &request,
                               Http::ResponseWriter response)
{
    json api_request = {
        { "type", request_type::CREATE_INTERFACE },
        { "data", request.body() }
    };

    json api_reply = submit_request(socket.get(), api_request);

    response.headers().add<Http::Header::ContentType>(MIME(Application, Json));
    switch (api_reply["code"].get<reply_code>()) {
    case reply_code::OK:
        response.send(Http::Code::Ok,
                      api_reply["data"].get<std::string>());
        break;
    case reply_code::BAD_INPUT:
        response.send(Http::Code::Bad_Request,
                      api_reply["error"].get<std::string>());
        break;
    default:
        response.send(Http::Code::Internal_Server_Error,
                      api_reply["error"].get<std::string>());
    }
}

/**
 * Our id might not be an int.  In that case, the Pistache request will throw
 * an exception.  This macro handles catching the exception and returning
 * the appropriate response.  Unfortunately, we cannot make this a function
 * because the response is an unmovable object.
 */

#define SAFE_GET_ID(id_, response_, code_)                              \
    do {                                                                \
        try {                                                           \
            id_ = request.param(":id").as<int>();                       \
        } catch (const std::runtime_error&) {                           \
            response_.send(code_);                                      \
            return;                                                     \
        }                                                               \
    } while (0)

void handler::get_interface(const Rest::Request &request,
                            Http::ResponseWriter response)
{
    int id = -1;
    SAFE_GET_ID(id, response, Http::Code::Not_Found);

    json api_request = {
        { "type", request_type::GET_INTERFACE },
        { "id", id }
    };

    json api_reply = submit_request(socket.get(), api_request);

    switch (api_reply["code"].get<reply_code>()) {
    case reply_code::OK:
        response.headers().add<Http::Header::ContentType>(MIME(Application, Json));
        response.send(Http::Code::Ok,
                      api_reply["data"].get<std::string>());
        break;
    case reply_code::NO_INTERFACE:
        response.send(Http::Code::Not_Found);
        break;
    default:
        response.headers().add<Http::Header::ContentType>(MIME(Application, Json));
        response.send(Http::Code::Internal_Server_Error,
                      api_reply["error"].get<std::string>());
    }
}

void handler::delete_interface(const Rest::Request &request,
                               Http::ResponseWriter response)
{
    int id = -1;
    SAFE_GET_ID(id, response, Http::Code::No_Content);

    json api_request = {
        { "type", request_type::DELETE_INTERFACE },
        { "id", id }
    };

    /* We don't care about any reply, here */
    submit_request(socket.get(), api_request);
    response.send(Http::Code::No_Content);
}

void handler::bulk_create_interfaces(const Rest::Request &request,
                                     Http::ResponseWriter response)
{
    response.headers().add<Http::Header::ContentType>(MIME(Application, Json));

    try {
        json body = json::parse(request.body());

        if (body.find("items") != body.end()) {
            json api_request = {
                { "type", request_type::BULK_CREATE_INTERFACES },
                { "items", body["items"] }
            };

            json api_reply = submit_request(socket.get(), api_request);

            switch (api_reply["code"].get<reply_code>()) {
            case reply_code::OK:
                response.send(Http::Code::Ok,
                              api_reply["data"].get<std::string>());
                break;
            case reply_code::BAD_INPUT:
                response.send(Http::Code::Bad_Request,
                      api_reply["error"].get<std::string>());
                break;
            default:
                response.send(Http::Code::Internal_Server_Error,
                              api_reply["error"].get<std::string>());
            }
        }
    } catch (const json::parse_error &e) {
        response.send(Http::Code::Bad_Request, json_error(e.id, e.what()));
    }
}

void handler::bulk_delete_interfaces(const Rest::Request &request,
                                     Http::ResponseWriter response)
{
    try {
        json body = json::parse(request.body());

        if (body.find("ids") != body.end()) {
            json api_request = {
                { "type", request_type::BULK_DELETE_INTERFACES },
                { "ids", body["ids"] }
            };
            submit_request(socket.get(), api_request);
        }
        response.send(Http::Code::No_Content);
    } catch (const json::parse_error &e) {
        response.headers().add<Http::Header::ContentType>(MIME(Application, Json));
        response.send(Http::Code::Bad_Request, json_error(e.id, e.what()));
    }
}

}
}
}
}
