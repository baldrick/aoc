package com.gavinmanning;

import org.junit.Assert;
import org.junit.Ignore;
import org.junit.Test;

import java.util.Arrays;

public class Day7Test {
    @Test
    @Ignore
    public void part2() {
        final String[] inputString = {
            "pbga (66)",
            "xhth (57)",
            "ebii (61)",
            "havc (66)",
            "ktlj (57)",
            "fwft (72) -> ktlj, cntj, xhth",
            "qoyq (66)",
            "padx (45) -> pbga, havc, qoyq",
            "tknk (41) -> ugml, padx, fwft",
            "jptl (61)",
            "ugml (68) -> gyxo, ebii, jptl",
            "gyxo (61)",
            "cntj (57)"
        };
        final Day7 day7 = new Day7(Arrays.asList(inputString));
        Assert.assertEquals(60, day7.part2("tknk"));
    }
}
