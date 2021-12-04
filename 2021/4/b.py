import aoc

class Board:
    def __init__(self, lines, start):
        boardSlice = slice(start, start+5)
        board = lines[boardSlice]
        self.rows = []
        for row in board:
            self.rows.append(row.split())
        self.matchedRow = [0,0,0,0,0]
        self.matchedCol = [0,0,0,0,0]
    
    def __repr__(self):
        return "%s" % self.rows

    def wins(self, call):
        nrow = -1
        for row in self.rows:
            nrow += 1
            ncol = -1
            for col in row:
                ncol += 1
                #print("testing ", nrow, ",", ncol, " - ", self.rows[nrow][ncol])
                if (col.strip() == call):
                    self.matchedRow[nrow] += 1
                    self.matchedCol[ncol] += 1
                    #print("board matched ", nrow, ",", ncol, " to ", call, " #matchedRows=", self.matchedRow[nrow], ", #matchedCols=", self.matchedCol[ncol])
                    if (self.matchedRow[nrow] == 5 or self.matchedCol[ncol] == 5):
                        return True
        return False

    def unmarked(self, calls):
        sum = 0
        nrow = -1
        for row in self.rows:
            nrow += 1
            ncol = -1
            for col in row:
                ncol += 1
                if (not calls.contains(int(col))):
                    print(col, " not found in ", calls)
                    sum += int(col)
        return sum

class Numbers:
    def __init__(self, numbers):
        self.numbers = numbers
    
    def __repr__(self):
        return "%s" % self.numbers

    def contains(self, n):
        for call in self.numbers:
            if (n == int(call)):
                return True
        return False

lines = aoc.getInput()

boards = []
start = 2
n = 0
while start < len(lines):
    boards.append(Board(lines, start))
    print("board ", n, " = ", boards[n])
    start += 6
    n += 1

def lastWinningBoard(callOrder, boards):
    finished = []
    for board in boards:
        finished.append(False)
    finishedBoards = 0

    complete = False
    n = 0
    while not complete:
        call = callOrder[n]
        #print("called ", call, ", n=", n)
        nboard = -1
        for board in boards:
            nboard += 1
            if (finished[nboard]):
                continue
            if (board.wins(call.strip())):
                print("board ", nboard, " won")
                finished[nboard] = True
                finishedBoards += 1
                if (finishedBoards == len(boards)):
                    print("board ", nboard, " is the last to win after ", call, " has been called")
                    return nboard, n
        n += 1

callOrder = lines[0].split(",")
print("call order = ", callOrder)

n, winningCallIndex = lastWinningBoard(callOrder, boards)
print("board ", n, " won after ", callOrder[winningCallIndex], " was called")

callSlice = slice(0, winningCallIndex+1)
called = Numbers(callOrder[callSlice])
sum = boards[n].unmarked(called)
print("sum = ", sum, " called = ", called, ", product = ", sum * int(callOrder[winningCallIndex]))
