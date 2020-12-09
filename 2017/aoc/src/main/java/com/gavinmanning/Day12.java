package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

public class Day12 {
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day12.txt");
        final Day12 day = new Day12(input);
        System.out.println("Part 1 = " + day.part1());
        System.out.println("Part 2 = " + day.part2());
    }

    private final List<String> input;

    public Day12(final List<String> input) {
        this.input = input;
    }

    // How many programs are in group 0?
    public int part1() {
        final Map<Integer, Set<Integer>> programs = input.stream().collect(Collectors.toMap(Day12::getKey, Day12::getPrograms));
        return count(programs, new HashSet<>(), 0);
    }

    private int count(Map<Integer, Set<Integer>> programs, final Set<Integer> encountered, final int start) {
        encountered.add(start);
        if (programs.containsKey(start)) {
            final Set<Integer> children = programs.get(start);
            int count = 1;
            for (Integer child : children) {
                if (!encountered.contains(child)) {
                    System.out.println("Adding child " + child);
                    count += count(programs, encountered, child);
                }
            }
            return count;
        } else {
            throw new RuntimeException("Failed to find program " + start);
        }
    }

    private static int getKey(final String row) {
        //1999 <-> 1239, 1364
        final int arrows = row.indexOf("<->");
        if (-1 == arrows) {
            throw new RuntimeException("No arrows for " + row);
        }
        return Integer.valueOf(row.substring(0, arrows).trim());
    }

    private static Set<Integer> getPrograms(final String row) {
        final int arrows = row.indexOf("<->");
        if (-1 == arrows) {
            throw new RuntimeException("No arrows for " + row);
        }
        return Arrays.asList(row.substring(arrows + 3).trim().split(",")).stream().map(String::trim).map(Integer::valueOf).collect(Collectors.toSet());
    }

    public int part2() {
        final Map<Integer, Set<Integer>> programs = input.stream().collect(Collectors.toMap(Day12::getKey, Day12::getPrograms));
        final Set<Integer> encountered = new HashSet<>();
        int groups = 0;
        while (!programs.isEmpty()) {
            final int rootProgram = programs.entrySet().iterator().next().getKey();
            final int count = count(programs, encountered, rootProgram);
            System.out.println("Removed " + encountered.size() + " (" + count + ") programs connected to " + rootProgram);
            encountered.forEach(programs::remove);
            encountered.clear();
            groups++;
        }
        return groups;
    }
}
