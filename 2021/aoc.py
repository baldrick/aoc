import sys


class colors:  # You may need to change color settings
    RED = '\033[31m'
    ENDC = '\033[m'
    GREEN = '\033[32m'
    YELLOW = '\033[33m'
    BLUE = '\033[34m'


def getInput():
    filename = sys.argv[1]
    with open(filename) as f:
        lines = f.readlines()
    stripped = []
    for line in lines:
        stripped.append(line.strip('\n'))
    return stripped


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


class SetOfCoords:
    def __init__(self, name="Anon"):
        self.coords = {}
        self.name = name

    def add(self, coord):
        if self.contains(coord):
            return False
        self.coords[coord] = coord
        return True

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
        return ",".join(str(n) for n in self.row)

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

    def inGrid(self, p, dx, dy):
        return p.row + dy >= 0 and p.row + dy < len(self.grid) and p.col + dx >= 0 and p.col + dx < self.grid[0].len()
