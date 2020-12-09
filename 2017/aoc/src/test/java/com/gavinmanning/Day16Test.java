package com.gavinmanning;

import org.junit.Assert;
import org.junit.Test;

import java.util.Collections;

public class Day16Test {
    @Test
    public void part1() {
        Assert.assertEquals("cdeab", part1Test("s3"));
        Assert.assertEquals("eabcd", part1Test("s1"));
        Assert.assertEquals("bcdea", part1Test("s4"));
        Assert.assertEquals("abced", part1Test("x3/4"));
        Assert.assertEquals("abedc", part1Test("pc/e"));
        Assert.assertEquals("baedc", part1Test("s1,x3/4,pe/b"));
    }

    @Test
    public void part1b() {
        final Day16 day = new Day16(Collections.singletonList("s1"), "baedc");
        Assert.assertEquals("cbaed", day.part1());
    }

    private String part1Test(final String input) {
        final Day16 day = new Day16(Collections.singletonList(input), 5);
        return day.part1();
    }

    @Test
    public void part2() {
        final Day16 day = new Day16(Collections.singletonList("s1,x3/4,pe/b"), "baedc");
        Assert.assertEquals("ceadb", day.part2());
    }

    @Test
    public void xorTest() {
        char[] ca = new char[3];
        ca[0] = 'a';
        ca[1] = 'b';
        ca[2] = 'c';
        System.out.println(ca);
        ca[0] = (char) (ca[0] ^ ca[2]);
        ca[2] = (char) (ca[0] ^ ca[2]);
        ca[0] = (char) (ca[0] ^ ca[2]);
        System.out.println(ca);
    }
}
