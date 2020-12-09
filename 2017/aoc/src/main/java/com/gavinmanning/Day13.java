package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Day13 {
    private static final Logger log = LoggerFactory.getLogger(Day13.class);
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day13.txt");
        final Day13 day = new Day13(input);
        log.info("Part 1 = {}", day.part1());
        log.info("Part 2 = {}", day.part2(1, 5000000));
    }

    private final List<String> input;

    public Day13(final List<String> input) {
        this.input = input;
    }

    static class Firewall {
        private final int range;
        private int pos;
        private int direction;

        Firewall(final int range) {
            this.range = range;
            this.pos = 0;
            this.direction = -1;
        }

        void move() {
            if ((0 == pos) || (pos == range - 1)) {
                direction = -direction;
            }
            pos += direction;
        }

        boolean caught(final int ps) {
            return ps % ((range - 1) * 2) == 0;
        }
    }
/*
    .       0
    ..      2
    ...     4
    ....    6
    .....   8
    ......  10
*/
    public int part1() {
        final Map<Integer, Firewall> firewalls = input.stream().collect(Collectors.toMap(Day13::getDepth, Day13::createFirewall));
        return firePacket(firewalls);
    }

    private int firePacket(final Map<Integer, Firewall> firewalls) {
        int caught = 0;
        final int lastFirewall = firewalls.keySet().stream().max(Integer::compare).get();
        for (int i = 0;  i <= lastFirewall;  i++) {
            final Firewall firewall = firewalls.get(i);
            if (null != firewall) {
                if (firewall.pos == 0) {
                    log.debug("Caught at {}", i);
                    caught += i * firewall.range;
                }
            } else {
                log.trace("No firewall at {}", i);
            }
            firewalls.values().forEach(Firewall::move);
        }
        return caught;
    }

    private static int getDepth(final String firewall) {
        final int colon = firewall.indexOf(':');
        return Integer.valueOf(firewall.substring(0, colon));
    }

    private static Firewall createFirewall(final String firewall) {
        final int colon = firewall.indexOf(':');
        return new Firewall(Integer.valueOf(firewall.substring(colon + 1).trim()));
    }

    public int part2(final int startAt, final int stopAfter) {
        int delay = startAt;
        while (true) {
            if (delay % 10000 == 0) {
                log.info("Trying delay {}ps onwards", delay);
            }
            final Map<Integer, Firewall> firewalls = input.stream().collect(Collectors.toMap(Day13::getDepth, Day13::createFirewall));
//            for (int i = 0; i < delay; i++) {
//                firewalls.values().forEach(Firewall::move);
//            }
            //dumpFirewalls("After " + delay + "ps delay", firewalls);
            if (!wouldBeCaught(firewalls, delay)) {
                log.info("Not caught after {}ps delay", delay);
                break;
            } else {
                log.debug("Caught for {}ps delay", delay);
                delay++;
            }
            if (delay > stopAfter) {
                log.warn("Failed to find solution after trying 0..{}ps delays", delay);
                return -1;
            }
        }
        return delay;
    }

    private boolean wouldBeCaught(final Map<Integer, Firewall> firewalls, final int start) {
        final int lastFirewall = firewalls.keySet().stream().max(Integer::compare).get();
        for (int ps = start;  ps <= start + lastFirewall;  ps++) {
            final Firewall firewall = firewalls.get(ps - start);
            if (null != firewall) {
                log.debug("Checking firewall {}, ps={}, range={}", ps - start, ps, firewall.range);
                if (firewall.caught(ps)) {
                    log.debug("Caught by firewall {}", ps - start);
                    return true;
                }
            } else {
                log.debug("No firewall at {}", ps - start);
            }
        }
        return false;
    }

    private void dumpFirewalls(final String message, final Map<Integer, Firewall> firewalls) {
        /*
            0   1   2   3   4   5   6
            [ ] [S] ... ... [ ] ... [ ]
            [ ] [ ]         [ ]     [ ]
            [S]             [S]     [S]
                            [ ]     [ ]
         */
        final int lastFirewall = firewalls.keySet().stream().max(Integer::compare).get();
        final StringBuilder sb = new StringBuilder();
        sb.append(message).append(System.lineSeparator());
        for (int i = 0;  i <= lastFirewall;  i++) {
            sb.append(String.format("%-3d", i)).append(" ");
        }
        sb.append(System.lineSeparator());
        for (int range = 0;  range <= 4;  range++) {
            for (int i = 0;  i <= lastFirewall;  i++) {
                final Firewall firewall = firewalls.get(i);
                if (null == firewall) {
                    sb.append("    ");
                } else if (firewall.pos == range) {
                    sb.append("[S] ");
                } else if (range < firewall.range) {
                    sb.append("[ ] ");
                } else {
                    sb.append("    ");
                }
            }
            sb.append(System.lineSeparator());
        }
        log.info(sb.toString());
    }
}
