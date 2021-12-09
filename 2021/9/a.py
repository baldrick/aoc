import aoc

class Coord:
    def __init__(self, row, col):
        self.row = row
        self.col = col
    
    def __repr__(self):
        return "r%s,c%s" % (self.row, self.col)

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
                    #print("r", row, ", c", col, "(", self.grid[row].get(col), ") > r", row+dr, ", c", col+dc, "(", self.grid[row+dr].get(col+dc), ")")
                    return False
                dc += 1
            dr += 1
        #print("LOWPOINT r", row, ", c", col, "(", self.grid[row].get(col), ")")
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
    
lines = aoc.getInput()

grid = Grid(lines)

print(grid)
lp = grid.lowPoints()
sum = 0
for p in lp:
    print(p, "=", grid.grid[p.row].get(p.col))
    sum += grid.grid[p.row].get(p.col)
print(sum + len(lp))