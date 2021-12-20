import grid
import util

class Algo:
    def __init__(self, input):
        s = ""
        for line in input:
            s += line
            if s == "": break
        self.iea = []
        for n in range(0, 512):
            self.iea.append(s[n])
    
    def __repr__(self):
        return f"{self.iea}"

    def __getitem__(self, item):
        # convert item to number
        n = 0
        for bit in item:
            n *= 2
            if bit in ["#", "-"]:
                n += 1
        #print(f"decode {item} to {n} which maps to {self.iea[n]}")
        return self.iea[n]

    def background(self, step):
        if self.iea[0] == ".":
            return "."
        else:
            return "#" if step % 2 else "."


class Image(grid.CharacterGrid):
    def __init__(self, input):
        imageStart = 0
        for n in range(0, len(input)):
            if input[n] == "":
                imageStart = n
                break
        super().__init__(input[imageStart+1:len(input)])

    def enhance(self, algo, step):
        enhancedImage = Image("")
        for y in range(-1, len(self)+1):
            row = []
            #print(f"y:{y} len(self[0]):{len(self[0])}, self[0]:{self[0]}")
            for x in range(-1, len(self[0])+1):
                row.append(self.enhancePixel(algo, x, y, step))
            enhancedImage.grid.append(row)
        return enhancedImage

    def enhancePixel(self, algo, x, y, step):
        n = ''
        for dy in range(-1, 2):
            for dx in range(-1, 2):
                xyp = grid.xy(x+dx, y+dy)
                if self.inGrid(xyp):
                    n += self.cell(xyp)
                else:
                    n += algo.background(step)
        return algo[n]

    def countLit(self):
        count = 0
        for row in self:
            for col in row:
                if col == "#":
                    count += 1
        return count
    
def enhance(input, repetitions):
    algo = Algo(input)
    image = Image(input)
    # print(f"algo: {algo}")
    # print(f"image:\n{image}")
    for n in range(0, repetitions):
        print(f"enhancing, step {n}", flush=True)
        image = image.enhance(algo, n)
    print(f"enhanced image:\n{image}")
    print(f"{image.countLit()} lit cells")
    return image

if __name__ == "__main__":
    input = util.getInput()
    enhance(input, 2)
    enhance(input, 50)
