import aoc
import sys


class FoldingGrid():
    def __init__(self, lines):
        self.folds = []
        self.coords = aoc.SetOfCoords("start")
        folds = False
        for line in lines:
            if folds:
                self.folds.append(line)
            else:
                if line == "":
                    folds = True
                    continue
                xy = line.split(',')
                x = int(xy[0])
                y = int(xy[1])
                self.coords.add(aoc.Coord(y, x))

    def __repr__(self):
        xmax = 0
        ymax = 0
        for c in self.coords.coords:
            if c.row > ymax:
                ymax = c.row
            if c.col > xmax:
                xmax = c.col
        s = f"{self.coords.name}\n"
        for y in range(0, ymax+1):
            s += f"{y}: "
            for x in range(0, xmax+1):
                if self.coords.contains(aoc.Coord(y, x)):
                    s += '#'
                else:
                    s += ' '
            s += '\n'
        s += f"Num dots: {self.coords.size()}\n"
        return s

    def fold(self, index):
        f = self.folds[index].split()
        c = f[2].split('=')
        axis = c[0]
        xy = int(c[1])
        foldedCoords = aoc.SetOfCoords(f"After fold {index} - {self.folds[index]}")
        if axis == 'x':
            #print(f"fold on x at {xy}")
            for coord in self.coords.coords:
                if coord.col == xy:
                    print(f"Error, folding along a line with a #: x={xy}")
                    return
                if coord.col < xy:
                    foldedCoords.add(coord)
                else:
                    #print(f"folding {coord.xy()} at x{xy} to {xy - (coord.col - xy)}, {coord.row}")
                    foldedCoords.add(aoc.Coord(coord.row, xy - (coord.col - xy)))
        else:
            #print(f"fold on y at {xy}")
            for coord in self.coords.coords:
                if coord.row == xy:
                    print(f"Error, folding along a line with a #: y={xy}")
                    return
                if coord.row < xy:
                    foldedCoords.add(coord)
                else:
                    #print(f"folding {coord.xy()} at y{xy} to {coord.col}, {xy - (coord.row - xy)}")
                    foldedCoords.add(aoc.Coord(xy - (coord.row - xy), coord.col))
        self.coords = foldedCoords


lines = aoc.getInput()

grid = FoldingGrid(lines)

for fold in range(0, len(grid.folds)):
    grid.fold(fold)

print(grid)
