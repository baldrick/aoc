import aoc

lines = aoc.getInput()

digits = [0,0,0,0,0,0,0,0,0]

for line in lines:
    halves = line.split("|")
    input = halves[0]
    output = halves[1]
    for digit in output.split():
        ld = len(digit)
        if ld == 2 or ld == 4 or ld == 3 or ld == 7:
            print(digit, "has length", ld)
        digits[ld] += 1

print("1 appears", digits[2], "times")
print("4 appears", digits[4], "times")
print("7 appears", digits[3], "times")
print("8 appears", digits[7], "times")
print("total:", digits[2] + digits[4] + digits[3] + digits[7])
