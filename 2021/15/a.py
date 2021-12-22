from collections import defaultdict
import grid
import heapq as heap
import util

class ChitonGrid(grid.IntegerGrid):
    def scale(self):
        self.scaleWidth()
        self.scaleHeight()

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
                newRow = self.cellProvider("")
                for col in row.row:
                    x = col + m
                    if x > 9:
                        x -= 9
                    newRow.row.append(x)
                newGrid.append(newRow)
        self.grid = newGrid

# Inspired by (ok, basically copied from)
# https://levelup.gitconnected.com/dijkstra-algorithm-in-python-8f0e75e3f16e
def dijkstraCopy(G, startingNode):
    visited = set()
    parentsMap = {}
    pq = []
    nodeCosts = defaultdict(lambda: float('inf'))
    nodeCosts[startingNode] = 0
    heap.heappush(pq, (0, startingNode))

    while pq:
        # go greedily by always extending the shorter cost nodes first
        _, node = heap.heappop(pq)
        visited.add(node)

        for adjNode in G.neighbours(node):
            if adjNode in visited:	continue

            weight = G.cell(adjNode)
            newCost = nodeCosts[node] + weight
            if nodeCosts[adjNode] > newCost:
                parentsMap[adjNode] = node
                nodeCosts[adjNode] = newCost
                heap.heappush(pq, (newCost, adjNode))

    return parentsMap, nodeCosts

g = ChitonGrid(util.getInput())
g.scale()

(pmap, dists) = dijkstraCopy(g, grid.xy(0,0))
(x,y) = g.xySize()
print(f"dist={dists[grid.xy(x-1, y-1)]}")
