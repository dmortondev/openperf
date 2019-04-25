#ifndef _ICP_API_INTERNAL_CLIENT_H_
#define _ICP_API_INTERNAL_CLIENT_H_

namespace icp::api::client {

// Send a GET request to the REST API.
// @param resource - REST API resource perform GET operation on.
// @returns Pair of Http reponse code and response body.
std::pair<Pistache::Http::Code, std::string> internal_api_get(std::string_view resource);

// Sent a POST request to the REST API.
// @param resource - REST API resource to perform POST operation on.
// @param body - Request body to include in the POST request.
// @returns Pair of Http reponse code and response body.
std::pair<Pistache::Http::Code, std::string> internal_api_post(std::string_view resource,
                                                               const std::string &body);

}  // namespace icp::api::client
#endif
