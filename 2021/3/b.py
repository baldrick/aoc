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

def filter(bitPos, numbers, includeBit):
    filtered = []
    for n in numbers:
        if n[bitPos] == includeBit:
            filtered.append(n)
    return filtered

'''
To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position,
and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with
a 1 in the position being considered.
'''
def mostCommon(bitPos, numbers):
    countOn = 0
    for n in numbers:
        bits = split(n)
        if bits[bitPos] == "1":
            countOn += 1
    if countOn >= len(numbers)/2:
        return filter(bitPos, numbers, "1")
    else:
        return filter(bitPos, numbers, "0")

def leastCommon(bitPos, numbers):
    countOn = 0
    for n in numbers:
        bits = split(n)
        if bits[bitPos] == "1":
            countOn += 1
    if countOn >= len(numbers)/2:
        return filter(bitPos, numbers, "0")
    else:
        return filter(bitPos, numbers, "1")

def binStringToNumber(s, bitCount):
    bit = 0
    n = 0
    while bit < bitCount:
        if s[bit] == "1":
            n |= 1 << bitCount - bit - 1
        bit += 1
    return n

numbers = lines
bitPos = 0
while len(numbers) > 1:
    numbers = mostCommon(bitPos, numbers)
    bitPos += 1
    print(numbers)

o2 = binStringToNumber(numbers[0], bitCount)
print("o2=", o2)

numbers = lines
bitPos = 0
while len(numbers) > 1:
    numbers = leastCommon(bitPos, numbers)
    bitPos += 1
    print(numbers)

co2 = binStringToNumber(numbers[0], bitCount)
print("co2=", co2)

print("life support = ", o2 * co2)