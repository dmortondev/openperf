/**
* OpenPerf API
* REST API interface for OpenPerf
*
* OpenAPI spec version: 1
* Contact: support@spirent.com
*
* NOTE: This class is auto generated by the swagger code generator program.
* https://github.com/swagger-api/swagger-codegen.git
* Do not edit the class manually.
*/
/*
 * TvlpProfile_packet.h
 *
 * 
 */

#ifndef TvlpProfile_packet_H_
#define TvlpProfile_packet_H_


#include "ModelBase.h"

#include "TvlpProfile_packet_series.h"
#include <vector>

namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// 
/// </summary>
class  TvlpProfile_packet
    : public ModelBase
{
public:
    TvlpProfile_packet();
    virtual ~TvlpProfile_packet();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// TvlpProfile_packet members

    /// <summary>
    /// 
    /// </summary>
    std::vector<std::shared_ptr<TvlpProfile_packet_series>>& getSeries();
    
protected:
    std::vector<std::shared_ptr<TvlpProfile_packet_series>> m_Series;

};

}
}
}

#endif /* TvlpProfile_packet_H_ */
