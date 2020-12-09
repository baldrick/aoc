def parse_coordinate(coordinate):
    split_coordinate = coordinate.split(",")
    x = int(split_coordinate[0].strip())
    y = int(split_coordinate[1].strip())
    return x, y


def find_grid_size(coordinates):
    x = 0
    y = 0
    for coordinate in coordinates:
        new_x, new_y = parse_coordinate(coordinate)
        if new_x > x:
            x = new_x
        if new_y > y:
            y = new_y
    return x+1, y+1


def append_row(grid, x):
    row = []
    for col in range(x):
        row.append("")
    grid.append(row)


def create_grid(coordinates):
    x, y = find_grid_size(coordinates)
    grid = []
    for row in range(y):
        append_row(grid, x)
    return grid


def create_coordinate_id(coordinate_id):
    ids = "*ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()-=_+[]{};:<>?,./abcdefghijklmn"
    return ids[coordinate_id]


def populate_coordinates(grid, coordinates):
    coordinate_id = 1
    for coordinate in coordinates:
        x, y = parse_coordinate(coordinate)
        grid[y][x] = create_coordinate_id(coordinate_id)
        coordinate_id += 1


def calculate_distance(x1, y1, x2, y2):
    return abs(x1 - x2) + abs(y1 - y2)


def find_closest_coordinate(x, y, grid, coordinates):
    closest_coordinate = -1
    closest_coordinate_distance = len(grid) + len(grid[0]) + 1
    tie = False
    for coordinate in coordinates:
        coordinate_x, coordinate_y = parse_coordinate(coordinate)
        if coordinate_x != x or coordinate_y != y:
            distance = calculate_distance(x, y, coordinate_x, coordinate_y)
            if distance < closest_coordinate_distance:
                closest_coordinate_distance = distance
                closest_coordinate = grid[coordinate_y][coordinate_x].lower()
                tie = False
            elif distance == closest_coordinate_distance:
                tie = True
    if tie:
        return "."
    else:
        return closest_coordinate


def populate_closest_coordinates(grid, coordinates):
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if grid[y][x] == "":
                grid[y][x] = find_closest_coordinate(x, y, grid, coordinates)


def dump_grid(message, grid, threshold):
    print(message)
    for row in grid:
        str_col = ""
        for col in row:
            if col < threshold:
                str_col += str(col) + " "
            else:
                str_col += ". "
        print(str_col)


def valid_coordinate(grid, x, y):
    return 0 <= y < len(grid) and 0 <= x < len(grid[0])


def search(grid, coordinate, dx, dy):
    x, y = parse_coordinate(coordinate)
    coordinate_id = grid[y][x].lower()
    while valid_coordinate(grid, x, y) and grid[y][x].lower() == coordinate_id:
        x += dx
        y += dy
    if valid_coordinate(grid, x, y):
        return False, x - dx, y - dy
    else:
        return True, -1, -1


# Return -1 if area is infinite
def calculate_area(grid, left, right, top, bottom, coordinate_id):
    calculated_area = 0
    for x in range(left, right + 1):
        for y in range(top, bottom + 1):
            if grid[y][x].upper() == coordinate_id:
                calculated_area += 1
    return calculated_area


def area(grid, coordinate):
    infinite, left, ignore = search(grid, coordinate, -1, 0)
    if not infinite:
        infinite, right, ignore = search(grid, coordinate, 1, 0)
    if not infinite:
        infinite, ignore, top = search(grid, coordinate, 0, -1)
    if not infinite:
        infinite, ignore, bottom = search(grid, coordinate, 0, 1)

    x, y = parse_coordinate(coordinate)
    if infinite:
        return -1

    coordinate_area = calculate_area(grid, left, right, top, bottom, grid[y][x])
    return coordinate_area


def largest_finite_area(grid, coordinates):
    largest_area = 0
    for coordinate in coordinates:
        coordinate_area = area(grid, coordinate)
        if coordinate_area > largest_area:
            largest_area = coordinate_area
    return largest_area


def part1(filename):
    coordinates = open(filename).readlines()
    grid = create_grid(coordinates)
    populate_coordinates(grid, coordinates)
    populate_closest_coordinates(grid, coordinates)
    print(largest_finite_area(grid, coordinates))


def calculate_distances_to_all_coordinates(grid, coordinates):
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            total_distance = 0
            for coordinate in coordinates:
                cx, cy = parse_coordinate(coordinate)
                total_distance += calculate_distance(x, y, cx, cy)
            grid[y][x] = total_distance


def count_of_locations_within_threshold(grid, threshold):
    count = 0
    for row in grid:
        for col in row:
            if col < threshold:
                count += 1
    return count


def part2(filename, threshold):
    coordinates = open(filename).readlines()
    grid = create_grid(coordinates)
    populate_coordinates(grid, coordinates)
    calculate_distances_to_all_coordinates(grid, coordinates)
    print(count_of_locations_within_threshold(grid, threshold))


part2("input.txt", 10000)
