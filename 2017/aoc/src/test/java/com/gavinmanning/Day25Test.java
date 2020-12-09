package com.gavinmanning;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;

public class Day25Test {
    private Day25 day;

    @Before
    public void setup() {
        final String[] input = {
            "Begin in state A.",
            "Perform a diagnostic checksum after 6 steps.",
            "",
            "In state A:",
            "  If the current value is 0:",
            "    - Write the value 1.",
            "    - Move one slot to the right.",
            "    - Continue with state B.",
            "  If the current value is 1:",
            "    - Write the value 0.",
            "    - Move one slot to the left.",
            "    - Continue with state B.",
            "",
            "In state B:",
            "  If the current value is 0:",
            "    - Write the value 1.",
            "    - Move one slot to the left.",
            "    - Continue with state A.",
            "  If the current value is 1:",
            "    - Write the value 1.",
            "    - Move one slot to the right.",
            "    - Continue with state A."
        };
        day = new Day25(Arrays.asList(input));
    }

    @Test
    public void part1() {
        Assert.assertEquals(3, day.part1());
    }

    @Test
    public void part2() {
        Assert.assertEquals(19, day.part2());
    }
}
