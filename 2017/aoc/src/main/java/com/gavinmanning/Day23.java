package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Day23 {
    private static final Logger log = LoggerFactory.getLogger(Day23.class);
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day23.txt");
        final Day23 day = new Day23(input);
        log.info("Part 1 = {}", day.part1());
        log.info("Part 2 = {}", day.part2());
    }

    private final List<String> input;
    private boolean prepared;
    private final Map<String, Integer> registers;
    private final List<Instruction> instructions;
    private int mulCount;

    public Day23(final List<String> input) {
        this.input = input;
        prepared = false;
        registers = new HashMap<>();
        resetRegisters();
        instructions = new ArrayList<>();
    }

    private void resetRegisters() {
        registers.put("a", 0);
        registers.put("b", 0);
        registers.put("c", 0);
        registers.put("d", 0);
        registers.put("e", 0);
        registers.put("f", 0);
        registers.put("g", 0);
        registers.put("h", 0);
    }

    enum Command {
        SET,
        SUB,
        MUL,
        JNZ
    };

    class Instruction {
        private final Command command;
        private final String register; // can also be a value!
        private final String value; // can also be a register!

        Instruction(final String in) {
            final String strCommand = in.substring(0, 3);
            command = Command.valueOf(strCommand.toUpperCase());
            register = in.substring(4, 5);
            value = in.substring(6);
        }

        int process() {
            switch (command) {
                case SET: {
                    registers.put(register, getValue());
                    break;
                }
                case SUB: {
                    final int currentValue = registers.get(register);
                    registers.put(register, currentValue - getValue());
                    break;
                }
                case MUL: {
                    mulCount++;
                    final int currentValue = registers.get(register);
                    registers.put(register, currentValue * getValue());
                    break;
                }
                case JNZ: {
                    if (getRegister() != 0) {
                        return getValue();
                    }
                    break;
                }
                default:
                    throw new RuntimeException("Unhandled command " + command);
            }
            return 1;
        }

        private int getRegister() {
            try {
                return Integer.valueOf(register);
            } catch (NumberFormatException ex) {
                return registers.get(register);
            }
        }

        private int getValue() {
            try {
                return Integer.valueOf(value);
            } catch (NumberFormatException ex) {
                return registers.get(value);
            }
        }

        @Override
        public String toString() {
            return command + " " + register + " " + value + " (" + getValue() + ")";
        }
    }

    /*
    set X Y sets register X to the value of Y.
    sub X Y decreases register X by the value of Y.
    mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
    jnz X Y jumps with an offset of the value of Y, but only if the value of X is not zero.
           (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
     */

    private void prepare() {
        if (!prepared) {
            instructions.addAll(input.stream().map(Instruction::new).collect(Collectors.toList()));
        }
    }

    public int part1() {
        prepare();
        int pc = 0;
        mulCount = 0;
        while (pc < instructions.size()) {
            log.info("Executing #{}: {}", pc, instructions.get(pc));
            pc += instructions.get(pc).process();
        }
        return mulCount;
    }

    public int part2() {
        prepare();
        int pc = 0;
        mulCount = 0;
        resetRegisters();
        registers.put("a", 1);
        while (pc < instructions.size()) {
            log.info("Executing #{}: {}", pc, instructions.get(pc));
            pc += instructions.get(pc).process();
        }
        return registers.get("h");
    }
}
