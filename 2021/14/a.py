import aoc

lines = aoc.getInput()

polymer = lines[0]
pairs = {}
for i in range(2, len(lines)):
    pair = lines[i].split(' -> ')
    pairs[pair[0]] = pair[1]

def polymerize(polymer, pairs):
    result = ""
    for i in range(0, len(polymer)-1):
        pair = f"{polymer[i]}{polymer[i+1]}"
        insert = pairs[pair]
        result += polymer[i] + insert
    result += polymer[len(polymer)-1]
    return result

for i in range(1,11):
    polymer = polymerize(polymer, pairs)
    print(f"After step {i} polymer has length {len(polymer)}", flush=True)

counts = {}
for c in polymer:
    if c in counts:
        counts[c] += 1
    else:
        counts[c] = 1

min = 9999999999999999
max = 0
for c in counts:
    if counts[c] < min:
        min = counts[c]
    if counts[c] > max:
        max = counts[c]

print(max - min)