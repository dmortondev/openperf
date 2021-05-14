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
 * PacketAnalyzerResult.h
 *
 * Results collected by a running packet analyzer
 */

#ifndef PacketAnalyzerResult_H_
#define PacketAnalyzerResult_H_


#include "ModelBase.h"

#include "PacketAnalyzerFlowDigests.h"
#include <string>
#include "PacketAnalyzerFlowCounters.h"
#include <vector>
#include "PacketAnalyzerProtocolCounters.h"

namespace swagger {
namespace v1 {
namespace model {

/// <summary>
/// Results collected by a running packet analyzer
/// </summary>
class  PacketAnalyzerResult
    : public ModelBase
{
public:
    PacketAnalyzerResult();
    virtual ~PacketAnalyzerResult();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    nlohmann::json toJson() const override;
    void fromJson(nlohmann::json& json) override;

    /////////////////////////////////////////////
    /// PacketAnalyzerResult members

    /// <summary>
    /// Unique analyzer result identifier
    /// </summary>
    std::string getId() const;
    void setId(std::string value);
        /// <summary>
    /// Unique analyzer identifier that generated this result
    /// </summary>
    std::string getAnalyzerId() const;
    void setAnalyzerId(std::string value);
    bool analyzerIdIsSet() const;
    void unsetAnalyzer_id();
    /// <summary>
    /// Indicates whether the result is currently receiving updates
    /// </summary>
    bool isActive() const;
    void setActive(bool value);
        /// <summary>
    /// 
    /// </summary>
    std::shared_ptr<PacketAnalyzerProtocolCounters> getProtocolCounters() const;
    void setProtocolCounters(std::shared_ptr<PacketAnalyzerProtocolCounters> value);
        /// <summary>
    /// 
    /// </summary>
    std::shared_ptr<PacketAnalyzerFlowCounters> getFlowCounters() const;
    void setFlowCounters(std::shared_ptr<PacketAnalyzerFlowCounters> value);
        /// <summary>
    /// 
    /// </summary>
    std::shared_ptr<PacketAnalyzerFlowDigests> getFlowDigests() const;
    void setFlowDigests(std::shared_ptr<PacketAnalyzerFlowDigests> value);
    bool flowDigestsIsSet() const;
    void unsetFlow_digests();
    /// <summary>
    /// List of unique flow ids included in stats. Individual flow statistics may be queried via the &#x60;rx-flows&#x60; endpoint. 
    /// </summary>
    std::vector<std::string>& getFlows();
    bool flowsIsSet() const;
    void unsetFlows();

protected:
    std::string m_Id;

    std::string m_Analyzer_id;
    bool m_Analyzer_idIsSet;
    bool m_Active;

    std::shared_ptr<PacketAnalyzerProtocolCounters> m_Protocol_counters;

    std::shared_ptr<PacketAnalyzerFlowCounters> m_Flow_counters;

    std::shared_ptr<PacketAnalyzerFlowDigests> m_Flow_digests;
    bool m_Flow_digestsIsSet;
    std::vector<std::string> m_Flows;
    bool m_FlowsIsSet;
};

}
}
}

#endif /* PacketAnalyzerResult_H_ */
