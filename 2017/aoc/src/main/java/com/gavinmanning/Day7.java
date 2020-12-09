package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

public class Day7 {
    public static void main(String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day7.txt");
        final Day7 day = new Day7(input);
        System.out.println("Part 1 = " + day.part1("hmvwl"));
        System.out.println("Part 2 = " + day.part2("hmvwl"));
    }

    private final List<String> input;
    private final Map<String, Node> nodes;

    public Day7(final List<String> input) {
        this.input = input;
        this.nodes = new HashMap<>();
    }

    private String part1(final String root) {
        nodes.putAll(input.stream().map(this::parseNode).collect(Collectors.toMap(Node::getName, Node::get)));
        final Nodes hierarchy = getHierarchy(nodes.get(root));
        return hierarchy.parent.getName();
    }

    private Nodes getHierarchy(final Node node) {
        if (null == node) {
            System.out.println("WTF node is null");
            return null;
        }
        if (null == node.children) {
            System.out.println("Adding leaf " + node.toString());
            return new Nodes(node, null);
        }
        final Set<Nodes> children = new HashSet<>();
        for (String child : node.children) {
            final Node childNode = nodes.get(child);
            System.out.println("Adding child " + child + ": " + childNode);
            children.add(getHierarchy(childNode));
        }
        return new Nodes(node, children);
    }

    private Node parseNode(final String row) {
        final int arrow = row.indexOf("->");
        if (-1 == arrow) {
            return createChild(row);
        } else {
            return createParent(row.substring(0, arrow).trim(), row.substring(arrow + 2).trim());
        }
    }

    private Node createChild(final String row) {
        return new Node(getName(row), getWeight(row));
    }

    private Node createParent(final String parent, final String children) {
        return new Node(getName(parent), getWeight(parent), Arrays.asList(children.split(",")));
    }

    private String getName(final String row) {
        final int open = row.indexOf("(");
        return row.substring(0, open - 1).trim();
    }

    private int getWeight(final String row) {
        final int open = row.indexOf("(");
        final int close = row.indexOf(")");
        final String strWeight = row.substring(open + 1, close);
        //System.out.println("Checking " + row + "; open=" + open + ",close=" + close + ", " + strWeight);
        return Integer.valueOf(strWeight);
    }

    private void dumpHierarchy(final Nodes nodes, final int level) {
        if (nodes == null) {
            return;
        }
        if (nodes.parent == null) {
            System.out.println("WTF nodes.parent is null!");
            return;
        }
        System.out.println(Utils.repeat(" ", level) + nodes.parent.getName() + " (" + nodes.parent.weight + "):");
        if (nodes.children != null) {
            for (Nodes child : nodes.children) {
                dumpHierarchy(child, level + 1);
            }
        }
    }

    private void dumpNodes() {
        for (Map.Entry<String, Node> entry : nodes.entrySet()) {
            System.out.println(entry.getKey() + " = " + entry.getValue().toString());
        }
    }

    public int part2(final String root) {
        nodes.putAll(input.stream().map(this::parseNode).collect(Collectors.toMap(Node::getName, Node::get)));
        //dumpNodes();
        final Nodes hierarchy = getHierarchy(nodes.get(root));
        //dumpHierarchy(hierarchy, 0);
        hierarchy.weight = accumulateWeights(hierarchy);
        dumpWeights(hierarchy, 0);
        return 0;
    }

    private void dumpWeights(final Nodes hierarchy, final int level) {
        System.out.println(Utils.repeat(" ", level * 4) + hierarchy.parent.getName() + " (" + hierarchy.weight + "):");
        if (hierarchy.children != null) {
            for (Nodes child : hierarchy.children) {
                dumpWeights(child, level + 1);
            }
        }
    }

    private int accumulateWeights(final Nodes nodes) {
        if (null == nodes) {
            return 0;
        }
        if (null == nodes.children) {
            return nodes.parent.weight;
        }
        int childrenWeight = 0;
        for (Nodes children : nodes.children) {
            if (children != null) {
                children.weight = accumulateWeights(children);
                childrenWeight += children.weight;
            }
        }
        return nodes.parent.weight + childrenWeight;
    }

    class Node {
        private String name;
        private int weight;
        private List<String> children;

        Node(final String name, final int weight) {
            this(name, weight, null);
        }

        Node(final String name, final int weight, final List<String> children) {
            this.name = name.trim();
            this.weight = weight;
            if (null != children) {
                this.children = children.stream().map(String::trim).collect(Collectors.toList());
            } else {
                this.children = null;
            }
        }

        public String getName() {
            return name;
        }

        Node get() {
            return this;
        }

        @Override
        public String toString() {
            if (children == null) {
                return name + " (" + weight + ")";
            } else {
                return name + " (" + weight + ") -> " + children.stream().collect(Collectors.joining(","));
            }
        }
    }

    class Nodes {
        private final Node parent;
        private final Set<Nodes> children;
        private int weight;

        Nodes(final Node parent, final Set<Nodes> children) {
            this.parent = parent;
            this.children = children;
            this.weight = 0;
        }

        String getParentName() {
            return parent.getName();
        }

        @Override
        public String toString() {
            if (null == children) {
                return parent + " (" + weight + ")";
            } else {
                final String strChildren = children.stream().map(Nodes::getParentName).collect(Collectors.joining(","));
                return parent + " (" + weight + ") -> " + strChildren;
            }
        }
    }
}
