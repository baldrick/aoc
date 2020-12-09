package com.gavinmanning;

import org.junit.Assert;
import org.junit.Test;

import java.util.Arrays;

public class Day8Test {
    @Test
    public void part1and2() {
        final String[] input = {
            "b inc 5 if a > 1",
            "a inc 1 if b < 5",
            "c dec -10 if a >= 1",
            "c inc -20 if c == 10"
        };
        final Day8 day8 = new Day8(Arrays.asList(input));
        Assert.assertEquals(1, day8.part1());
        Assert.assertEquals(10, day8.part2());
    }
}
