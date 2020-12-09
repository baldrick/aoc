package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.List;

public class Day9 {
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day9.txt");
        final Day9 day = new Day9(input);
        System.out.println("Part 1 = " + day.part1());
        System.out.println("Part 2 = " + day.part2());
    }

    private final String input;
    private String mutatedInput;
    private int characterCount;

    public Day9(final List<String> input) {
        this.input = input.get(0);
    }

    public int part1() {
        mutatedInput = input;
        characterCount = 0;
        while (removeGarbage());
        System.out.println("input is " + input.length() + " chars, but " + mutatedInput.length()
            + " with garbage removed (" + characterCount + " unescaped garbage characters)");
        System.out.println("Garbage-free input is " + mutatedInput);
        return score(mutatedInput);
    }

    private boolean removeGarbage() {
        final int garbageStart = findUnescaped('<', 0, false);
        if (-1 == garbageStart) {
            return false;
        }
        final int garbageEnd = findUnescaped('>', garbageStart + 1, true);
        if (-1 == garbageEnd) {
            throw new RuntimeException("Unclosed garbage starting at " + garbageStart);
        }
        final StringBuilder sb = new StringBuilder();
        if (garbageStart > 0) {
            sb.append(mutatedInput.substring(0, garbageStart));
        }
        if (garbageEnd < mutatedInput.length()) {
            sb.append(mutatedInput.substring(garbageEnd + 1));
        }
        mutatedInput = sb.toString();
        return true;
    }

    private int findUnescaped(final char target, final int start, final boolean countCharacters) {
        int pos = start;
        while (pos < mutatedInput.length()) {
            final char charPos = mutatedInput.charAt(pos);
            if ('!' == charPos) {
                pos += 2;
                continue;
            }
            if (target == charPos) {
                return pos;
            }
            if (countCharacters) {
                characterCount++;
            }
            pos++;
        }
        return -1;
    }

    private int score(final String in) {
        int pos = 0;
        long group = 0L;
        int score = 0;
        while (pos < in.length()) {
            final char charPos = in.charAt(pos);
            if ('!' == charPos) {
                pos += 2;
                continue;
            }
            if ('{' == charPos) {
                group++;
            }
            if ('}' == charPos) {
                score += group;
                group--;
            }
            pos++;
        }
        return score;
    }

    public int part2() {
        return characterCount;
    }
}
