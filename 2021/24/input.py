
import sys

def getNextNumber(n):
    i = int(n.pop())
    return i


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
    x = x % 26 # x in range 0..25
    z = int(z/26)
    x += -6 # x=z%26 -6 => x in range -6..19
    x = 1 if x == w else 0 #
    x = 1 if x == 0 else 0 # x = !x
    y *= 0
    y += 25
    y *= x
    y += 1 # y=1 or 26
    z *= y # z=zy
    y *= 0
    y += w
    y += 8 # y = n+8
    y *= x # y = nx+8x => y = 0 or nx+8x
    z += y
    return (w,x,y,z)

def block14(n,w,x,y,z):
    w = getNextNumber(n)
    x *= 0
    x += z
    x = x % 26  # x=z%26
    z = int(z/26) # z=z/26, from below z%26 must be 12..20 so z starts where (z/26)%26 in range 12..20
    x += -11 # x=z%26 -11 = w => z%26 = w+11 => z%26 in range 12..20
    x = 1 if x == w else 0 # x=1 if z%26==n+11 => z=1, z%26=1 != n+11 (min 12) => z%26 >= 12, from below x must = w at this point
    x = 1 if x == 0 else 0 # x = !x, from below x must be non-zero at this point
    y *= 0
    y += 25 # y=25
    y *= x # y=0 or 25
    y += 1 # y=1 or 26
    z *= y # z *= 1 or 26
    y *= 0
    y += w
    y += 5 # y = n+5
    y *= x # y = nx+5x => x is -ve or 0, must be 0 because above it's set to 1 or 0
    z += y # z += nx+5x

    return (w,x,y,z)

def run(num):
    for i in range(1,100): #100_000):
        n = list(num)
        n.reverse()
        (w,x,y,z) = (0,0,0,i)

        # for each block w=number read in, x=0, y=0, z=carried over
        # (w,x,y,z) = block1(n,w,x,y,z)
        # (w,x,y,z) = block2(n,w,x,y,z)
        # (w,x,y,z) = block3(n,w,x,y,z)
        # (w,x,y,z) = block4(n,w,x,y,z)
        # (w,x,y,z) = block5(n,w,x,y,z)
        # (w,x,y,z) = block6(n,w,x,y,z)
        # (w,x,y,z) = block7(n,w,x,y,z)
        # (w,x,y,z) = block8(n,w,x,y,z)
        # (w,x,y,z) = block9(n,w,x,y,z)
        # (w,x,y,z) = block10(n,w,x,y,z)
        # (w,x,y,z) = block11(n,w,x,y,z)

        # for z to be 1..6 after this...
        # input number = 1 => z must start at 34 (= 26 + 2n),60 = 26*2 + n, 86 = 26*3 + n
        #                2
        (w,x,y,z) = block12(n,w,x,y,z)
        if z < 1 or z > 6:
            #print(f"z={z} for i={i}")
            continue
        n12 = z+3
        print(f"n={n12}, z start={i}")

        # for z to be 12..20 after this it must start between 1..6
        # input number = 4 => z must be 1
        # in fact z = n-3, input number 4..9
        # (w,x,y,z) = block13(n,w,x,y,z)
        # if not (z >= 12 and z <= 20):
        #     #print(f"z={z} which is outside 12..20")
        #     continue
        # n13 = z-11
        # print(f"n={n13}")
        #n.append(str(z-11))

        # for z to be zero after this it must be 12..20 going in
        # in fact z = n+11
        #(w,x,y,z) = block14(n,w,x,y,z)

        #if z == 0 or (z >= 12 and z <= 20):
            #print(f"i={n13}-{i} ... w={w}, x={x}, y={y}, z={z}")
    return z

if __name__ == "__main__":
    for n in range(1, 10): #11111111111111, -1):
        s = str(n)
        print(f"running for {s}")
        if s.find("0") != -1: continue
        z = run(s)

