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
 * PortConfig.h
 *
 * Port configuration
 */

#ifndef PortConfig_H_
#define PortConfig_H_


#include "ModelBase.h"

#include "PortConfig_bond.h"
#include "PortConfig_dpdk.h"

namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// Port configuration
/// </summary>
class  PortConfig
    : public ModelBase
{
public:
    PortConfig();
    virtual ~PortConfig();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// PortConfig members

    /// <summary>
    /// 
    /// </summary>
    std::shared_ptr<PortConfig_bond> getBond() const;
    void setBond(std::shared_ptr<PortConfig_bond> value);
    bool bondIsSet() const;
    void unsetBond();
    /// <summary>
    /// 
    /// </summary>
    std::shared_ptr<PortConfig_dpdk> getDpdk() const;
    void setDpdk(std::shared_ptr<PortConfig_dpdk> value);
    bool dpdkIsSet() const;
    void unsetDpdk();

protected:
    std::shared_ptr<PortConfig_bond> m_Bond;
    bool m_BondIsSet;
    std::shared_ptr<PortConfig_dpdk> m_Dpdk;
    bool m_DpdkIsSet;
};

}
}
}

#endif /* PortConfig_H_ */
