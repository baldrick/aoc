package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Day24 {
    private static final Logger log = LoggerFactory.getLogger(Day24.class);
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day24.txt");
        final Day24 day = new Day24(input);
        //log.info("Part 1 = {}", day.part1());
        log.info("Part 2 = {}", day.part2());
    }

    private final List<String> input;
    private boolean prepared;
    private List<Integer> weights;
    private Bridge bridges;

    public Day24(final List<String> input) {
        this.input = input;
        this.prepared = false;
    }

    static class Component {
        private final int in;
        private final int out;

        Component(final int in, final int out) {
            this.in = in;
            this.out = out;
        }

        Component reverse() {
            return new Component(out, in);
        }
    }

    static class Bridge {
        private final Component start;
        private final List<Bridge> options;

        Bridge(final Component start) {
            this.start = start;
            options = new ArrayList<>();
        }
    }

    private void prepare() {
        if (!prepared) {
            final List<Component> components = input.stream().map(Day24::getComponent).collect(Collectors.toList());
            final Map<Integer, List<Component>> componentMap = new HashMap<>();
            final Map<Integer, List<Component>> reverseComponentMap = new HashMap<>();
            for (Component component : components) {
                addComponentToMap(componentMap, reverseComponentMap, component);
            }
            bridges = createBridges(new Bridge(null), componentMap, reverseComponentMap);
            //dumpBridges(bridges, 0);
        }
    }

    public int part1() {
        prepare();
        weights = new ArrayList<>();
        calculateWeights(bridges, weights, 0);
        return weights.stream().max(Integer::compare).get();
    }

    private void dumpBridges(final Bridge bridges, final int level) {
        if (null != bridges.start) {
            log.info("{}{}/{}", Utils.repeat(" ", level), bridges.start.in, bridges.start.out);
        }
        for (Bridge bridge : bridges.options) {
            dumpBridges(bridge, level + 1);
        }
    }

    private void calculateWeights(final Bridge bridges, final List<Integer> weights, final int weight) {
        final int thisWeight;
        if (null != bridges.start) {
            thisWeight = bridges.start.in + bridges.start.out;
        } else {
            thisWeight = 0;
        }
        if (bridges.options.isEmpty()) {
            weights.add(weight + thisWeight);
        } else {
            for (Bridge bridge : bridges.options) {
                calculateWeights(bridge, weights, weight + thisWeight);
            }
        }
    }

    private void addComponentToMap(
        final Map<Integer, List<Component>> componentMap,
        final Map<Integer, List<Component>> reverseComponentMap,
        final Component component) {
        if (componentMap.containsKey(component.in)) {
            componentMap.get(component.in).add(component);
        } else {
            final List<Component> outList = new ArrayList<>();
            outList.add(component);
            componentMap.put(component.in, outList);
        }

        if (reverseComponentMap.containsKey(component.out)) {
            reverseComponentMap.get(component.out).add(component);
        } else {
            final List<Component> outList = new ArrayList<>();
            outList.add(component);
            reverseComponentMap.put(component.out, outList);
        }
    }

    public Bridge createBridges(
        final Bridge bridge,
        final Map<Integer, List<Component>> componentMap,
        final Map<Integer, List<Component>> reverseComponentMap) {
        final int findComponent;
        if (null == bridge.start) {
            findComponent = 0;
        } else {
            findComponent = bridge.start.out;
        }
        if (componentMap.containsKey(findComponent)) {
            log.debug("Adding {} options to in={}", componentMap.get(findComponent).size(), findComponent);
            for (Component out : componentMap.get(findComponent)) {
                bridge.options.add(
                    createBridges(new Bridge(out), componentMapCopy(componentMap, out.in, out), componentMapCopy(reverseComponentMap, out.out, out)));
            }
        }
        if (reverseComponentMap.containsKey(findComponent)) {
            log.debug("Adding {} options to in={}", reverseComponentMap.get(findComponent).size(), findComponent);
            for (Component out : reverseComponentMap.get(findComponent)) {
                bridge.options.add(
                    createBridges(new Bridge(out.reverse()), componentMapCopy(componentMap, out.in, out), componentMapCopy(reverseComponentMap, out.out, out)));
            }
        }
        return bridge;
    }

    private Map<Integer, List<Component>> componentMapCopy(final Map<Integer, List<Component>> source, final int port, final Component omit) {
        final Map<Integer, List<Component>> copy = new HashMap<>();
        for (Map.Entry<Integer, List<Component>> entry : source.entrySet()) {
            final List<Component> components = new ArrayList<>();
            components.addAll(entry.getValue());
            copy.put(entry.getKey(), components);
        }
        final List<Component> components = copy.get(port);
        components.remove(omit);
        return copy;
    }

    private static Component getComponent(final String row) {
        return new Component(getIn(row), getOut(row));
    }

    private static int getIn(final String row) {
        final int slash = row.indexOf('/');
        return Integer.valueOf(row.substring(0, slash));
    }

    private static int getOut(final String row) {
        final int slash = row.indexOf('/');
        return Integer.valueOf(row.substring(slash + 1));
    }

    class BridgeInfo {
        private int length;
        private int weight;

        BridgeInfo(final int length, final int weight) {
            this.length = length;
            this.weight = weight;
        }

        BridgeInfo(final BridgeInfo bi) {
            this.length = bi.length;
            this.weight = bi.weight;
        }

        void reset() {
            length = 0;
            weight = 0;
        }

        void add(final Component component) {
            length++;
            weight += component.in + component.out;
        }

        void copyIfLonger(final BridgeInfo bi) {
            if (bi.length > length) {
                length = bi.length;
                weight = bi.weight;
            } else if (bi.length == length) {
                if (bi.weight > weight) {
                    weight = bi.weight;
                }
            }
        }
    }

    private void findLongestHeaviestBridge(final Bridge bridges, BridgeInfo longest, BridgeInfo current) {
        if (null != bridges.start) {
            current.add(bridges.start);
        }
        if (bridges.options.isEmpty()) {
            // bridge is complete - is it the longest so far?
            longest.copyIfLonger(current);
            current.reset();
        } else {
            for (Bridge bridge : bridges.options) {
                findLongestHeaviestBridge(bridge, longest, new BridgeInfo(current));
            }
        }
    }

    public int part2() {
        prepare();
        final BridgeInfo longest = new BridgeInfo(0, 0);
        findLongestHeaviestBridge(bridges, longest, new BridgeInfo(0, 0));
        return longest.weight;
    }
}
