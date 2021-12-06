import aoc
import sys

lines = aoc.getInput()

days = int(sys.argv[2])

dayCount = [0,0,0,0,0,0,0,0,0]

for day in lines[0].split(","):
    dayCount[int(day)] += 1

print(dayCount)

day = 0
while day < days:
    newFish = dayCount[0]
    n = 1
    while n < len(dayCount):
        dayCount[n-1] += dayCount[n]
        dayCount[n] = 0
        n += 1
    dayCount[8] += newFish
    dayCount[6] += newFish
    dayCount[0] -= newFish
    print(dayCount)
    day += 1

sum = 0
for dc in dayCount:
    sum += dc
print(sum)