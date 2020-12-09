package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

public class Day16 {
    private static final Logger log = LoggerFactory.getLogger(Day16.class);
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day16.txt");
        final Day16 day = new Day16(input, 16);
        log.info("Part 1 = {}", day.part1());
        log.info("Part 2 = {}", day.part2());
    }

    private final List<String> input;
    private int programCount;
    private boolean prepared;
    private List<Character> programs;

    public Day16(final List<String> input, final int programCount) {
        this.input = input;
        this.programCount = programCount;
        prepared = false;
        programs = new ArrayList<>();
    }

    public Day16(final List<String> input, final String programs) {
        this(input, 0);
        for (int i = 0;  i < programs.length();  i++) {
            this.programs.add(programs.charAt(i));
        }
        programCount = programs.length();
    }

    private void prepare() {
        if (!prepared) {
            if (programs.isEmpty()) {
                for (int i = 0; i < programCount; i++) {
                    programs.add((char) ((int) 'a' + i));
                }
            }
            log.info("Start programs: {}", programs);
        }
    }

    public String part1() {
        prepare();
        Arrays.stream(input.get(0).split(",")).forEach(this::process);
        return programs.stream().map(c -> Character.toString(c)).collect(Collectors.joining());
    }

    public void process(final String i) {
        // sX - spin x
        // xA/B - exchange positions A & B
        // pA/B - exchange programs A & B
        //final String before = programs.toString();
        switch (i.charAt(0)) {
            case 's':
                spin(Integer.valueOf(i.substring(1)));
                break;
            case 'x':
                final int sep = i.indexOf('/');
                exchangePlaces(Integer.valueOf(i.substring(1, sep)), Integer.valueOf(i.substring(sep + 1)));
                break;
            case 'p':
                exchangePrograms(i.substring(1,2), i.substring(3,4));
                break;
            default:
                throw new RuntimeException("Unhandled instruction " + i);
        }
        //log.debug("{} --> {} --> {}", before, i, programs);
    }

    private void spin(final int count) {
        Collections.rotate(programs, count);
    }

    private void exchangePlaces(final int a, final int b) {
        programs.set(a, (char) (programs.get(a) ^ programs.get(b)));
        programs.set(b, (char) (programs.get(a) ^ programs.get(b)));
        programs.set(a, (char) (programs.get(a) ^ programs.get(b)));
    }

    private void exchangePrograms(final String a, final String b) {
        exchangePlaces(findProgram(a), findProgram(b));
    }

    private int findProgram(final String p) {
        final char pc = p.charAt(0);
        for (int i = 0;  i < this.programCount;  i++) {
            if (programs.get(i) == pc) {
                return i;
            }
        }
        throw new RuntimeException("Program " + p + " not found");
    }

    public String part2() {
        prepare();
        for (int round = 0;  round < 1e9;  round++) {
            if (round % 1e6 == 0) {
                log.info("Round {}", round);
            }
            Arrays.stream(input.get(0).split(",")).forEach(this::process);
        }
        return programs.stream().map(c -> Character.toString(c)).collect(Collectors.joining());
    }
}
