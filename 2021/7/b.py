import aoc

lines = aoc.getInput()

positions = [0]

max = 0
for pos in lines[0].split(','):
    if (int(pos) > max):
        max = int(pos)

sumOneTo = []
sum = 0
n = 0
while n <= max:
    sum += n
    sumOneTo.append(sum)
    #print("sum 0..", n, " = ", sum)
    n += 1

while max > 0:
    positions.append(0)
    max -= 1

for pos in lines[0].split(','):
    positions[int(pos)] += 1

def fuelToMoveTo(targetPosition, currentMinFuel):
    fuel = 0
    n = 0
    while fuel < currentMinFuel and n < len(positions):
        addFuel = positions[n] * sumOneTo[abs(n - targetPosition)]
        fuel += addFuel
    #    print("Adding", addFuel, "to move", positions[n], "crabs to", abs(targetPosition))
        n += 1
    #if fuel >= currentMinFuel:
    #    print("No improvement targeting position", targetPosition, "stopped at", n)
    return fuel

def minFuel():
    targetPosition = 0
    minFuel = 1e12
    while targetPosition < len(positions):
        fuel = fuelToMoveTo(targetPosition, minFuel)
        if fuel < minFuel:
            minFuel = fuel
            #print("minFuel=", minFuel, " at target ", targetPosition)
        targetPosition += 1
    return minFuel

print(minFuel())
