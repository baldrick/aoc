import aoc


class Coord:
    def __init__(self, row, col):
        self.row = row
        self.col = col

    def __repr__(self):
        return "r%s,c%s" % (self.row, self.col)

    def __hash__(self):
        return hash((self.row, self.col))

    def __eq__(self, other):
        return (self.row, self.col) == (other.row, other.col)


class Basin:
    def __init__(self, coord):
        self.coords = {coord: coord}
        self.name = coord

    def add(self, coord):
        self.coords[coord] = coord

    def size(self):
        return len(self.coords)

    def __repr__(self):
        return "%s: %s" % (self.name, self.coords)

    def contains(self, coord):
        return coord in self.coords


class Row:
    def __init__(self, line):
        self.row = []
        for c in line:
            if c == "\n":
                continue
            self.row.append(int(c))

    def __repr__(self):
        return "".join(str(n) for n in self.row)

    def len(self):
        return len(self.row)

    def get(self, col):
        return self.row[col]


class Grid:
    def __init__(self, lines):
        self.grid = []
        for line in lines:
            self.grid.append(Row(line))

    def __repr__(self):
        s = ""
        for row in self.grid:
            s += "%s\n" % row
        return s

    def highlight(self, basin):
        y = 0
        for row in self.grid:
            x = 0
            for col in row.row:
                if (basin.contains(Coord(y, x))):
                    print(f"{aoc.colors.RED}{self.grid[y].get(x)}{aoc.colors.ENDC}", end='')
                else:
                    print(f"{self.grid[y].get(x)}", end='')
                x += 1
            y += 1
            print()

    def highlightAll(self, basins):
        colours = [aoc.colors.RED, aoc.colors.BLUE, aoc.colors.GREEN]
        y = 0
        for row in self.grid:
            x = 0
            for col in row.row:
                n = 0
                found = False
                for basin in basins:
                    if (basin.contains(Coord(y, x))):
                        print(f"{colours[n]}{self.grid[y].get(x)}{aoc.colors.ENDC}", end='')
                        found = True
                        break
                    n += 1
                if not found:
                    print(f"{self.grid[y].get(x)}", end='')
                x += 1
            y += 1
            print()

    def diagonal(self, dc, dr):
        return dc != 0 and dr != 0

    def isLowPoint(self, row, col):
        rows = len(self.grid)
        cols = self.grid[0].len()
        dr = -1
        while dr <= 1:
            dc = -1
            while dc <= 1:
                if row+dr < 0 or col+dc < 0 or row+dr >= rows or col+dc >= cols or (dc == 0 and dr == 0) or self.diagonal(dc, dr):
                    dc += 1
                    continue
                if self.grid[row].get(col) >= self.grid[row+dr].get(col+dc):
                    # print("r", row, ", c", col, "(", self.grid[row].get(col), ") > r", row+dr, ", c", col+dc, "(", self.grid[row+dr].get(col+dc), ")")
                    return False
                dc += 1
            dr += 1
        # print("LOWPOINT r", row, ", c", col, "(", self.grid[row].get(col), ")")
        return True

    def lowPoints(self):
        lp = []
        row = 0
        while row < len(self.grid):
            col = 0
            while col < self.grid[row].len():
                if self.isLowPoint(row, col):
                    lp.append(Coord(row, col))
                col += 1
            row += 1
        return lp

    def findBasin(self, p):
        return self.addAdjacentHighPoints(Basin(p), p)

    def inGrid(self, p, dx, dy):
        return p.row + dy >= 0 and p.row + dy < len(self.grid) and p.col + dx >= 0 and p.col + dx < self.grid[0].len()

    def addAdjacentHighPoints(self, basin, p):
        dx = -1
        while dx <= 1:
            dy = -1
            while dy <= 1:
                if (dx, dy) != (0, 0) and not self.diagonal(dx, dy) and self.inGrid(p, dx, dy):
                    hp = self.grid[p.row + dy].get(p.col + dx)
                    if (self.grid[p.row].get(p.col) < hp and hp != 9):
                        c = Coord(p.row + dy, p.col + dx)
                        #print("Adding", c, "to basin", basin.name)
                        basin.add(c)
                        self.addAdjacentHighPoints(basin, c)
                dy += 1
            dx += 1
        return basin


lines = aoc.getInput()

grid = Grid(lines)

# print(grid)
lp = grid.lowPoints()
sum = 0
basins = []
for p in lp:
    basins.append(grid.findBasin(p))

basins.sort(key=lambda x: x.size(), reverse=True)
product = 1
print(grid.highlightAll(basins[0:3:1]))
for basin in basins[0:3:1]:
    # print(f"{aoc.colors.GREEN}size={basin.size()}{aoc.colors.ENDC}")
    # print(grid.highlight(basin))
    product *= basin.size()
print(product)
