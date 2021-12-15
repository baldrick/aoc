import aoc
import sys

lines = aoc.getInput()

polymer = lines[0]

pairs = {}
for i in range(2, len(lines)):
    pair = lines[i].split(' -> ')
    pairs[pair[0]] = pair[1]

# map pair+depth to distribution of letters -> counts
pd_dist = {}

def count(s):
    counts = {}
    for i in range(0, len(s)-1):
        c = s[i]
        if c in counts:
            counts[c] += 1
        else:
            counts[c] = 1
    return counts

# Merge two dictionaries of letters -> counts
def merge(a, b):
    for k in b:
        if k in a:
            a[k] += b[k]
        else:
            a[k] = b[k]
    return a

def dist(distributions, pairs, polymer, depth):
    if f"{polymer}{depth}" in distributions:
        return distributions[f"{polymer}{depth}"]
    if depth == 0:
        print(f"depth:0, polymer:{polymer}")
        counts = count(polymer)
        distributions[f"{polymer}0"] = counts
        return counts
    print(f"depth:{depth}, polymer:{polymer}")
    d = {}
    for i in range(0, len(polymer)-1):
        pair = f"{polymer[i]}{polymer[i+1]}"
        insert = pairs[pair]
        merge(d, dist(distributions, pairs, f"{polymer[i]}{insert}{polymer[i+1]}", depth - 1))
        #merge(d, dist(distributions, pairs, f"{insert}{polymer[i+1]}", depth - 1))
    distributions[f"{polymer}{depth}"] = d
    return d


print(f"1: {count('NCNBCHB')}")
print(f"2: {count('NBCCNBBBCBHCB')}")
print(f"3: {count('NBBBCNCCNBBNBNBBCHBHHBCHB')}")
print(f"4: {count('NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB')}")

d = dist({}, pairs, polymer, int(sys.argv[2]))
# add one for the final letter of the polymer
d[polymer[len(polymer)-1]] += 1
print(d)

min = 9999999999999999
max = 0
for c in d:
    if d[c] < min:
        min = d[c]
    if d[c] > max:
        max = d[c]

print(max - min)
