package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Day5 {
    public static void main(String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day5.txt");
        final Day5 day5 = new Day5(input);
        System.out.println("Steps to escape the maze = " + day5.part1());
        System.out.println("Steps to escape the strange maze = " + day5.part2());
    }

    private final List<String> input;

    private Day5(final List<String> input) {
        this.input = input;
    }

    private int part1() {
        final ArrayList<Integer> intInput = input.stream().map(Integer::valueOf).collect(Collectors.toCollection(ArrayList::new));
        int sp = 0;
        int steps = 0;
        while (sp >= 0 && sp < intInput.size()) {
            final int jumpBy = intInput.get(sp);
            intInput.set(sp, jumpBy + 1);
            sp += jumpBy;
            steps++;
        }
        return steps;
    }

    private int part2() {
        final ArrayList<Integer> intInput = input.stream().map(Integer::valueOf).collect(Collectors.toCollection(ArrayList::new));
        int sp = 0;
        int steps = 0;
        while (sp >= 0 && sp < intInput.size()) {
            final int jumpBy = intInput.get(sp);

            // after each jump, if the offset was three or more, instead decrease it by 1. Otherwise, increase it by 1 as before.
            if (jumpBy >= 3) {
                intInput.set(sp, jumpBy - 1);
            } else {
                intInput.set(sp, jumpBy + 1);
            }
            sp += jumpBy;
            steps++;
        }
        return steps;
    }
}
