import a
import util

testInput = [
"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##",
"#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###",
".######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#.",
".#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#.....",
".#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#..",
"...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.....",
"..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
"",
"#..#.",
"#....",
"##..#",
"..#..",
"..###",
]

def test_decode():
    algo = a.Algo(testInput)
    assert(algo["...#...#."] == "#"), "Decode failed"

def test_enhancePixel():
    algo = a.Algo(testInput)
    image = a.Image(testInput)
    assert(image.enhancePixel(algo, 2, 2) == "#"), "Enhance pixel failed"

def test_enhance():
    assert(a.enhance(testInput, 2).countLit() == 35), "Image enhancement failed"

def test_enhanceLarge():
    assert(a.enhance(util.readFile("test2"), 2).countLit() == 5326), "Larger image enhancement failed"
