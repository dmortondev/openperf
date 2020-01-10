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


#include "TimeKeeperStats.h"

namespace swagger {
namespace v1 {
namespace model {

TimeKeeperStats::TimeKeeperStats()
{
    m_Frequency_accept = 0.0;
    m_Frequency_reject = 0.0;
    m_Local_frequency_accept = 0.0;
    m_Local_frequency_reject = 0.0;
    m_Theta_accept = 0.0;
    m_Theta_reject = 0.0;
    m_Timestamps = 0.0;
    
}

TimeKeeperStats::~TimeKeeperStats()
{
}

void TimeKeeperStats::validate()
{
    // TODO: implement validation
}

nlohmann::json TimeKeeperStats::toJson() const
{
    nlohmann::json val = nlohmann::json::object();

    val["frequency_accept"] = m_Frequency_accept;
    val["frequency_reject"] = m_Frequency_reject;
    val["local_frequency_accept"] = m_Local_frequency_accept;
    val["local_frequency_reject"] = m_Local_frequency_reject;
    val["round_trip_times"] = ModelBase::toJson(m_Round_trip_times);
    val["theta_accept"] = m_Theta_accept;
    val["theta_reject"] = m_Theta_reject;
    val["timestamps"] = m_Timestamps;
    

    return val;
}

void TimeKeeperStats::fromJson(nlohmann::json& val)
{
    setFrequencyAccept(val.at("frequency_accept"));
    setFrequencyReject(val.at("frequency_reject"));
    setLocalFrequencyAccept(val.at("local_frequency_accept"));
    setLocalFrequencyReject(val.at("local_frequency_reject"));
    setThetaAccept(val.at("theta_accept"));
    setThetaReject(val.at("theta_reject"));
    setTimestamps(val.at("timestamps"));
    
}


double TimeKeeperStats::getFrequencyAccept() const
{
    return m_Frequency_accept;
}
void TimeKeeperStats::setFrequencyAccept(double value)
{
    m_Frequency_accept = value;
    
}
double TimeKeeperStats::getFrequencyReject() const
{
    return m_Frequency_reject;
}
void TimeKeeperStats::setFrequencyReject(double value)
{
    m_Frequency_reject = value;
    
}
double TimeKeeperStats::getLocalFrequencyAccept() const
{
    return m_Local_frequency_accept;
}
void TimeKeeperStats::setLocalFrequencyAccept(double value)
{
    m_Local_frequency_accept = value;
    
}
double TimeKeeperStats::getLocalFrequencyReject() const
{
    return m_Local_frequency_reject;
}
void TimeKeeperStats::setLocalFrequencyReject(double value)
{
    m_Local_frequency_reject = value;
    
}
std::shared_ptr<TimeKeeperStats_round_trip_times> TimeKeeperStats::getRoundTripTimes() const
{
    return m_Round_trip_times;
}
void TimeKeeperStats::setRoundTripTimes(std::shared_ptr<TimeKeeperStats_round_trip_times> value)
{
    m_Round_trip_times = value;
    
}
double TimeKeeperStats::getThetaAccept() const
{
    return m_Theta_accept;
}
void TimeKeeperStats::setThetaAccept(double value)
{
    m_Theta_accept = value;
    
}
double TimeKeeperStats::getThetaReject() const
{
    return m_Theta_reject;
}
void TimeKeeperStats::setThetaReject(double value)
{
    m_Theta_reject = value;
    
}
double TimeKeeperStats::getTimestamps() const
{
    return m_Timestamps;
}
void TimeKeeperStats::setTimestamps(double value)
{
    m_Timestamps = value;
    
}

}
}
}

