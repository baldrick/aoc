import sys

filename = sys.argv[1]
with open(filename) as f:
    lines = f.readlines()

countOfOnBits = [0,0,0,0,0,0,0,0,0,0,0,0,0]

def split(word):
    return list(word)

for line in lines:
    bits = split(line)
    pos = 0
    bitCount = 0
    for bit in bits:
        if bit == "1":
            countOfOnBits[pos] += 1
        if bit != "0" and bit != "1":
            continue
        pos += 1
        bitCount += 1

print("bitCount = ", bitCount)
print("countOfOnBits = ", countOfOnBits)
print("len lines = ", len(lines))
bit = 0
gamma = 0
epsilon = 0
while bit < bitCount:
    if len(lines)/2 < countOfOnBits[bit]:
        print("bit ", bitCount - bit, " should be set")
        gamma |= 1 << bitCount - bit - 1
    else:
        epsilon |= 1 << bitCount - bit - 1
    bit += 1

print("gamma = ", gamma)
print("epsilon = ", epsilon)
print(epsilon * gamma)
