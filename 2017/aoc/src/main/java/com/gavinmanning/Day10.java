package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

public class Day10 {
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day10.txt");
        final Day10 day = new Day10(input, 256);
        System.out.println("Part 1 = " + day.part1());
        System.out.println("Part 2 = " + day.part2());
    }

    private final String input;
    private final int listSize;
    private int start;
    private int skip;

    public Day10(final List<String> input, final int listSize) {
        this.input = input.get(0);
        this.listSize = listSize;
    }

    public int part1() {
        List<Integer> list = createList(listSize);
        final List<Integer> lengths = Arrays.asList(input.split(",")).stream().map(Integer::valueOf).collect(Collectors.toList());
        skip = 0;
        start = 0;
        list = hash(list, lengths);
        return list.get(0) * list.get(1);
    }

    private List<Integer> hash(List<Integer> list, final List<Integer> lengths) {
        for (Integer length : lengths) {
            System.out.print("(skip=" + skip + ") reversing " + length + " from " + list + " starting at " + start);
            list = reverse(list, start, length);
            System.out.println(" = " + list);
            start += length;
            start += skip;
            start = start % listSize;
            skip++;
        }
        System.out.println("Mutated list is " + list);
        return list;
    }

    private List<Integer> reverse(final List<Integer> list, final int start, final int length) {
        final List<Integer> toReverse = new ArrayList<>();
        if (start + length > list.size()) {
            toReverse.addAll(list.subList(start, listSize));
            toReverse.addAll(list.subList(0, length - (listSize - start)));
        } else {
            toReverse.addAll(list.subList(start, start + length));
        }
        System.out.print("; reversing " + toReverse);
        Collections.reverse(toReverse);
        System.out.print(" = " + toReverse);

        final List<Integer> result = new ArrayList<>();
        if (start + length > list.size()) {
            final int overspill = length - (listSize - start);
            // add start of reversed string
            result.addAll(toReverse.subList(length - overspill, length));
            // add unreversed
            result.addAll(list.subList(overspill, start));
            // add end of reversed string
            result.addAll(toReverse.subList(0, length - overspill));
        } else {
            result.addAll(list.subList(0, start));
            result.addAll(toReverse);
            result.addAll(list.subList(start + length, listSize));
        }
        return result;
    }

    private List<Integer> createList(final int end) {
        final List<Integer> list = new ArrayList<>();
        for (int pos = 0;  pos < end;  pos++) {
            list.add(pos);
        }
        return list;
    }

    private List<Integer> createList(final String in) {
        final List<Integer> list = new ArrayList<>();
        for (int i = 0;  i < in.length();  i++) {
            list.add((int) in.charAt(i));
        }
        list.addAll(Arrays.asList(17, 31, 73, 47, 23));
        return list;
    }

    public String part2() {
        List<Integer> list = createList(listSize);
        final List<Integer> lengths = createList(input);
        start = 0;
        skip = 0;
        for (int round = 1;  round <= 64;  round++) {
            System.out.println("round " + round);
            list = hash(list, lengths);
        }
        final StringBuilder sb = new StringBuilder();
        for (int block = 0;  block < 16;  block++) {
            final int result = xor(list, block);
            sb.append(String.format("%02x", result));
        }
        return sb.toString();
    }

    private int xor(final List<Integer> list, final int start) {
        int result = list.get(start * 16);
        for (int i = 1;  i < 16;  i++) {
            result ^= list.get((start * 16) + i);
        }
        return result;
    }
}
