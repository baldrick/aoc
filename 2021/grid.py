class xy:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return f"{self.x},{self.y}"

    def __hash__(self):
        return hash((self.x, self.y))

    def __eq__(self, other):
        return (self.x, self.y) == (other.x, other.y)

class RowProvider:
    # Iterates over all rows.
    def __iter__(self):
        pass
    
    # Reports the number of rows.
    def __len__(self):
        pass

class SimpleLinesToRows(RowProvider):
    def __init__(self, lines):
        super().__init__()
        self.lines = lines

    def __iter__(self):
        for line in self.lines:
            # Trim carriage returns and whitespace
            yield line.strip('\n').strip()

    def __len__(self):
        return len(self.lines)

def simpleLinesToRowsProvider(lines):
    return SimpleLinesToRows(lines)

class CellProvider:
    def __init__(self):
        self.row = []

    # Iterates over all cells in this row.
    def __iter__(self):
        for c in self.row:
            yield c
    
    # Reports the number of cells in this row.
    def __len__(self):
        return len(self.row)
    
class SimpleRowToLetterCells(CellProvider):
    def __init__(self, line):
        super().__init__()
        for c in line:
            self.row.append(c)
        self.line = line

class IntegerCells(CellProvider):
    def __init__(self, line):
        super().__init__()
        for c in line:
            self.row.append(int(c))

class CSV(CellProvider):
    def __init__(self, line):
        super().__init__()
        for c in line.split(','):
            self.row.append(c)

def simpleRowToLetterCellsProvider(line):
    return SimpleRowToLetterCells(line)

def integerCellsProvider(line):
    return IntegerCells(line)

def csvCellsProvider(line):
    return CSV(line)

class Grid:
    def __init__(self, input, cellProvider, rowProvider = simpleLinesToRowsProvider):
        self.grid = []
        for row in rowProvider(input):
            self.grid.append(cellProvider(row))

    def __repr__(self):
        s = ""
        for row in self.grid:
            for cell in row:
                s += f"{cell}"
            s += "\n"
        return s

    def __iter__(self):
        for row in self.grid:
            yield row

    def colRowSize(self):
        return self.xySize()

    def rowColSize(self):
        return (len(self.grid), len(self.grid[0]))

    def xySize(self):
        return (len(self.grid[0]), len(self.grid))

    def inGrid(self, xyp):
        return xyp.x >= 0 and xyp.y >= 0 and xyp.y < len(self.grid) and xyp.x < len(self.grid[0])
