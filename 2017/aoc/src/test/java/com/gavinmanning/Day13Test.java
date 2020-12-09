package com.gavinmanning;

import org.junit.Assert;
import org.junit.Test;

import java.util.Arrays;

public class Day13Test {
    @Test
    public void part1and2() {
        final String[] input = {
            "0: 3",
            "1: 2",
            "4: 4",
            "6: 4"
        };
        final Day13 day = new Day13(Arrays.asList(input));
        Assert.assertEquals(24, day.part1());
        Assert.assertEquals(10, day.part2(6,20));
    }
}
