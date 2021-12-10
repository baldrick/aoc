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

def checkLine(line):
    stack = []
    for c in line:
        if not bracket(c):
            continue
        if opener(c):
            stack.append(c)
        else:
            o = stack.pop()
            if closer(o) != c:
                #print(line, "is corrupt, got", c, "expected", closer(o))
                return None # means the line is corrupt
    return stack

def score(c):
    match c:
        case ")": return 1
        case "]": return 2
        case "}": return 3
        case ">": return 4
    print("Unexpected closer, cannot score", c)
    return 0

def scoreCompletion(stack):
    s = 0
    while len(stack) > 0:
        o = stack.pop()
        c = closer(o)
        s *= 5
        s += score(c)
    return s

scores = []
for line in lines:
    stack = checkLine(line)
    if stack != None:
        scores.append(scoreCompletion(stack))

scores.sort()
middle = int((len(scores)-1)/2)
print(scores[middle])