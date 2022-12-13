# on x=-28..25,y=-34..15,z=-36..13
import util

class Range:
    def __init__(self, xyz=None, min=None, max=None):
        if xyz is None:
            self.min = min
            self.max = max
        else:
            minMax = xyz.split("=")[1].split("..")
            self.min = int(minMax[0])
            self.max = int(minMax[1])

    def __repr__(self):
        return f"{self.min} to {self.max}"

    def inside(self, r):
        return self.min >= r.min and self.max <= r.max

    def size(self):
        return self.max - self.min

class Instruction:
    def __init__(self, input, r):
        cmdCoord = input.split(" ")
        self.cmd = cmdCoord[0]
        coord = cmdCoord[1]
        xyz = coord.split(",")
        self.xr = Range(xyz[0])
        self.yr = Range(xyz[1])
        self.zr = Range(xyz[2])
        self.valid = True
        if not self.xr.inside(r) and not self.yr.inside(r) and not self.zr.inside(r):
            self.valid = False
    
    def __repr__(self):
        cells = self.xr.size() * self.yr.size() * self.zr.size()
        return f"{self.cmd}: x={self.xr}, y={self.yr}, z={self.zr} ({cells} cells)"

class Cube:
    def __init__(self, r):
        self.xr = Range(min=r.min, max=r.max)
        self.yr = Range(min=r.min, max=r.max)
        self.zr = Range(min=r.min, max=r.max)
        self.cube = {}
    
    def __repr__(self):
        return f"x={self.xr}, y={self.yr}, z={self.zr}"

    def apply(self, i):
        print(f"turning {i}", flush=True)
        match i.cmd:
            case "on":
                for x in range(i.xr.min, i.xr.max+1):
                    for y in range(i.yr.min, i.yr.max+1):
                        for z in range(i.zr.min, i.zr.max+1):
                            self.cube[(x,y,z)] = 1
            case "off":
                for x in range(i.xr.min, i.xr.max+1):
                    for y in range(i.yr.min, i.yr.max+1):
                        for z in range(i.zr.min, i.zr.max+1):
                            self.cube[(x,y,z)] = 0
            case _:
                print(f"error {i}!")

input = util.getInput()

r=Range(min=-50, max=50)
c = Cube(r)

for line in input:
    i = Instruction(line, r)
    if i.valid:
        c.apply(i)

count = 0
for coord, value in c.cube.items():
    if value == 1:
        count += 1
print(count)