
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
    x *= 0
    x += z
    x = x % 26
    z = int(z/1)
    x += 12
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 7
    y *= x
    z += y
    return (w,x,y,z)

def block2(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/1)
    x += 12
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 8
    y *= x
    z += y
    return (w,x,y,z)

def block3(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/1)
    x += 13
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 2
    y *= x
    z += y
    return (w,x,y,z)

def block4(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/1)
    x += 12
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 11
    y *= x
    z += y
    return (w,x,y,z)

def block5(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/26)
    x += -3
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 6
    y *= x
    z += y
    return (w,x,y,z)

def block6(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/1)
    x += 10
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 12
    y *= x
    z += y
    return (w,x,y,z)

def block7(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/1)
    x += 14
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 14
    y *= x
    z += y
    return (w,x,y,z)

def block8(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/26)
    x += -16
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 13
    y *= x
    z += y
    return (w,x,y,z)

def block9(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/1)
    x += 12
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 15
    y *= x
    z += y
    return (w,x,y,z)

def block10(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/26)
    x += -8
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 10
    y *= x
    z += y
    return (w,x,y,z)

def block11(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/26)
    x += -12
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 6
    y *= x
    z += y
    return (w,x,y,z)

def block12(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/26)
    x += -7
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 10
    y *= x
    z += y
    return (w,x,y,z)

def block13(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/26)
    x += -6
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 8
    y *= x
    z += y
    return (w,x,y,z)

def block14(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26
    z = int(z/26)
    x += -11
    x = 1 if x == w else 0
    x = 1 if x == 0 else 0
    y *= 0
    y += 25
    y *= x
    y += 1
    z *= y
    y *= 0
    y += w
    y += 5
    y *= x
    z += y
    return (w,x,y,z)

def run(num):
    n = list(num)
    n.reverse()
    (w,x,y,z) = (0,0,0,0)
    (w,x,y,z) = block1(n,w,x,y,z)
    print(f"block 1 {num}: {w},{x},{y},{z}")
    (w,x,y,z) = block2(n,w,x,y,z)
    print(f"block 2 {num}: {w},{x},{y},{z}")
    return z

if __name__ == "__main__":
    print('running')
    for n in range(11, 21):
        run(str(n))
