import grid

def test_SimpleLinesToRows():
    input = ["123\n", "456\n", "789"]
    s = grid.SimpleLinesToRows(input)
    assert(len(s) == 3), "Should have 3 lines"

def test_SimpleRowToLetterCells():
    input = "abc"
    r = grid.SimpleRowToLetterCells(input)
    assert(len(r) == 3), "Should have 3 cells"
    i = 0
    for cell in r:
        assert(cell == input[i]), f"Cell {i} should be {input[i]} not {cell}"
        i += 1

def test_IntegerCells():
    input = "012345"
    r = grid.IntegerCells(input)
    assert(len(r) == 6), "Should have 6 cells"
    i = 0
    for n in r:
        assert(n == i), f"Cell {i} should be {i}, not {n}"
        i += 1

def test_CSV():
    input = "this,is,a,simple,test"
    csv = input.split(',')
    r = grid.CSV(input)
    assert(len(r) == 5), "Should have 5 cells"
    i = 0
    for cell in r:
        assert(cell == csv[i].strip('\n')), f"Cell {i} should be {csv[i]} not {cell}"
        i += 1

def test_SimpleGrid():
    input = ["1234\n", "5678\n", "9000"]
    g = grid.Grid(input, grid.simpleRowToLetterCellsProvider)
    assert(g.colRowSize() == (4,3)), f"Grid should report having size 4 cols x 3 rows, not {g.colRowSize()}"
    assert(g.rowColSize() == (3,4)), f"Grid should report having size 3 rows x 4 cols, not {g.rowColSize()}"
    assert(g.xySize() == (4,3)), f"Grid should report having size 4x3, not {g.xySize()}"
    assert(g.inGrid(grid.xy(3,2))), f"xy 3,2 should be in the grid"
    assert(not g.inGrid(grid.xy(4,2))), f"xy 4,2 should not be in the grid"
    assert(not g.inGrid(grid.xy(3,3))), f"xy 3,3 should not be in the grid"

def test_IntegerGrid():
    input = ["123\n", "456\n", "789\n", "999"]
    g = grid.Grid(input, grid.integerCellsProvider)
    assert(g.xySize() == (3,4)), f"Grid should report size 3x3, not {g.xySize()}"
    sum = 0
    for row in g:
        for cell in row:
            sum += cell
    assert(sum == 1+2+3+4+5+6+7+8+9+9+9+9), f"Sum of all cells in grid should be {1+2+3+4+5+6+7+8+9+9+9+9}"

def test_csvGrid():
    input = ["this,is,a,small\n", "grid,for,testing,purposes"]
    g = grid.Grid(input, grid.csvCellsProvider)
    assert(g.xySize() == (4,2)), f"Grid should report size 4x2, not {g.xySize()}"
    s = ""
    for row in g:
        for cell in row:
            s += f"{cell} "
    assert(s == "this is a small grid for testing purposes "), f"Concatenating cells gave unexpected results: {s}"

if __name__ == "__main__":
    test_SimpleLinesToRows()
    test_SimpleRowToLetterCells()
    test_IntegerCells()
    test_CSV()
    test_inGrid()