package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.List;

public class Day17 {
    private static final Logger log = LoggerFactory.getLogger(Day17.class);
    public static void main(final String[] args) {
        final Day17 dayA = new Day17(349);
        log.info("Part 1 = {}", dayA.part1());
        final Day17 dayB = new Day17(349);
        log.info("Part 2 = {}", dayB.part2());
    }

    private final int steps;
    private final List<Integer> buffer;

    public Day17(final int steps) {
        this.steps = steps;
        buffer = new ArrayList<>();
    }

    public int part1() {
        buffer.add(0);
        final int pos = spin(1500);
        return buffer.get(pos + 1);
    }

    public int part2() {
        buffer.add(0);

        int indexOne = 0;
        int pos = 0;
        for (int i = 1;  i <= 50_000_000;  i++) {
            if (i % 100_000 == 0) log.info("Spin {}", i);
            pos = (pos + steps) % i;
            pos++;
            if (pos == 1) {
                indexOne = i;
            }
        }
        return indexOne;
    }
/*
    step   index 1
    0
    1       1
    2       2
    3       2
    4       2
    5       5
    6       5
    7       5
    8       5
    9       9
    10      9
    11      9
    12      12
    13      12
    14      12
    15      12
    16      16
    17      16
    18      16
    19      16
    20
    */

    private int spin(final int count) {
        int pos = 0;
        for (int i = 1;  i <= count;  i++) {
            if (i % 100_000 == 0) log.info("Spin {}", i);
            pos = (pos + steps) % i;
            pos++;
            buffer.add(pos, i);
            dumpBuffer(i, pos);
        }
        //dumpBuffer(pos);
        return pos;
    }

    private void dumpBuffer(final int step, final int pos) {
        final StringBuilder sb = new StringBuilder();
        for (int i = 0;  i < buffer.size();  i++) {
            if (i == pos) sb.append('(');
            sb.append(buffer.get(i));
            if (i == pos) sb.append(')');
            sb.append(',');
        }
        log.info("{}: {}", step, sb);
    }
}
