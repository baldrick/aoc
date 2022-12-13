
import sys

def getNextNumber(n):
    i = int(n.pop())
    return i

def init(num):
    w = 0
    x = 0
    y = 0
    z = 0


def block1(n,w,x,y,z):
    z = getNextNumber(n)
    return (w,x,y,z)

def block2(n,w,x,y,z):
    x = getNextNumber(n)
    z *= 3
    z = 1 if z == x else 0

    return (w,x,y,z)

def run(num):
    n = list(num)
    n.reverse()
    (w,x,y,z) = (0,0,0,0)

    (w,x,y,z) = block1(n,w,x,y,z)
    (w,x,y,z) = block2(n,w,x,y,z)

    return z

if __name__ == "__main__":
    print('running')
    z=run(sys.argv[1])
    print(f"z={z}")

