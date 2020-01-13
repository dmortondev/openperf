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


#include "TimeKeeper.h"

namespace swagger {
namespace v1 {
namespace model {

TimeKeeper::TimeKeeper()
{
    m_Time = "";
    m_Time_counter_id = "";
    m_Time_source_id = "";
    
}

TimeKeeper::~TimeKeeper()
{
}

void TimeKeeper::validate()
{
    // TODO: implement validation
}

nlohmann::json TimeKeeper::toJson() const
{
    nlohmann::json val = nlohmann::json::object();

    val["state"] = ModelBase::toJson(m_State);
    val["stats"] = ModelBase::toJson(m_Stats);
    val["time"] = ModelBase::toJson(m_Time);
    val["time_counter_id"] = ModelBase::toJson(m_Time_counter_id);
    val["time_source_id"] = ModelBase::toJson(m_Time_source_id);
    

    return val;
}

void TimeKeeper::fromJson(nlohmann::json& val)
{
    setTime(val.at("time"));
    setTimeCounterId(val.at("time_counter_id"));
    setTimeSourceId(val.at("time_source_id"));
    
}


std::shared_ptr<TimeKeeperState> TimeKeeper::getState() const
{
    return m_State;
}
void TimeKeeper::setState(std::shared_ptr<TimeKeeperState> value)
{
    m_State = value;
    
}
std::shared_ptr<TimeKeeperStats> TimeKeeper::getStats() const
{
    return m_Stats;
}
void TimeKeeper::setStats(std::shared_ptr<TimeKeeperStats> value)
{
    m_Stats = value;
    
}
std::string TimeKeeper::getTime() const
{
    return m_Time;
}
void TimeKeeper::setTime(std::string value)
{
    m_Time = value;
    
}
std::string TimeKeeper::getTimeCounterId() const
{
    return m_Time_counter_id;
}
void TimeKeeper::setTimeCounterId(std::string value)
{
    m_Time_counter_id = value;
    
}
std::string TimeKeeper::getTimeSourceId() const
{
    return m_Time_source_id;
}
void TimeKeeper::setTimeSourceId(std::string value)
{
    m_Time_source_id = value;
    
}

}
}
}

