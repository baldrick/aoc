package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Day25 {
    private static final Logger log = LoggerFactory.getLogger(Day25.class);
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day25.txt");
        final Day25 day = new Day25(input);
        log.info("Part 1 = {}", day.part1());
        log.info("Part 2 = {}", day.part2());
    }

    private final List<String> input;
    private int[] tape = new int[10000];
    private boolean prepared;
    private String startState;
    private int diagnosticChecksumStep;
    private final Map<String, State> states;

    public Day25(final List<String> input) {
        this.input = input;
        this.prepared = false;
        states = new HashMap<>();
    }

    class State {
        private final String name;
        private final Action[] actions = new Action[2];

        class Action {
            private final int writeValue;
            private final int direction;
            private final String nextState;

            Action(final List<String> instructions) {
                writeValue = Integer.valueOf(getAfter(instructions.get(0), "    - Write the value ", 1));
                final String strDirection = getAfter(instructions.get(1), "    - Move one slot to the ", 1);
                if ("l".equalsIgnoreCase(strDirection)) {
                    direction = -1;
                } else if ("r".equalsIgnoreCase(strDirection)) {
                    direction = 1;
                } else {
                    throw new RuntimeException("Could not parse direction from " + instructions.get(1));
                }
                nextState = getAfter(instructions.get(2), "    - Continue with state ", 1);
            }

            void process(final ProgramState ps) {
                log.debug("{}, writing {}, moving {}, next state is {}", name, writeValue, direction == 1 ? "right" : "left", nextState);
                tape[ps.ptr] = writeValue;
                ps.ptr += direction;
                ps.state = nextState;
            }

            @Override
            public String toString() {
                return "w" + writeValue + (direction == 1 ? ">" : "<") + nextState;
            }
        }

        /*
        In state A:
          If the current value is 0:
            - Write the value 1.
            - Move one slot to the right.
            - Continue with state B.
        If the current value is 1:
            - Write the value 0.
            - Move one slot to the right.
            - Continue with state C.
         */

        State(final List<String> instructions) {
            this.name = getAfter(instructions.get(0), "In state ", 1);
            actions[0] = new Action(instructions.subList(2, 5));
            actions[1] = new Action(instructions.subList(6, 9));
        }

        public String getName() {
            return name;
        }

        public void process(final ProgramState ps) {
            actions[tape[ps.ptr]].process(ps);
        }

        @Override
        public String toString() {
            return "(" + name + "), 0=" + actions[0] + ", 1=" + actions[1];
        }
    }

    private void prepare() {
        if (!prepared) {
            for (int i = 0;  i < tape.length;  i++) {
                tape[i] = 0;
            }

            final List<String> steps = new ArrayList<>();
            for (String instruction : input) {
                if (instruction.startsWith("Begin in state ")) {
                    startState = parseStartState(instruction);
                } else if (instruction.startsWith("Perform a diagnostic checksum after")) {
                    diagnosticChecksumStep = parseDiagnosticChecksumStep(instruction);
                } else if (!instruction.trim().isEmpty()) {
                    steps.add(instruction);
                } else {
                    processSteps(steps);
                }
            }
            processSteps(steps);
            dumpStateMap();
        }
    }

    private void processSteps(final List<String> steps) {
        if (!steps.isEmpty()) {
            parseStep(steps);
            steps.clear();
        }
    }

    private void dumpStateMap() {
        for (Map.Entry<String, State> state : states.entrySet()) {
            log.info("{}: {}", state.getKey(), state.getValue());
        }
    }

    private String getAfter(final String s, final String search, final int chars) {
        final int start = search.length();
        return s.substring(start, start + chars);
    }

    private String parseStartState(final String instruction) {
        return getAfter(instruction, "Begin in state ", 1);
    }

    private int parseDiagnosticChecksumStep(final String instruction) {
        final int start = "Perform a diagnostic checksum after ".length();
        return Integer.valueOf(instruction.substring(start, instruction.indexOf(" ", start)));
    }

    private void parseStep(final List<String> steps) {
        final State state = new State(steps);
        states.put(state.getName(), state);
    }

    class ProgramState {
        private int ptr;
        private String state;

        ProgramState(final int ptr, final String state) {
            this.ptr = ptr;
            this.state = state;
        }
    }

    public int part1() {
        prepare();
        final ProgramState ps = new ProgramState(tape.length / 2, startState);
        for (int step = 0;  step < diagnosticChecksumStep;  step++) {
            if (step % 200000 == 0) {
                dumpState(step, ps);
            }
            states.get(ps.state).process(ps);
        }
        return Arrays.stream(tape).sum();
    }

    private void dumpState(final int step, final ProgramState ps) {
        final StringBuilder sb = new StringBuilder();
        final int bracket = 3;
        for (int i = tape.length / 2 - bracket;  i < tape.length / 2 + bracket;  i++) {
            if (ps.ptr == i) {
                sb.append("[").append(tape[i]).append("] ");
            } else {
                sb.append(" ").append(tape[i]).append("  ");
            }
        }
        log.info("... {}... after {} steps; about to run state {}.", sb.toString(), step, ps.state);
    }

    public int part2() {
        prepare();
        return 0;
    }
}
