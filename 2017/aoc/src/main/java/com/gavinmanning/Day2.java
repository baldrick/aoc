package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class Day2 {
    public static void main(String[] args) throws URISyntaxException, IOException {
        final List<String> input = Utils.getInput("day2.txt");
        System.out.println("part1 = " + part1(input) + ", part2 = " + part2(input));
    }

    public static int part1(final List<String> input) {
        final List<List<Integer>> inputList = Day2.parse(input);
        return inputList.stream().map(Day2::maxDiff).collect(Collectors.summingInt(Integer::intValue));
    }

    public static int part2(final List<String> input) {
        final List<List<Integer>> inputList = Day2.parse(input);
        return inputList.stream().map(Day2::evenDivision).collect(Collectors.summingInt(Integer::intValue));
    }

    private static int maxDiff(final List<Integer> row) {
        return Math.abs(row.stream().min(Integer::compare).get() - row.stream().max(Integer::compare).get());
    }

    private static int evenDivision(final List<Integer> row) {
        final int len = row.size();
        for (int i = 0; i < len; i++) {
            for (int j = 0; j < len; j++) {
                if (i != j) {
                    final int iVal = row.get(i);
                    final int jVal = row.get(j);
                    if (iVal % jVal == 0) {
                        return iVal / jVal;
                    }
                }
            }
        }
        throw new RuntimeException("Failed to find even division in " + row);
    }

    private static List<List<Integer>> parse(final List<String> input) {
        final List<List<Integer>> output = new ArrayList<>();
        for (String row : input) {
            final List<String> rowList = Arrays.asList(row.split("\\s"));
            final List<Integer> rowListInteger = rowList.stream().map(Integer::valueOf).collect(Collectors.toList());
            output.add(rowListInteger);
        }
        return output;
    }
}
