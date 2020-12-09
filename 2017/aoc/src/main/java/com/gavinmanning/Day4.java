package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class Day4 {
    public static void main(String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day4.txt");
        final Day4 day4 = new Day4();
        day4.part1(input);
        day4.part2(input);
    }

    private void part1(final List<String> input) throws URISyntaxException, IOException {
        System.out.println(input.stream().map(Day4::allUnique).collect(Collectors.summingInt(Integer::intValue)) + " valid passphrases");
    }

    private void part2(final List<String> input) {
        System.out.println(input.stream().map(Day4::anyAnagram).collect(Collectors.summingInt(Integer::intValue)) + " valid non-anagram passphrases");
    }

    private static int allUnique(final String row) {
        final List<String> rowList = Arrays.asList(row.split("\\s"));
        return rowList.stream().count() == rowList.stream().distinct().count() ? 1 : 0;
    }

    private static int anyAnagram(final String row) {
        final List<String> rowList = Arrays.asList(row.split("\\s"));
        return rowList.stream().count() == rowList.stream().map(Day4::sortWord).distinct().count() ? 1 : 0;
    }

    private static String sortWord(final String word) {
        final char[] caWord = word.toCharArray();
        Arrays.sort(caWord);
        return new String(caWord);
    }
}
