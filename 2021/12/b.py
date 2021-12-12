import aoc

lines = aoc.getInput()


class Node:
    def __init__(self, connections):
        self.connections = [connections]

    def __repr__(self):
        return "%s" % self.connections

    def connect(self, connection):
        self.connections.append(connection)


class Path:
    def __init__(self, nodes):
        self.nodes = nodes

    def __repr__(self):
        return "%s" % self.nodes

    def __hash__(self):
        return hash(','.join(self.nodes))

    def __eq__(self, other):
        return (','.join(self.nodes) == ','.join(other.nodes))

    def __lt__(self, other):
        return ','.join(self.nodes) < ','.join(other.nodes)


def connect(left, right):
    if not left in nodes:
        nodes[left] = Node(right)
    else:
        nodes[left].connect(right)


nodes = {}
for line in lines:
    split = line.split("-")
    left = split[0]
    right = split[1]
    connect(left, right)
    connect(right, left)


def pathsFrom(connections, path, paths, revisitSmallCave):
    #print(f"connections: {connections}, length: {len(connections)}")
    if len(connections) == 0:
        return
    for connection in connections:
        #print(f"{path}: {connection}")
        if connection == "end":
            completePath = path.copy()
            completePath.append("end")
            paths.add(Path(completePath))
            continue
        if connection in path:
            if not connection == connection.upper():
                if revisitSmallCave != connection:
                    # We've visited this cave more than once before but it's not revisitable.
                    continue
                count = 0
                for c in path:
                    if c == connection:
                        count += 1
                if count > 1:
                    continue
        if connection in nodes:
            newPath = path.copy()
            newPath.append(connection)
            pathsFrom(nodes[connection].connections, newPath, paths, revisitSmallCave)


print(f"nodes: {nodes}")

paths = set()
smallCaves = []
for cave in nodes:
    if not cave == cave.upper() and cave != "start" and cave != "end":
        smallCaves.append(cave)
print(f"small caves: {smallCaves}")

for smallCave in smallCaves:
    print(f"getting paths where we can revisit '{smallCave}'")
    pathsFrom(nodes["start"].connections, ["start"], paths, smallCave)

for path in sorted(paths):
    print(path)
print(len(paths))
