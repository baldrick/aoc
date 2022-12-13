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

    def size(self):
        return self.max - self.min

    def overlap(self, other):
        # a 2-3, 1-2, 3-4
        # count overlap with
        # b 3-4, 2-6, 4-10

        # a, b overlap 3-3, 2-2, 4-4
        # a, b x overlaps from max(a.minx, b.minx) to min(a.maxx, b.maxx)
        min = max(self.min, other.min)
        max = min(self.max, other.max)
        return Range(min=min, max=max)

class Instruction:
    def __init__(self, input, r):
        cmdCoord = input.split(" ")
        self.cmd = cmdCoord[0]
        coord = cmdCoord[1]
        xyz = coord.split(",")
        self.xr = Range(xyz[0])
        self.yr = Range(xyz[1])
        self.zr = Range(xyz[2])
        self.value = 0
        if self.cmd == "on":
            self.value = 1
    
    def __repr__(self):
        cells = self.xr.size() * self.yr.size() * self.zr.size()
        return f"{self.cmd}: x={self.xr}, y={self.yr}, z={self.zr} ({cells} cells)"

class Cube:
    def __init__(self, xr, yr, zr, value):
        self.xr = xr
        self.yr = yr
        self.zr = zr
        self.value = value
    
    def __repr__(self):
        return f"x={self.xr}, y={self.yr}, z={self.zr} --> {self.value}"
    
    def size(self):
        return self.xr.size() * self.yr.size() * self.zr.size()

    def overlap(self, other):
        xoverlap = self.xr.overlap(other.xr)
        yoverlap = self.yr.overlap(other.yr)
        zoverlap = self.zr.overlap(other.zr)
        return Cube(overlap, yoverlap, zoverlap)

cubes = []
input = util.getInput()
for line in input:
    i = Instruction(line)
    cubes.append(Cube(i.xr, i.yr, i.zr, i.value))

size A
+ size B - overlap(A, B)
+ size C - overlap(a,c) - overlap(b,c) + overlap(a,b,c)
+ size D - overlap(a,d) - overlap(b,d) - overlap(c,d) + overlap(a,b,c,d)

1
1 1   1
1 1 1 1
  1 1 1
  1 1
    1
    1

3
+ 4 - 2
+ 5 - 1 - 3 + 1


1 1 1

size a
+ size b - overlap a,b
+ size c - overlap 

split regions?

a1
a2 b1
a3 b2 c1
   b3 c2
      c3

region A is kept
region B splits into a smaller region (b3) and the overlap discarded
regions A and (now smaller) B are kept, the overlap discarded so region C becomes c3

how to do this efficiently in 3D? each overlap could create 3 regions (think corner of a cube missing)
A x1-3, y1-4, z1-5
B x1-2, y4-5, z5-6
C x3, y4-5, z6
D x2, y2, z2
E x2, y2, z5-6

A,B overlap is x1-2,y4,z5
B becomes non-overlapping parts x3,y5,z6

A,C overlap is nothing (z takes it outside A)
B,C overlap is nothing (x takes it outside B)

A,D overlap is x2, y2, z2
D is entirely within A so it disappears

A,E overlap is x2, y2, z5
E becomes non-overlapping region x2,y2,z6