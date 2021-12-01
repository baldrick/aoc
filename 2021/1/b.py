import sys

filename = sys.argv[1]
with open(filename) as f:
    lines = f.readlines()

def three_sum(lines, start):
    return int(lines[start]) + int(lines[start+1]) + int(lines[start+2])

i = 0
count = 0
while i < len(lines)-3:
    if three_sum(lines, i) < three_sum(lines, i+1):
        count += 1
    i += 1

print(count)
