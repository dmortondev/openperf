#include "packet/analyzer/statistics/generic_flow_counters.hpp"

namespace openperf::packet::analyzer::statistics {

template <typename Key, typename Value, typename... Pairs>
constexpr auto associative_array(Pairs&&... pairs)
    -> std::array<std::pair<Key, Value>, sizeof...(pairs)>
{
    return {{std::forward<Pairs>(pairs)...}};
}

static constexpr auto flow_counter_pairs =
    associative_array<flow_counter_flags, std::string_view>(
        std::pair(flow_counter_flags::frame_count, "frame_count"),
        std::pair(flow_counter_flags::frame_length, "frame_length"),
        std::pair(flow_counter_flags::header, "header"),
        std::pair(flow_counter_flags::interarrival, "interarrival_time"),
        std::pair(flow_counter_flags::jitter_ipdv, "jitter_ipdv"),
        std::pair(flow_counter_flags::jitter_rfc, "jitter_rfc"),
        std::pair(flow_counter_flags::latency, "latency"),
        std::pair(flow_counter_flags::prbs, "prbs"),
        std::pair(flow_counter_flags::sequencing, "advanced_sequencing"));

enum flow_counter_flags to_flow_counter_flag(std::string_view name)
{
    auto cursor = std::begin(flow_counter_pairs),
         end = std::end(flow_counter_pairs);
    while (cursor != end) {
        if (cursor->second == name) return (cursor->first);
        cursor++;
    }

    return (flow_counter_flags::none);
}

std::string_view to_name(enum flow_counter_flags flag)
{
    auto cursor = std::begin(flow_counter_pairs),
         end = std::end(flow_counter_pairs);
    while (cursor != end) {
        if (cursor->first == flag) return (cursor->second);
        cursor++;
    }

    return ("none");
}

} // namespace openperf::packet::analyzer::statistics
