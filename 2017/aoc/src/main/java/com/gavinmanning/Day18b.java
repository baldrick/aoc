package com.gavinmanning;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Queue;
import java.util.stream.Collectors;

public class Day18b {
    private static final Logger log = LoggerFactory.getLogger(Day18b.class);
    public static void main(final String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day18.txt");
        final Day18b day = new Day18b(input);
        log.info("Part 2 = {}", day.part2());
    }

    private final List<String> input;
    private boolean prepared;
    private final List<Map<String, Long>> registers;
    private final List<Instruction> instructions;

    public Day18b(final List<String> input) {
        this.input = input;
        prepared = false;
        registers = new ArrayList<>();
        registers.add(new HashMap<>());
        registers.add(new HashMap<>());
        instructions = new ArrayList<>();
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

        long process(final Map<String, Long> registers, final Queue<Long> receiveQueue, final Queue<Long> sendQueue) {
            switch (command) {
                case SET: {
                    registers.put(register, getValue(registers));
                    break;
                }
                case MUL: {
                    final long currentValue = getRegisterValue(registers, register);
                    registers.put(register, currentValue * getValue(registers));
                    break;
                }
                case JGZ: {
                    if (getRegister(registers) > 0) {
                        return getValue(registers);
                    }
                    break;
                }
                case ADD: {
                    final long currentValue = getRegisterValue(registers, register);
                    registers.put(register, currentValue + getValue(registers));
                    break;
                }
                case MOD: {
                    final long currentValue = getRegisterValue(registers, register);
                    registers.put(register, currentValue % getValue(registers));
                    break;
                }
                case SND: {
                    sendQueue.add(getRegister(registers));
                    registers.put("send", getRegisterValue(registers, "send") + 1);
                    break;
                }
                case RCV: {
                    if (receiveQueue.isEmpty()) {
                        return 0; // don't change pc => wait until messageQueue contains something
                    }
                    registers.put(register, receiveQueue.poll());
                    break;
                }
                default:
                    throw new RuntimeException("Unhandled command " + command);
            }
            return 1;
        }

        private long getRegister(final Map<String, Long> registers) {
            try {
                return Long.valueOf(register);
            } catch (NumberFormatException ex) {
                return getRegisterValue(registers, register);
            }
        }

        private long getRegisterValue(final Map<String, Long> registers ,final String registerToRetrieve) {
            final Long registerValue = registers.get(registerToRetrieve);
            if (null == registerValue) {
                return 0;
            } else {
                return registerValue;
            }
        }

        private long getValue(final Map<String, Long> registers) {
            try {
                return Long.valueOf(value);
            } catch (NumberFormatException ex) {
                return getRegisterValue(registers, value);
            }
        }

        @Override
        public String toString() {
            return command + " " + register + " " + value;
        }
    }

    private void prepare() {
        if (!prepared) {
            instructions.addAll(input.stream().map(Instruction::new).collect(Collectors.toList()));
        }
    }

    public long part2() {
        prepare();
        int[] pc = new int[2];
        pc[0] = 0;
        pc[1] = 0;
        long[] pcChange = new long[2];
        pcChange[0] = 0;
        pcChange[1] = 0;
        final List<Queue<Long>> queues = new ArrayList<>();
        queues.add(new ArrayDeque<>());
        queues.add(new ArrayDeque<>());
        // Run until neither pc changes - at which point we've reached deadlock
        do {
            for (int instance = 0;  instance < 2;  instance++) {
                do {
                    log.info("Executing {} - #{}: {} (inQ={}, outQ={})",
                        instance, pc[instance], instructions.get(pc[instance]),
                        Utils.dumpQueue(queues.get(instance)), Utils.dumpQueue(queues.get(1 - instance)));
                    pcChange[instance] = instructions.get(pc[instance]).process(
                        registers.get(instance), queues.get(instance), queues.get(1 - instance));
                    pc[instance] += pcChange[instance];
                } while (pcChange[instance] != 0);
            }
        } while ((pcChange[0] != 0) || (pcChange[1] != 0));
        log.info("0: {}", Utils.dumpMap(registers.get(0)));
        log.info("1: {}", Utils.dumpMap(registers.get(1)));
        return registers.get(1).get("send");
    }
}
