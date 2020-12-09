def part1():
    total = 0
    input_file = open("input.txt", "r")
    for line in input_file:
        total += int(line)
    print total


def part2():
    total = 0
    frequency_seen = { total }
    finished = False
    while not finished:
        print "Reading file"
        input_file = open("input.txt", "r")
        for line in input_file:
            total += int(line)
            if total in frequency_seen:
                print total
                finished = True
                break
            frequency_seen.add(total)


part2()
