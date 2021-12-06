import aoc

def getXY(s):
    coord = s.split(",")
    x = int(coord[0])
    y = int(coord[1])
    return x, y


class Coord:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def __repr__(self):
        return "%s,%s" % (self.x, self.y)

class Line:
    def __init__(self, start, end):
        self.start = start
        self.end = end
    
    def __repr__(self):
        return "%s,%s to %s,%s" % (self.start.x, self.start.y, self.end.x, self.end.y)

    def isDiagonal(self):
        return self.start.x != self.end.x and self.start.y != self.end.y

    def ignore(self):
        if self.start.x == self.end.x or self.start.y == self.end.y:
            return False
        return abs(self.start.x - self.end.x) != abs(self.start.y - self.end.y)
    def maxX(self):
        return max(self.start.x, self.end.x)
    
    def maxY(self):
        return max(self.start.y, self.end.y)

    def dx(self):
        if self.start.x > self.end.x:
            return -1
        if self.start.x < self.end.x:
            return 1
        return 0

    def dy(self):
        if self.start.y > self.end.y:
            return -1
        if self.start.y < self.end.y:
            return 1
        return 0

    def getFillCoords(self):
        dx = self.dx()
        dy = self.dy()
        x = self.start.x
        y = self.start.y
        coords = []
        while not self.finished(x, y, dx, dy):
            coords.append(Coord(x,y))
            x += dx
            y += dy
        return coords
    
    def finished(self, x, y, dx, dy):
        if dx < 0 and x < self.end.x:
            #print("Finished: dx=", dx, " and x<self.end.x:", x, self.end.x)
            return True
        if dx > 0 and x > self.end.x:
            #print("Finished: dx=", dx, " and x>self.end.x:", x, self.end.x)
            return True
        if dy < 0 and y < self.end.y:
            #print("Finished: dy=", dy, " and y<self.end.y:", y, self.end.y)
            return True
        if dy > 0 and y > self.end.y:
            #print("Finished: dy=", dy, " and y>self.end.y:", y, self.end.y)
            return True
        return False

class Map:
    def __init__(self, lines):
        self.mapLines = []
        maxX = 0
        maxY = 0
        for line in lines:
            split = line.split("->")
            sx, sy = getXY(split[0])
            ex, ey = getXY(split[1])
            mapLine = Line(Coord(sx,sy), Coord(ex,ey))
            self.mapLines.append(mapLine)
            if maxX < mapLine.maxX():
                maxX = mapLine.maxX()
            if maxY < mapLine.maxY():
                maxY = mapLine.maxY()

        self.rows = []
        y = 0
        while y < maxY+1:
            row = []
            x = 0
            while x < maxX+1:
                row.append(0)
                x += 1
            self.rows.append(row)
            y += 1

        for mapLine in self.mapLines:
            if mapLine.ignore():
                print("Ignoring", mapLine, " - it's diagonal but not 45 degrees")
                continue
            #print("Filling", mapLine, "with", mapLine.getFillCoords())
            for coord in mapLine.getFillCoords():
                self.rows[coord.y][coord.x] += 1
    
    def __repr__(self):
        rep = ""
        for row in self.rows:
            s = ""
            for cell in row:
                if cell == 0:
                    c = "."
                else:
                    c = cell
                s = "%s%s" % (s, c)
            rep = "%s\n%s" % (rep, s)
        return rep

    def doubleMarked(self):
        count = 0
        nrow = -1
        for row in self.rows:
            nrow += 1
            ncol = -1
            for col in row:
                ncol += 1
                if (col > 1):
                    count += 1
        return count

lines = aoc.getInput()

map = []
n = 0
map = Map(lines)
print(map)
count = map.doubleMarked()
print(count)
