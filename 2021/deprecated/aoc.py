import deprecation

@deprecation.deprecated(deprecated_in="20211218", removed_in="2022",
                        current_version="20211218",
                        details="Use the broken out colors / grid / util modules instead")

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

    def xy(self):
        return f"{self.col},{self.row}"

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
        return "".join(str(n) for n in self.row)

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

    # Oh what was I thinking mixing row,col and x,y!
    def inGrid(self, p, dx, dy):
        return p.row + dy >= 0 and p.row + dy < len(self.grid) and p.col + dx >= 0 and p.col + dx < self.grid[0].len()

    def cell(self, p):
        return self.grid[p.row].row[p.col]

    def isLastCell(self, p):
        return p.row == len(self.grid)-1 and p.col == len(self.grid[0].row)-1


class NodeFactory:
    _id = 0

    def create_node(self, data):
        self._id += 1
        return Node(data, self._id)


class Node:
    def __init__(self, data, id):
        self.data = data
        self.id = id
        self.next = None

    def __repr__(self):
        return f"{self.id}:{self.data}"

    def __eq__(self, other):
        if self is None and other is None:
            return True
        if (self is None and other is not None) or (self is not None and other is None):
            return False
        return self.id == other.id


class LinkedList:
    def __init__(self):
        self.head = None

    def __repr__(self):
        node = self.head
        nodes = []
        while node is not None:
            nodes.append(f"{node}")
            node = node.next
        nodes.append("None")
        return " -> ".join(nodes)

    def simple(self):
        node = self.head
        nodes = []
        while node is not None:
            nodes.append(f"{node.data}")
            node = node.next
        return "".join(nodes)

    def __iter__(self):
        node = self.head
        while node is not None:
            yield node
            node = node.next

    def __len__(self):
        n = 0
        for c in self:
            n += 1
            pass
        return n

    def prepend(self, node):
        node.prev = None
        node.next = self.head
        self.head.prev = node
        self.head = node

    def append(self, node):
        if self.head is None:
            self.head = node
            return
        for current_node in self:
            pass
        current_node.next = node
        node.prev = current_node

    def append_after(self, target_node, new_node):
        if self.head is None:
            raise Exception("List is empty")

        old_next = target_node.next
        target_node.next = new_node
        new_node.next = old_next
        new_node.prev = target_node
        old_next.prev = new_node

    def prepend_before(self, target_node, new_node):
        if self.head is None:
            raise Exception("List is empty")

        old_prev = target_node.prev
        target_node.prev = new_node
        new_node.prev = old_prev
        old_prev.next = new_node
        new_node.next = target_node


def log(level, s):
    if len(sys.argv) > 2 and sys.argv[2] == 'debug':
        if int(sys.argv[3]) >= level:
            print(s, flush=True)
