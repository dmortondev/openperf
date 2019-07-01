#include <zmq.h>

#include "json.hpp"

#include "api/api_route_handler.h"
#include "core/icp_core.h"
#include "packetio/stack_api.h"
#include "config/icp_config_utils.h"

namespace icp {
namespace packetio {
namespace stack {
namespace api {

using namespace Pistache;
using json = nlohmann::json;

json submit_request(void *socket, json& request)
{
    auto type = request["type"].get<request_type>();
    switch (type) {
    case request_type::GET_STACK:
        ICP_LOG(ICP_LOG_TRACE, "Sending %s request for stack %s\n",
                to_string(type).c_str(),
                request["id"].get<std::string>().c_str());
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
    handler(void* context, Rest::Router& router);

    void list_stacks(const Rest::Request& request, Http::ResponseWriter response);
    void   get_stack(const Rest::Request& request, Http::ResponseWriter response);

private:
    std::unique_ptr<void, icp_socket_deleter> m_socket;
};

handler::handler(void* context, Rest::Router& router)
    : m_socket(icp_socket_get_client(context, ZMQ_REQ, endpoint.c_str()))
{
    Rest::Routes::Get(router, "/stacks",
                      Rest::Routes::bind(&handler::list_stacks, this));
    Rest::Routes::Get(router, "/stacks/:id",
                      Rest::Routes::bind(&handler::get_stack, this));
}


void handler::list_stacks(const Rest::Request& requst __attribute__((unused)),
                          Http::ResponseWriter response)
{
    json api_request = {
        {"type", request_type::LIST_STACKS }
    };

    json api_reply = submit_request(m_socket.get(), api_request);

    response.headers().add<Http::Header::ContentType>(MIME(Application, Json));

    if (api_reply["code"].get<reply_code>() == reply_code::OK) {
        response.send(Http::Code::Ok,
                      api_reply["data"].get<std::string>());
    } else {
        response.send(Http::Code::Internal_Server_Error,
                      api_reply["error"].get<std::string>());
    }
}

/**
 * Our id might not be a valid string. This macro handles validating the id
 * value using the framework. If the id is invalid this macro will send a reply
 * using the supplied code and return. Unfortunately, we cannot make this a function
 * because the response is an unmovable object.
 */

#define VALIDATE_ID(id_, response_, code_)                     \
    do {                                                       \
        auto res = config::icp_config_validate_id_string(id_); \
        if (!res) {                                            \
            response_.send(code_, res.error());                \
            return;                                            \
        }                                                      \
    } while (0)

void handler::get_stack(const Rest::Request& request,
                        Http::ResponseWriter response)
{
    auto id = request.param(":id").as<std::string>();
    VALIDATE_ID(id, response, Http::Code::Not_Found);

    json api_request = {
        { "type", request_type::GET_STACK },
        { "id", id }
    };

    json api_reply = submit_request(m_socket.get(), api_request);

    switch (api_reply["code"].get<reply_code>()) {
    case reply_code::OK:
        response.headers().add<Http::Header::ContentType>(MIME(Application, Json));
        response.send(Http::Code::Ok,
                      api_reply["data"].get<std::string>());
        break;
    case reply_code::NO_STACK:
        response.send(Http::Code::Not_Found);
        break;
    default:
        response.headers().add<Http::Header::ContentType>(MIME(Application, Json));
        response.send(Http::Code::Internal_Server_Error,
                      api_reply["error"].get<std::string>());
     }
}

}
}
}
}
