package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.List;

public class Day15 {
    private static final Logger log = LoggerFactory.getLogger(Day15.class);
    public static void main(final String[] args) {
        final Day15 dayA = new Day15(116, 299, 1, 1);
        log.info("Part 1 = {}", dayA.part1());
        final Day15 dayB = new Day15(116, 299, 4, 8);
        log.info("Part 2 = {}", dayB.part2());
    }

    private final List<Generator> generators;

    public Day15(final long g1, final long g2, final long multiple1, final long multiple2) {
        this.generators = new ArrayList<>();
        generators.add(new Generator(g1, 16807, multiple1));
        generators.add(new Generator(g2,48271, multiple2));
    }

    class Generator {
        private final long factor;
        private long previous;
        private final long multiple;

        Generator(final long previous, final long factor, final long multiple) {
            this.previous = previous;
            this.factor = factor;
            this.multiple = multiple;
        }

        long generate() {
            do {
                previous = (previous * factor) % 2147483647;
            } while ((previous % multiple) != 0);
            return previous;
        }
    }

    public long generate(final int generator) {
        return generators.get(generator - 1).generate();
    }

    public int part1() {
        int matches = 0;
        for (int round = 0;  round < 40e6;  round++) {
            final long one = generate(1);
            final long two = generate(2);
            if ((one & 0xffff) == (two & 0xffff)) {
                matches++;
            }
        }
        return matches;
    }

    public int part2() {
        int matches = 0;
        for (int round = 0;  round < 5e6;  round++) {
            final long one = generate(1);
            final long two = generate(2);
            if ((one & 0xffff) == (two & 0xffff)) {
                matches++;
            }
        }
        return matches;
    }
}
