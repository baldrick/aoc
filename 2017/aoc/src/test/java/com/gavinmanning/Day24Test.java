package com.gavinmanning;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;

public class Day24Test {
    private Day24 day;

    @Before
    public void setup() {
        final String[] input = {
            "0/2",
            "2/2",
            "2/3",
            "3/4",
            "3/5",
            "0/1",
            "10/1",
            "9/10"
        };
        day = new Day24(Arrays.asList(input));
    }

    @Test
    public void part1() {
        Assert.assertEquals(31, day.part1());
    }

    @Test
    public void part2() {
        Assert.assertEquals(19, day.part2());
    }
}
