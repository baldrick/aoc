package com.gavinmanning;

import org.junit.Assert;
import org.junit.Test;

import java.util.Arrays;

public class Day12Test {
    @Test
    public void part1() {
        final String[] input = {
            "0 <-> 2",
            "1 <-> 1",
            "2 <-> 0, 3, 4",
            "3 <-> 2, 4",
            "4 <-> 2, 3, 6",
            "5 <-> 6",
            "6 <-> 4, 5"
        };
        final Day12 day = new Day12(Arrays.asList(input));
        Assert.assertEquals(6, day.part1());
    }

    @Test
    public void part2() {
    }
}
