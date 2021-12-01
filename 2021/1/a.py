import sys

filename = sys.argv[1]
with open(filename) as f:
    lines = f.readlines()

count = -1
prev = 0
for n in lines:
    current = int(n)
    if current > prev:
        count += 1
    prev = current
print(count)
