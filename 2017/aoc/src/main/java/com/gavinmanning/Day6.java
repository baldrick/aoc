package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class Day6 {
    public static void main(String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day6.txt");
        final Day6 day = new Day6(input);
        System.out.println("Part 1 = " + day.part1());
        System.out.println("Part 2 = " + day.part2());
    }

    private final List<String> input;

    private Day6(final List<String> input) {
        this.input = input;
    }

    private int part1() {
        final ArrayList<Integer> memory = Arrays.asList(input.get(0).split("\\s"))
            .stream().map(Integer::valueOf).collect(Collectors.toCollection(ArrayList::new));
        final Set<ArrayList<Integer>> states = new HashSet<>();
        while (states.add(memory)) {
            final int max = memory.stream().max(Integer::compare).get();
            final int index = findMax(memory, max);
            redistribute(memory, index);
        }
        return states.size();
    }

    private int findMax(final ArrayList<Integer> memory, final int max) {
        int index = 0;
        while (memory.get(index) != max) {
            index++;
        }
        return index;
    }

    private void redistribute(final ArrayList<Integer> memory, final int index) {
        final int redistributeCount = memory.get(index);
        memory.set(index, 0);
        for (int i = 1;  i <= redistributeCount;  i++) {
            final int changeIndex = (index + i) % memory.size();
            memory.set(changeIndex, memory.get(changeIndex) + 1);
        }
    }

    private int part2() {
        final ArrayList<Integer> memory = Arrays.asList(input.get(0).split("\\s"))
            .stream().map(Integer::valueOf).collect(Collectors.toCollection(ArrayList::new));
        final Set<ArrayList<Integer>> states = new HashSet<>();
        while (states.add(memory)) {
            final int max = memory.stream().max(Integer::compare).get();
            final int index = findMax(memory, max);
            redistribute(memory, index);
        }
        final ArrayList<Integer> targetState = new ArrayList<>();
        targetState.addAll(memory);
        int steps = 0;
        while (true) {
            final int max = memory.stream().max(Integer::compare).get();
            final int index = findMax(memory, max);
            redistribute(memory, index);
            steps++;
            if (targetState.equals(memory)) {
                break;
            }
        }
        return steps;
    }
}
