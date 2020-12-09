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

public class Day18 {
    private static final Logger log = LoggerFactory.getLogger(Day18.class);
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day18.txt");
        final Day18 day = new Day18(input);
        log.info("Part 1 = {}", day.part1());
        log.info("Part 2 = {}", day.part2());
    }

    private final List<String> input;
    private boolean prepared;
    private final Map<String, Long> registers;
    private final List<Instruction> instructions;
    private long sound;

    public Day18(final List<String> input) {
        this.input = input;
        prepared = false;
        registers = new HashMap<>();
        resetRegisters();
        instructions = new ArrayList<>();
    }

    private void resetRegisters() {
        registers.clear();
    }

    enum Command {
        SET,
        MUL,
        JGZ,
        ADD,
        MOD,
        SND,
        RCV
    }

    class Recover extends Throwable  {
        private final long frequency;

        Recover(final long frequency) {
            this.frequency = frequency;
        }
    }

    class Instruction {
        private final Command command;
        private final String register; // can also be a value!
        private final String value; // can also be a register!

        Instruction(final String in) {
            final String strCommand = in.substring(0, 3);
            command = Command.valueOf(strCommand.toUpperCase());
            register = in.substring(4, 5);
            if (in.length() > 5) {
                value = in.substring(6);
            } else {
                value = null;
            }
        }

        long process() throws Recover {
            switch (command) {
                case SET: {
                    registers.put(register, getValue());
                    break;
                }
                case MUL: {
                    final long currentValue = getRegisterValue(register);
                    registers.put(register, currentValue * getValue());
                    break;
                }
                case JGZ: {
                    if (getRegister() > 0) {
                        return getValue();
                    }
                    break;
                }
                case ADD: {
                    final long currentValue = getRegisterValue(register);
                    registers.put(register, currentValue + getValue());
                    break;
                }
                case MOD: {
                    final long currentValue = getRegisterValue(register);
                    registers.put(register, currentValue % getValue());
                    break;
                }
                case SND: {
                    sound = getRegister();
                    break;
                }
                case RCV: {
                    if (getRegister() != 0) {
                        throw new Recover(sound);
                    }
                    break;
                }
                default:
                    throw new RuntimeException("Unhandled command " + command);
            }
            return 1;
        }

        private long getRegister() {
            try {
                return Long.valueOf(register);
            } catch (NumberFormatException ex) {
                return getRegisterValue(register);
            }
        }

        private long getRegisterValue(final String registerToRetrieve) {
            final Long registerValue = registers.get(registerToRetrieve);
            if (null == registerValue) {
                return 0;
            } else {
                return registerValue;
            }
        }

        private long getValue() {
            try {
                return Long.valueOf(value);
            } catch (NumberFormatException ex) {
                return getRegisterValue(value);
            }
        }

        @Override
        public String toString() {
            return command + " " + register + "(" + getRegisterValue(register) + ") " + value + " (" + getValue() + ")";
        }
    }

    private void prepare() {
        if (!prepared) {
            instructions.addAll(input.stream().map(Instruction::new).collect(Collectors.toList()));
        }
    }

    public long part1() {
        prepare();
        int pc = 0;
        sound = 0;
        try {
            while (pc < instructions.size()) {
                log.info("Executing #{}: {}", pc, instructions.get(pc));
                pc += instructions.get(pc).process();
            }
            throw new RuntimeException("Finished without recover...");
        } catch (Recover r) {
            return r.frequency;
        }
    }

    public int part2() {
        prepare();
        return 0;
    }
}
