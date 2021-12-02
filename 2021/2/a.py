import sys

filename = sys.argv[1]
with open(filename) as f:
    lines = f.readlines()

depth=0
x=0
for line in lines:
    instruction = line.split()
    match instruction[0]:
        case 'forward':
            x += int(instruction[1])
        case 'down':
            depth += int(instruction[1])
        case 'up':
            depth -= int(instruction[1])

print(x*depth)
