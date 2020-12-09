package com.gavinmanning;

import org.junit.Assert;
import org.junit.Test;

public class Day17Test {
    @Test
    public void part1() {
        final Day17 day = new Day17(3);
        Assert.assertEquals(638, day.part1());
    }
}
