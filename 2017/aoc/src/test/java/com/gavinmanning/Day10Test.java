package com.gavinmanning;

import org.junit.Assert;
import org.junit.Ignore;
import org.junit.Test;

import java.util.Arrays;

public class Day10Test {
    @Test
    public void part1() {
        final String[] input = { "3,4,1,5" };
        final Day10 day = new Day10(Arrays.asList(input), 5);
        Assert.assertEquals(12, day.part1());
    }

    @Test
    @Ignore
    public void part2() {
        Assert.assertEquals("a2582a3a0e66e6e86e3812dcb672a272", part2Test("The empty string"));
        Assert.assertEquals("33efeb34ea91902bb2f59c9920caa6cd", part2Test("AoC 2017"));
        Assert.assertEquals("3efbe78a8d82f29979031a4aa0b16a9d", part2Test("1,2,3"));
        Assert.assertEquals("63960835bcdc130f0b66d7ff4f6a5a8e", part2Test("1,2,4"));
    }

    private String part2Test(final String in) {
        final String[] input = { in };
        final Day10 day = new Day10(Arrays.asList(input), 256);
        return day.part2();
    }
}
