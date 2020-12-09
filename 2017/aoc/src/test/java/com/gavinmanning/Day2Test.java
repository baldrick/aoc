package com.gavinmanning;

import org.junit.Test;

import java.util.Arrays;

import static org.junit.Assert.assertEquals;

public class Day2Test {
    @Test
    public void part1() {
        final String[] inputString = {
            "5 1 9 5",
            "7 5 3",
            "2 4 6 8"
        };
        assertEquals(18, Day2.part1(Arrays.asList(inputString)));
    }

    @Test
    public void part2() {
        final String[] inputString = {
            "5 9 2 8",
            "9 4 7 3",
            "3 8 6 5"
        };
        assertEquals(9, Day2.part2(Arrays.asList(inputString)));
    }
}
