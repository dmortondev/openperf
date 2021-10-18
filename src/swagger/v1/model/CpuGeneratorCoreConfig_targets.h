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
 * CpuGeneratorCoreConfig_targets.h
 *
 * 
 */

#ifndef CpuGeneratorCoreConfig_targets_H_
#define CpuGeneratorCoreConfig_targets_H_


#include "ModelBase.h"

#include "CpuGeneratorCoreLoad.h"

namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// 
/// </summary>
class  CpuGeneratorCoreConfig_targets
    : public ModelBase
{
public:
    CpuGeneratorCoreConfig_targets();
    virtual ~CpuGeneratorCoreConfig_targets();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// CpuGeneratorCoreConfig_targets members

    /// <summary>
    /// 
    /// </summary>
    std::shared_ptr<CpuGeneratorCoreLoad> getLoad() const;
    void setLoad(std::shared_ptr<CpuGeneratorCoreLoad> value);
        /// <summary>
    /// Targeted load ratio
    /// </summary>
    int32_t getWeight() const;
    void setWeight(int32_t value);
    
protected:
    std::shared_ptr<CpuGeneratorCoreLoad> m_Load;

    int32_t m_Weight;

};

}
}
}

#endif /* CpuGeneratorCoreConfig_targets_H_ */
