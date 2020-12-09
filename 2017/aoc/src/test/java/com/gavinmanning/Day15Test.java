package com.gavinmanning;

import org.junit.Assert;
import org.junit.Test;

public class Day15Test {
    @Test
    public void part1a() {
        final Day15 day = new Day15(65, 8921, 1, 1);
        Assert.assertEquals(1092455, day.generate(1));
        Assert.assertEquals(430625591, day.generate(2));

        Assert.assertEquals(1181022009, day.generate(1));
        Assert.assertEquals(1233683848, day.generate(2));

        Assert.assertEquals(245556042, day.generate(1));
        Assert.assertEquals(1431495498, day.generate(2));

        Assert.assertEquals(1744312007, day.generate(1));
        Assert.assertEquals(137874439, day.generate(2));

        Assert.assertEquals(1352636452, day.generate(1));
        Assert.assertEquals(285222916, day.generate(2));
    }

    @Test
    public void part1b() {
        final Day15 day = new Day15(65, 8921, 1, 1);
        Assert.assertEquals(588, day.part1());
        //Assert.assertEquals(10, day.part2(6,20));
    }
}
