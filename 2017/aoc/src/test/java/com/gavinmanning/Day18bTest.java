package com.gavinmanning;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

import java.util.Arrays;

public class Day18bTest {
    private Day18b day;

    @Before
    public void setup() {
        final String[] input = {
            "snd 1",
            "snd 2",
            "snd p",
            "rcv a",
            "rcv b",
            "rcv c",
            "rcv d"
        };
        day = new Day18b(Arrays.asList(input));
    }

    @Test
    public void part2() {
        Assert.assertEquals(3, day.part2());
    }
}
