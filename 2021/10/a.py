import aoc

lines = aoc.getInput()

def bracket(c):
    return c in ["(", "[", "{", "<", ")", "]", "}", ">"]

def opener(c):
    return c in ["(", "[", "{", "<"]

def closer(c):
    match c:
        case "(": return ")"
        case "[": return "]"
        case "{": return "}"
        case "<": return ">"
    print("Unexpected closer", c)
    return ""

def checkLine(line, illegal):
    stack = []
    for c in line:
        if not bracket(c):
            continue
        if opener(c):
            stack.append(c)
        else:
            o = stack.pop()
            if closer(o) != c:
                print(line, "is corrupt, got", c, "expected", closer(o))
                illegal[c] += 1

illegal = { ")": 0, "]": 0, "}": 0, ">": 0 }
for line in lines:
    checkLine(line, illegal)

sum = illegal[")"] * 3
sum += illegal["]"] * 57
sum += illegal["}"] * 1197
sum += illegal[">"] * 25137

print(sum)
