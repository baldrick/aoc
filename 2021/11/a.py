import aoc
import sys


class OctopusGrid(aoc.Grid):
    def flash(self, row, col):
        # Increment energy of surrounding octopuses
        dr = -1
        while dr <= 1:
            dc = -1
            while dc <= 1:
                if self.inGrid(aoc.Coord(row, col), dc, dr):
                    #print(f"Inc energy at {Coord(row,col)}")
                    self.grid[row+dr].row[col+dc] += 1
                dc += 1
            dr += 1

    def flashIfReady(self, flashes):
        row = 0
        while row < len(self.grid):
            col = 0
            while col < self.grid[row].len():
                if flashes.contains(aoc.Coord(row, col)):
                    #print(f"{Coord(row,col)} already flashed")
                    col += 1
                    continue
                if self.grid[row].row[col] > 9:
                    #print(f"Flashing {Coord(row,col)}")
                    self.flash(row, col)
                    flashes.add(aoc.Coord(row, col))
                col += 1
            row += 1

    def step(self):
        # increase energy for every octopus
        row = 0
        while row < len(self.grid):
            col = 0
            while col < self.grid[row].len():
                self.grid[row].row[col] += 1
                col += 1
            row += 1

        # flash every octopus at energy 9
        flashes = aoc.SetOfCoords()
        flashesDone = -1  # doesn't matter where we start as long as loop executes at least once
        while flashesDone != flashes.size():
            flashesDone = flashes.size()
            self.flashIfReady(flashes)
            #print(f"flashed {flashes.size()} times, increase of {flashes.size() - flashesDone}")
            # print(self)

        # Reset energy >9 to 0
        row = 0
        while row < len(self.grid):
            col = 0
            while col < self.grid[row].len():
                if self.grid[row].row[col] > 9:
                    self.grid[row].row[col] = 0
                col += 1
            row += 1

        return flashes.size()


lines = aoc.getInput()

grid = OctopusGrid(lines)

flashes = 0
steps = int(sys.argv[2])
for n in range(0, steps):
    print(f"step {n}")
    print(grid)
    extraFlashes = grid.step()
    flashes += extraFlashes
    if extraFlashes == 100:
        print(f"synchronization achieved at step {n+1}")
        exit()

print(f"completed {steps} steps, there have been {flashes} flashes")
print(grid)
