package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Day8 {
    public static void main(String[] args) throws IOException, URISyntaxException {
        final List<String> input = Utils.getInput("day8.txt");
        final Day8 day = new Day8(input);
        System.out.println("Part 1 = " + day.part1());
        System.out.println("Part 2 = " + day.part2());
    }

    private final List<String> input;

    public Day8(final List<String> input) {
        this.input = input;
    }

    public int part1() {
        final Map<String, RegisterValue> registers = new HashMap<>();
        input.stream().map(Instruction::new).forEach(instruction -> processInstruction(instruction, registers));
        return registers.values().stream().map(RegisterValue::getValue).max(Integer::compare).get();
    }

    public int part2() {
        final Map<String, RegisterValue> registers = new HashMap<>();
        input.stream().map(Instruction::new).forEach(instruction -> processInstruction(instruction, registers));
        return registers.values().stream().map(RegisterValue::getMaxValue).max(Integer::compare).get();
    }

    private static void processInstruction(final Instruction instruction, final Map<String, RegisterValue> registers) {
        if (instruction.conditional.isTrue(registers)) {
            instruction.process(registers);
        }
    }

    enum Comparator {
        GT,
        LT,
        EQUAL,
        GE,
        LE,
        NE
    }

    class Conditional {
        private final String register;
        private final Comparator comparator;
        private final int operand;

        Conditional(final String register, final String comparator, final int operand) {
            this.register = register;
            if (comparator.equals(">")) {
                this.comparator = Comparator.GT;
            } else if (comparator.equals("<")) {
                this.comparator = Comparator.LT;
            } else if (comparator.equals("==")) {
                this.comparator = Comparator.EQUAL;
            } else if (comparator.equals(">=")) {
                this.comparator = Comparator.GE;
            } else if (comparator.equals("<=")) {
                this.comparator = Comparator.LE;
            } else if (comparator.equals("!=")) {
                this.comparator = Comparator.NE;
            } else {
                throw new RuntimeException("Unknown comparator " + comparator);
            }
            this.operand = operand;
        }

        boolean isTrue(final Map<String, RegisterValue> registers) {
            final int registerValue;
            if (registers.containsKey(register)) {
                registerValue = registers.get(register).value;
            } else {
                registerValue = 0;
            }
            switch (comparator) {
                case GT:
                    return registerValue > operand;
                case LT:
                    return registerValue < operand;
                case EQUAL:
                    return registerValue == operand;
                case GE:
                    return registerValue >= operand;
                case LE:
                    return registerValue <= operand;
                case NE:
                    return registerValue != operand;
                default:
                    throw new RuntimeException("Unhanded comparator " + comparator);
            }
        }
    }

    enum Operator {
        DEC,
        INC
    }

    class Instruction {
        private final String register;
        private final Operator operator;
        private final int amount;
        private final Conditional conditional;

        Instruction(final String strInstruction) {
            final String[] splitInstruction = strInstruction.split("\\s");
            register = splitInstruction[0].trim();
            operator = Operator.valueOf(splitInstruction[1].trim().toUpperCase());
            amount = Integer.valueOf(splitInstruction[2]);
            conditional = new Conditional(splitInstruction[4], splitInstruction[5], Integer.valueOf(splitInstruction[6]));
        }

        void process(final Map<String, RegisterValue> registers) {
            final RegisterValue registerValue = getRegister(registers);
            registerValue.operate(operator, amount);
        }

        RegisterValue getRegister(final Map<String, RegisterValue> registers) {
            RegisterValue registerValue = registers.get(register);
            if (null == registerValue) {
                registerValue = new RegisterValue();
                registers.put(register, registerValue);
            }
            return registerValue;
        }
    }

    class RegisterValue {
        private int value;
        private int maxValue;

        void operate(final Operator operator, final int amount) {
            switch (operator) {
                case DEC:
                    value -= amount;
                    break;
                case INC:
                    value += amount;
                    break;
                default:
                    throw new RuntimeException("Unhandled operator " + operator);
            }
            if (value > maxValue) {
                maxValue = value;
            }
        }

        int getValue() {
            return value;
        }

        int getMaxValue() {
            return maxValue;
        }
    }
}
