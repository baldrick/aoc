package com.gavinmanning;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;

public class Day18Test {
    private Day18 day;

    @Before
    public void setup() {
        final String[] input = {
            "set a 1",
            "add a 2",
            "mul a a",
            "mod a 5",
            "snd a",
            "set a 0",
            "rcv a",
            "jgz a -1",
            "set a 1",
            "jgz a -2"
        };
        day = new Day18(Arrays.asList(input));
    }

    @Test
    public void part1() {
        Assert.assertEquals(4, day.part1());
    }

    @Test
    public void part2() {
        Assert.assertEquals(0, day.part2());
    }
}
