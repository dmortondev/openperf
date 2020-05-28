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


#include "AnalyzerResult.h"

namespace swagger {
namespace v1 {
namespace model {

AnalyzerResult::AnalyzerResult()
{
    m_Id = "";
    m_Analyzer_id = "";
    m_Analyzer_idIsSet = false;
    m_Active = false;
    m_FlowsIsSet = false;
    
}

AnalyzerResult::~AnalyzerResult()
{
}

void AnalyzerResult::validate()
{
    // TODO: implement validation
}

nlohmann::json AnalyzerResult::toJson() const
{
    nlohmann::json val = nlohmann::json::object();

    val["id"] = ModelBase::toJson(m_Id);
    if(m_Analyzer_idIsSet)
    {
        val["analyzer_id"] = ModelBase::toJson(m_Analyzer_id);
    }
    val["active"] = m_Active;
    val["protocol_counters"] = ModelBase::toJson(m_Protocol_counters);
    val["flow_counters"] = ModelBase::toJson(m_Flow_counters);
    {
        nlohmann::json jsonArray;
        for( auto& item : m_Flows )
        {
            jsonArray.push_back(ModelBase::toJson(item));
        }
        
        if(jsonArray.size() > 0)
        {
            val["flows"] = jsonArray;
        }
    }
    

    return val;
}

void AnalyzerResult::fromJson(nlohmann::json& val)
{
    setId(val.at("id"));
    if(val.find("analyzer_id") != val.end())
    {
        setAnalyzerId(val.at("analyzer_id"));
        
    }
    setActive(val.at("active"));
    {
        m_Flows.clear();
        nlohmann::json jsonArray;
        if(val.find("flows") != val.end())
        {
        for( auto& item : val["flows"] )
        {
            m_Flows.push_back(item);
            
        }
        }
    }
    
}


std::string AnalyzerResult::getId() const
{
    return m_Id;
}
void AnalyzerResult::setId(std::string value)
{
    m_Id = value;
    
}
std::string AnalyzerResult::getAnalyzerId() const
{
    return m_Analyzer_id;
}
void AnalyzerResult::setAnalyzerId(std::string value)
{
    m_Analyzer_id = value;
    m_Analyzer_idIsSet = true;
}
bool AnalyzerResult::analyzerIdIsSet() const
{
    return m_Analyzer_idIsSet;
}
void AnalyzerResult::unsetAnalyzer_id()
{
    m_Analyzer_idIsSet = false;
}
bool AnalyzerResult::isActive() const
{
    return m_Active;
}
void AnalyzerResult::setActive(bool value)
{
    m_Active = value;
    
}
std::shared_ptr<AnalyzerProtocolCounters> AnalyzerResult::getProtocolCounters() const
{
    return m_Protocol_counters;
}
void AnalyzerResult::setProtocolCounters(std::shared_ptr<AnalyzerProtocolCounters> value)
{
    m_Protocol_counters = value;
    
}
std::shared_ptr<AnalyzerFlowCounters> AnalyzerResult::getFlowCounters() const
{
    return m_Flow_counters;
}
void AnalyzerResult::setFlowCounters(std::shared_ptr<AnalyzerFlowCounters> value)
{
    m_Flow_counters = value;
    
}
std::vector<std::string>& AnalyzerResult::getFlows()
{
    return m_Flows;
}
bool AnalyzerResult::flowsIsSet() const
{
    return m_FlowsIsSet;
}
void AnalyzerResult::unsetFlows()
{
    m_FlowsIsSet = false;
}

}
}
}

