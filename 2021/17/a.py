import aoc


class Target:
    def __init__(self, input):
        xy = input[0].split()
        (self.xrmin, self.xrmax) = self.getRange(xy[2])
        (self.yrmax, self.yrmin) = self.getRange(xy[3])  # min/max swapped deliberately!
        aoc.log(1, f"x: {self.xrmin} to {self.xrmax}, y: {self.yrmin} to {self.yrmax}")

    def getRange(self, e):
        r = e.split('=')
        minmax = r[1].split('..')
        rmin = int(minmax[0])
        rmax = int(minmax[1].strip(','))
        return (rmin, rmax)

    def hit(self, xv, yv):
        step = 0
        x = 0
        y = 0
        while True:
            step += 1
            x += xv
            y += yv
            yv -= 1
            xv = max(xv - 1, 0)
            aoc.log(2, f"Step {step} at {x},{y} speed {xv},{yv}")
            if xv == 0 and x < self.xrmin:
                # Won't reach target
                aoc.log(2, f"Won't reach target, xv=0, x={x}")
                return False
            if x >= self.xrmin and x <= self.xrmax and y >= self.yrmax and y <= self.yrmin:
                aoc.log(2, f"Hit target at {x},{y} on step {step}")
                return True
            if x > self.xrmax or y < self.yrmax:
                # Gone past target
                aoc.log(2, f"Gone past target, x={x}, y={y} after {step} steps")
                return False


def getStepsTo(x):
    sum = 0
    n = 0
    while sum < x:
        n += 1
        sum += n
    return n


def sumTo(x):
    sum = 0
    for n in range(0, x+1):
        sum += n
    return sum


# e.g. target area: x=20..30, y=-10..-5
input = aoc.getInput()
t = Target(input)

# minx = sum(0..start x) >= range min x
minx = getStepsTo(t.xrmin)
aoc.log(2, f"Min x speed: {minx}")

# find max y start speed such that we hit the target
# speed could still be > target size if we get "lucky"
# let's just try a "sensible" range
speedSet = aoc.SetOfCoords()
maxy = 0
for y in range(0, 20):
    if t.hit(minx, y):
        speedSet.add(aoc.Coord(minx, y))
        aoc.log(1, f"target hit when yspeed = {y}")
        if y > maxy:
            maxy = y

h = sumTo(maxy)
aoc.log(1, f"max height {h} reached with min starting speed {maxy}")

# to get greater height, num steps must be > maxy
maxy2 = maxy
for x in range(minx, minx+100):
    aoc.log(1, f"trying x={x}")
    for y in range(maxy, maxy+1000):
        if t.hit(x, y):
            speedSet.add(aoc.Coord(x, y))
            if y > maxy2:
                aoc.log(1, f"target hit when yspeed = {y}, new max")
                maxy2 = y

# part B - find set of all starting velocities that hit the target
# means we need to look at starting y speeds < maxy
for y in range(maxy, t.yrmax-1, -1):
    aoc.log(1, f"trying y={y}")
    for x in range(minx, minx+200):
        if t.hit(x, y):
            aoc.log(1, f"hit with starting speed {x},{y}")
            speedSet.add(aoc.Coord(x, y))

h = sumTo(maxy2)
aoc.log(1, f"max height {h} reached, {speedSet.size()} starting velocities hit the target")
# maxx = sum(end step..start x) <= range max x
