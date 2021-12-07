import aoc

lines = aoc.getInput()

positions = [0]

max = 0
for pos in lines[0].split(','):
    if (int(pos) > max):
        max = int(pos)

while max > 0:
    positions.append(0)
    max -= 1

for pos in lines[0].split(','):
    positions[int(pos)] += 1

targetPosition = 0
minFuel = 1e12

while targetPosition < len(positions):
    fuel = 0
    n = 0
    while fuel < minFuel and n < len(positions):
        fuel += positions[n] * abs(n - targetPosition)
        n += 1
    if fuel < minFuel:
        minFuel = fuel
    targetPosition += 1

print(minFuel)