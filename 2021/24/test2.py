
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
    w = getNextNumber(n)
    z += w
    z = z % 2
    w = int(w/2)
    y += w
    y = y % 2
    w = int(w/2)
    x += w
    x = x % 2
    w = int(w/2)
    w = w % 2

    return (w,x,y,z)

def run(num):
    n = list(num)
    n.reverse()
    (w,x,y,z) = (0,0,0,0)

    (w,x,y,z) = block1(n,w,x,y,z)

    print(f"w={w}, x={x}, y={y}, z={z}")
    return z

if __name__ == "__main__":
    print('running')
    run(sys.argv[1])

