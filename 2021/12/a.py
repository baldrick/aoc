import aoc

lines = aoc.getInput()


class Node:
    def __init__(self, connections):
        self.connections = [connections]

    def __repr__(self):
        return "%s" % self.connections

    def connect(self, connection):
        self.connections.append(connection)


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


def pathsFrom(connections, path, paths):
    #print(f"connections: {connections}, length: {len(connections)}")
    if len(connections) == 0:
        return
    for connection in connections:
        #print(f"{path}: {connection}")
        if connection == "end":
            completePath = path.copy()
            completePath.append("end")
            paths.append(completePath)
            continue
        if connection in path:
            if not connection == connection.upper():
                # We've visited this cave before but it's not revisitable.
                continue
        if connection in nodes:
            newPath = path.copy()
            newPath.append(connection)
            pathsFrom(nodes[connection].connections, newPath, paths)


print(nodes)

paths = []
pathsFrom(nodes["start"].connections, ["start"], paths)
print(paths)
print(len(paths))
