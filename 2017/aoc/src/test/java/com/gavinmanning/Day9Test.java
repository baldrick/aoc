package com.gavinmanning;

import org.junit.Assert;
import org.junit.Test;

import java.util.Arrays;

public class Day9Test {
    @Test
    public void part1() {
        final String[] input = { "{{<!!>},{<!!>},{<!!>},{<!!>}}" };
        final Day9 day = new Day9(Arrays.asList(input));
        Assert.assertEquals(9, day.part1());
    }

    @Test
    public void part2() {
        Assert.assertEquals(0, part2Test("<>"));
        Assert.assertEquals(17, part2Test("<random characters>"));
        Assert.assertEquals(3, part2Test("<<<<>"));
        Assert.assertEquals(2, part2Test("<{!>}>"));
        Assert.assertEquals(0, part2Test("<!!>"));
        Assert.assertEquals(0, part2Test("<!!!>>"));
        Assert.assertEquals(10, part2Test("<{o\"i!a,<{i<a>"));
    }

    private int part2Test(final String in) {
        final String[] input = { in };
        final Day9 day = new Day9(Arrays.asList(input));
        day.part1();
        return day.part2();
    }
}
