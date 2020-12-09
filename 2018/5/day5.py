def get_polymer(filename):
    file = open(filename)
    polymer = file.readline()
    return polymer.strip()


def react(prev_unit, current_unit):
    if prev_unit == current_unit:
        return False
    if prev_unit.lower() != current_unit.lower():
        return False
    return True


def list_to_string_skip_spaces(polymer):
    s = ""
    for unit in polymer:
        if unit != " ":
            s += unit
    return s


def process_polymer(raw_polymer):
    prev_unit = 0
    unit = 1
    reacted = False
    polymer = list(raw_polymer)
    while unit < len(polymer):
        if react(polymer[prev_unit], polymer[unit]):
            reacted = True
            polymer[prev_unit] = " "
            polymer[unit] = " "
        prev_unit = unit
        unit += 1
    str_polymer = list_to_string_skip_spaces(polymer)
    return not reacted, str_polymer


def part1(polymer):
    # print("Starting length: ", len(polymer))
    finished = False
    while not finished:
        finished, polymer = process_polymer(polymer)
    # print(polymer)
    return len(polymer)


def part1_test(polymer, expected_result):
    result = part1(polymer)
    print(polymer, ": ", result == expected_result)


def test():
    part1_test("aA", 0)
    part1_test("abBA", 0)
    part1_test("abAB", 4)
    part1_test("aabAAB", 6)
    part1_test("dabAcCaCBAcCcaDA", 10)
    part1_test("aAxxxbB", 3)
    part1_test("aAxXbBc", 1)
    part1_test("aXxdEeD", 1)


def part2(polymer):
    shortest_polymer = len(polymer)
    for unit_to_remove in list("abcdefghijklmnopqrstuvwxyz"):
        print("Reacting without ", unit_to_remove)
        reduced_polymer = polymer.replace(unit_to_remove, "").replace(unit_to_remove.upper(), "")
        reacted_polymer_length = part1(reduced_polymer)
        if reacted_polymer_length < shortest_polymer:
            shortest_polymer = reacted_polymer_length
    print(shortest_polymer)


# print(part1(get_polymer("input.txt")))
# test()
part2(get_polymer("input.txt"))
# part2("dabAcCaCBAcCcaDA")
