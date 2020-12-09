def count_exactly(id, target_count):
    character_count = {}
    for char in id:
        if char in character_count:
            character_count[char] = character_count[char] + 1
        else:
            character_count[char] = 1
    for k, v in character_count.items():
        if v == target_count:
            return 1
    return 0


def part1():
    input_file = open("input.txt")
    exactly_two = 0
    exactly_three = 0
    for box in input_file:
        exactly_two += count_exactly(box, 2)
        exactly_three += count_exactly(box, 3)

    print exactly_two * exactly_three


def get_boxes():
    f = open("input.txt")
    boxes = []
    for box in f:
        boxes.append(box.strip())
    return boxes


def count_differences(str1, str2):
    len1 = len(str1)
    len2 = len(str2)
    differences = 0
    if len1 == len2:
        for c in range(len1):
            if str1[c] != str2[c]:
                differences += 1
    else:
        differences = 999
    return differences


def differs_by_one(box, boxes):
    for check_box in boxes:
        if count_differences(box, check_box) == 1:
            return True, check_box
    return False, ""


def part2():
    boxes = get_boxes()
    for box in boxes:
        close_match, matching_box = differs_by_one(box, boxes)
        if close_match:
            print box + " and " + matching_box
            break


part2()
