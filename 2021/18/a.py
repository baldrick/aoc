import math
import sys
import util

class SnailfishException(Exception):
    pass

class Explosion(Exception):
    pass

class ExplosionComplete(Exception):
    pass

class Split(Exception):
    pass

class SnailfishNumber:
    def __init__(self, left, right):
        self.left = left
        self.right = right
        self.parent = None
        if isinstance(left, SnailfishNumber):
            self.left.setParent(self)
        if isinstance(right, SnailfishNumber):
            self.right.setParent(self)

    def setParent(self, parent):
        if self.parent is not None:
            raise SnailfishException(f"{self} already has parent {self.parent}")
        #print(f"parent of {self} is now {parent}")
        self.parent = parent

    def depth(self):
        n = 1
        if self.parent is not None:
            n = self.parent.depth() + 1
        #print(f"{self} has depth {n}")
        return n

    @staticmethod
    def create(line):
        cb = SnailfishNumber.findClosingBracketOrComma(line, 1)
        cbLeftClose = SnailfishNumber.findClosingBracketOrComma(line, 1)
        left = SnailfishNumber.parse(line, 1, cbLeftClose)
        cbRightClose = SnailfishNumber.findClosingBracketOrComma(line, cbLeftClose+1)
        right = SnailfishNumber.parse(line, cbLeftClose+1, cbRightClose)
        sfn = SnailfishNumber(left, right)
        if isinstance(left, SnailfishNumber):
            left.parent = sfn
        if isinstance(right, SnailfishNumber):
            right.parent = sfn
        return sfn

    @staticmethod
    def parse(input, start, end):
        #print(f"parsing {input[start:end]}")
        if input[start] == '[':
            cb = SnailfishNumber.findClosingBracketOrComma(input, start)
            return SnailfishNumber.create(input[start:cb])
        elif input[start] in ['0','1','2','3','4','5','6','7','8','9']:
            n = input[start:end].split(',')
            return int(n[0])
        else:
            print(f"Failed to parse {input} from {start}-{end}")

    def __repr__(self):
        #return f"#{self.depth()}: [{self.left},{self.right}]"
        return f"[{self.left},{self.right}]"

    @staticmethod
    def findClosingBracketOrComma(s, start):
        openBracketCount = 1
        for c in range(start, len(s)):
            if s[c] == '[':
                openBracketCount += 1
            if s[c] == ']':
                openBracketCount -= 1
                if openBracketCount == 0:
                    return c
            if s[c] == ',' and openBracketCount == 1:
                return c
        print(f"Closing bracket not found from {start} in {s}")

    def add(self, right):
        return SnailfishNumber(self, right)

    def reduce(self):
        #print(f"reducing {self}")
        try:
            self.explodeIfDeeplyNested()
            #print(f"no explosion for {self}")
            self.splitIfOver9()
        except ExplosionComplete:
            return True
        except Split:
            return True
        return False
    
    def fullReduce(self):
        while self.reduce():
            pass

    def explodeIfDeeplyNested(self):
        if isinstance(self.left, SnailfishNumber):
            try:
                self.left.explodeIfDeeplyNested()
            except Explosion:
                self.left = 0
                raise ExplosionComplete # more eww using exceptions for standard flow

        if isinstance(self.right, SnailfishNumber):
            try:
                self.right.explodeIfDeeplyNested()
            except Explosion:
                self.right = 0
                raise ExplosionComplete # more eww using exceptions for standard flow
        
        #print(f"checking {self} for explodability")
        if self.depth() > 4:
            self.explode()
    
    def explode(self):
        #print(f"exploding {self}")
        if isinstance(self.left, SnailfishNumber) or isinstance(self.right, SnailfishNumber):
            raise SnailfishException(f"don't expect to explode a pair with a SnailfishNumber: {self}")

        # if we have [[a,b],[xl,xr]] we need to add xl to a and xr to b
        # so we find ancestor of [xl,xr] that has [xl,xr] on its right
        # then go to that ancestor's immediate left then down its right to find b
        #
        #          gp
        #          / \
        #        p1   p2
        #        /\   /\
        #        a b xl xr

        # add xl to b
        # find parent that has self on its right
        child = self
        parent = self.parent
        #print(f"child={self}, parent={self.parent}, parent.left={self.parent.left}, parent.right={self.parent.right}, looking for ancestor with child on its right")
        while parent is not None and parent.left == child:
            grandParent = parent.parent
            child = parent
            parent = grandParent
            #print(f"parent.left = {parent.left}, child = {child}")
        #print(f"{parent} is ancestor with {self} on RHS")

        # now go down left side of this ancestor finding right-most regular number
        if parent is not None:
            child = parent.left
            up = 0
            while isinstance(child, SnailfishNumber):
                parent = child
                child = child.right
                up += 1

            # If we're dealing with the immediate parent, add to the right
            if up == 0:
                #print(f"going to add {self.left} to left of {parent}")
                parent.left += self.left
            else:
                #print(f"going to add {self.left} to right of {parent}")
                parent.right += self.left

        # find parent whose right isn't an ancestor of self
        child = self
        parent = self.parent
        #print(f"child={self}, parent={self.parent}, parent.left={self.parent.left}, parent.right={self.parent.right}, looking for ancestor with child on its left")
        while parent is not None and parent.right == child:
            grandParent = parent.parent
            child = parent
            parent = grandParent
            #print(f"parent.right = {parent.right}, child = {child}")
        #print(f"{parent} is ancestor with {self} on LHS")

        # now go down right side of this ancestor finding left-most regular number
        if parent is not None:
            child = parent.right
            up = 0
            while isinstance(child, SnailfishNumber):
                parent = child
                child = child.left
                up += 1

            # If we're dealing with the immediate parent, add to the right
            if up == 0:
                #print(f"going to add {self.right} to right of {parent}")
                parent.right += self.right
            else:
                #print(f"going to add {self.right} to left of {parent}")
                parent.left += self.right
        
        raise Explosion # eww using exception handling for standard program flow!

    def splitIfOver9(self):
        #print(f"checking for split {self}")
        if isinstance(self.left, SnailfishNumber):
            self.left.splitIfOver9()
        else:
            if self.left > 9:
                #print(f"splitting {self.left}")
                self.left = SnailfishNumber.split(self.left, self)
                raise Split

        if isinstance(self.right, SnailfishNumber):
            self.right.splitIfOver9()
        else:
            if self.right > 9:
                #print(f"splitting {self.right}")
                self.right = SnailfishNumber.split(self.right, self)
                raise Split

    @staticmethod
    def split(value, parent):
        left = math.floor(value/2)
        right = math.ceil(value/2)
        sfn = SnailfishNumber(left, right)
        sfn.setParent(parent)
        return sfn
    
    def magnitude(self):
        '''
        The magnitude of a pair is 3 times the magnitude of its left element plus 2 times
        the magnitude of its right element. The magnitude of a regular number is just that number.
        '''
        if isinstance(self.left, SnailfishNumber):
            ml = self.left.magnitude()
        else:
            ml = self.left
        if isinstance(self.right, SnailfishNumber):
            mr = self.right.magnitude()
        else:
            mr = self.right
        return 3 * ml + 2 * mr

def main(filename):
    input = util.readFile(filename)

    print(input)
    sfn = SnailfishNumber.create(input[0])

    for n in range(1, len(input)):
        sfn = sfn.add(SnailfishNumber.create(input[n]))
        sfn.fullReduce()

    return f"{sfn}"

def magnitude(input):
    sfn = SnailfishNumber.create(input)
    return sfn.magnitude()

def largestFromTwo(filename):
    input = util.readFile(filename)

    max = 0
    for i in range(0, len(input)):
        for j in range(0, len(input)):
            if i != j:
                sfn1 = SnailfishNumber.create(input[i])
                sfn2 = SnailfishNumber.create(input[j])
                sum = sfn1.add(sfn2)
                sum.fullReduce()
                magnitude = sum.magnitude()
                if magnitude > max:
                    max = magnitude
    return max

if __name__ == "__main__":
    sfn = main(sys.argv[1])
    print(f"{magnitude(sfn)}")

    print(f"{largestFromTwo(sys.argv[1])}")
