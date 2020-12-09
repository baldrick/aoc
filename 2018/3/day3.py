def new_column(size):
    columns = []
    for column in range(size):
        columns.append(0)
    return columns


def add_row(rows, size):
    rows.append(new_column(size))


def create_fabric(size):
    rows = []
    for row in range(size):
        add_row(rows, size)
    return rows


def get_coordinates(fields):
    raw_coordinates = fields[2]
    raw_coordinates = raw_coordinates[0:len(raw_coordinates)-1]
    coordinates = raw_coordinates.split(",")
    return int(coordinates[0]), int(coordinates[1])


def get_size(fields):
    raw_size = fields[3]
    size = raw_size.split("x")
    return int(size[0]), int(size[1])


def mark_as_used(fabric, x, y, width, height):
    for w in range(width):
        for h in range(height):
            fabric[x+w][y+h] += 1


def mark_used_fabric(fabric):
    # Input is in this form: #4 @ 441,971: 21x15
    f = open("input.txt")
    for line in f:
        fields = line.split(" ")
        x, y = get_coordinates(fields)
        width, height = get_size(fields)
        mark_as_used(fabric, x, y, width, height)


def count_fabric_used_more_than_once(fabric):
    count = 0
    for x in range(len(fabric)):
        for y in range(len(fabric[x])):
            if fabric[x][y] > 1:
                count += 1
    return count


def claim_has_no_overlap(fabric, x, y, width, height):
    for w in range(width):
        for h in range(height):
            if fabric[x+w][y+h] != 1:
                return False
    return True


def claim_id_with_no_overlap(fabric):
    # Input is in this form: #4 @ 441,971: 21x15
    f = open("input.txt")
    for line in f:
        fields = line.split(" ")
        x, y = get_coordinates(fields)
        width, height = get_size(fields)
        if claim_has_no_overlap(fabric, x, y, width, height):
            print(fields[0] + " has no overlap")


def part1and2():
    size = 1000
    fabric = create_fabric(size)
    mark_used_fabric(fabric)
    print count_fabric_used_more_than_once(fabric)
    print claim_id_with_no_overlap(fabric)


part1and2()
