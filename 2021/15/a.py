import aoc
import sys

lines = aoc.getInput()

class ChitonGrid(aoc.Grid):
    safestByCell = {}
    def __init__(self, lines):
        super().__init__(lines)
        lastCell = aoc.Coord(len(self.grid)-1, len(self.grid[0].row)-1)
        self.safestByCell[lastCell] = self.cell(lastCell)

    # Calculate risk from the last cell back to the first.
    # Calculate risk for bottom row and last column first.
    # Then move on to the next row / column in.
    # This guarantees the cached risk will be the safest...
    def calculateSafestByCell(self, row, col):
        if row < 0 or col < 0:
            return

        p = aoc.Coord(row, col)
        if p not in self.safestByCell:
            self.safestByCell[p] = self.cell(p) + self.chooseOneFrom(p)

        #print(f"calculating safest route from cell {p} ({self.cell(p)})")
        for r in range(row-1, -1, -1):
            target = aoc.Coord(r, col)
            cellRisk = self.cell(target)
            extraRisk = self.chooseOneFrom(target)
            self.safestByCell[target] = cellRisk + extraRisk
            #print(f"calculated safest route from {target} ({self.cell(target)}) = {cellRisk} + {extraRisk} = {self.safestByCell[target]}")

        for c in range(col-1, -1, -1):
            target = aoc.Coord(row, c)
            cellRisk = self.cell(target)
            extraRisk = self.chooseOneFrom(target)
            self.safestByCell[target] = cellRisk + extraRisk
            #print(f"calculated safest route from {target} ({self.cell(target)}) = {cellRisk} + {extraRisk} = {self.safestByCell[target]}")

        self.calculateSafestByCell(row - 1, col - 1)

    def chooseOneFrom(self, p):
        right = aoc.Coord(p.row, p.col+1)
        down = aoc.Coord(p.row+1, p.col)
        if self.inGrid(right, 0, 0) and self.inGrid(down, 0, 0):
            #print(f"map: {self.safestByCell}")
            #print(f"right: {right} = {self.safestByCell[right]}")
            #print(f"down: {down} = {self.safestByCell[down]}")
            return min(self.safestByCell[right], self.safestByCell[down])
        elif self.inGrid(right, 0, 0):
            #print(f"returning right {right} = {self.safestByCell[right]}")
            return self.safestByCell[right]
        elif self.inGrid(down, 0, 0):
            #print(f"returning down {down} = {self.safestByCell[down]}")
            return self.safestByCell[down]

    def scale(self):
        self.scaleWidth()
        self.scaleHeight()
        self.safestByCell = {}
        lastCell = aoc.Coord(len(self.grid)-1, len(self.grid[0].row)-1)
        self.safestByCell[lastCell] = self.cell(lastCell)

    def scaleWidth(self):
        n = 0
        for row in self.grid:
            newRow = row.row.copy()
            for m in range(1, 5):
                for col in row.row:
                    x = col + m
                    if x > 9:
                        x -= 9
                    newRow.append(x)
            self.grid[n].row = newRow
            n += 1
        
    def scaleHeight(self):
        newGrid = self.grid.copy()
        for m in range(1, 5):
            for row in self.grid:
                newRow = aoc.Row("")
                for col in row.row:
                    x = col + m
                    if x > 9:
                        x -= 9
                    newRow.row.append(x)
                newGrid.append(newRow)
        
        self.grid = newGrid
                    
    def cheapestPath(self, path, next):
        if next is None:
            return path
        path.append(next)
        return self.cheapestPath(path, self.cheapestCellFrom(next))
    
    def cheapestCellFrom(self, p):
        if self.isLastCell(p):
            return None
        right = aoc.Coord(p.row, p.col+1)
        down = aoc.Coord(p.row+1, p.col)
        if self.inGrid(right, 0, 0) and self.inGrid(down, 0, 0):
            if self.safestByCell[right] < self.safestByCell[down]:
                return right
            return down
        if self.inGrid(right, 0, 0):
            return right
        return down
    
    def highlight(self, p):
        y = 0
        for row in self.grid:
            x = 0
            for col in row.row:
                if aoc.Coord(y, x) in p:
                    print(f"{aoc.colors.RED}{self.grid[y].get(x)}{aoc.colors.ENDC}", end='')
                else:
                    print(f"{self.grid[y].get(x)}", end='')
                x += 1
            y += 1
            print()        


g = ChitonGrid(lines)
#g.scale()

print(f"{len(g.grid)}x{len(g.grid[0].row)} grid:")
print(f"{g}")
g.calculateSafestByCell(len(g.grid)-1, len(g.grid[0].row)-1)
start = aoc.Coord(0,0)
g.safestByCell[start] = g.chooseOneFrom(start)
print(g.safestByCell[start])

#sys.setrecursionlimit(1500)
#path = g.cheapestPath([], start)
#g.highlight(path)
