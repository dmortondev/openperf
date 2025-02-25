import sys

from datetime import datetime
from expects import expect, be_a, be_empty, be_none, equal
from expects.matchers import Matcher

import client.models


class _be_valid_counter(Matcher):
    def _match(self, counter):
        expect(counter).to(be_a(client.models.TimeCounter))
        expect(counter.id).not_to(be_empty)
        expect(counter.name).not_to(be_empty)
        expect(counter.frequency).not_to(be_none)
        return True, ['is valid counter']


class _be_valid_keeper(Matcher):
    def _match(self, keeper):
        expect(keeper).to(be_a(client.models.TimeKeeper))
        expect(keeper.time).to(be_a(datetime))
        expect(keeper.time_counter_id).not_to(be_empty)
        expect(keeper.stats).to(be_a(client.models.TimeKeeperStats))
        expect(keeper.stats.round_trip_times).to(be_a(client.models.TimeKeeperStatsRoundTripTimes))
        expect(keeper.state).to(be_a(client.models.TimeKeeperState))
        if (keeper.state.synced):
            expect(keeper.time_source_id).not_to(be_empty)
        return True, ['is valid keeper']


class _be_valid_source_ntp_stats(Matcher):
    def _match(self, stats):
        expect(stats).to(be_a(client.models.TimeSourceStatsNtp))
        expect(stats.poll_period).not_to(be_none)
        expect(stats.rx_ignored).not_to(be_none)
        expect(stats.rx_packets).not_to(be_none)
        expect(stats.tx_packets).not_to(be_none)
        if stats.rx_packets > 0:
            expect(stats.last_rx_accepted).to(be_a(datetime))
            expect(stats.stratum).to(be_above_or_equal(1))
        if stats.rx_ignored > 0:
            expect(stats.last_rx_ingored).to(be_a(datetime))
        return True, ['is valid source ntp stats']


class _be_valid_source(Matcher):
    def _match(self, source):
        expect(source).to(be_a(client.models.TimeSource))
        expect(source.id).not_to(be_empty)
        expect(source.stats).to(be_a(client.models.TimeSourceStats))
        if source.kind == 'ntp':
            expect(source.config).to(be_a(client.models.TimeSourceConfig))
            ntp = source.config.ntp
            expect(ntp).to(be_a(client.models.TimeSourceConfigNtp))
            expect(ntp.hostname).not_to(be_empty)
            expect(source.stats.ntp).to(be_valid_source_ntp_stats)
            return True, ['is valid ntp source']
        elif source.kind == 'system':
            expect(source.stats.system).to(be_a(client.models.TimeSourceStatsSystem))
            return True, ['is valid system source']
        else:
            return False, ['is invalid source']


be_valid_counter = _be_valid_counter()
be_valid_keeper = _be_valid_keeper()
be_valid_source_ntp_stats = _be_valid_source_ntp_stats()
be_valid_source = _be_valid_source()
