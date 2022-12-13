import util

'''
inp a - Read an input value and write it to variable a.
add a b - Add the value of a to the value of b, then store the result in variable a.
mul a b - Multiply the value of a by the value of b, then store the result in variable a.
div a b - Divide the value of a by the value of b, truncate the result to an integer, then store the result in variable a. (Here, "truncate" means to round the value toward zero.)
mod a b - Divide the value of a by the value of b, then store the remainder in variable a. (This is also called the modulo operation.)
eql a b - If the value of a and b are equal, then store the value 1 in variable a. Otherwise, store the value 0 in variable a.
'''

def apply(cli, blockNum):
    s = cli.split()
    cmd = s[0]
    a = s[1]
    match cmd:
        case "inp":
            if blockNum > 1:
                print("    return (w,x,y,z)")
            print("")
            print(f"def block{blockNum}(n,w,x,y,z):")
            blockNum += 1
            print(f"    {a} = getNextNumber(n)")
        case 'add':
            b = s[2]
            print(f"    {a} += {b}")
        case 'mul':
            b = s[2]
            print(f"    {a} *= {b}")
        case 'div':
            b = s[2]
            print(f"    {a} = int({a}/{b})")
        case 'mod':
            b = s[2]
            print(f"    {a} = {a} % {b}")
        case 'eql':
            b = s[2]
            print(f"    {a} = 1 if {a} == {b} else 0")
        case _:
            print(f"    # illegal instruction {cli}")
    #print('    print(f"w={w}, x={x}, y={y}, z={z}")')
    return blockNum

input = util.getInput()
print("""
import sys

def getNextNumber(n):
    i = int(n.pop())
    return i
""")

blockNum = 1

for cmd in input:
    blockNum = apply(cmd, blockNum)

print("""
    return (w,x,y,z)

def run(num):
    n = list(num)
    n.reverse()
    (w,x,y,z) = (0,0,0,0)
""")

for n in range(1, blockNum):
    print(f"    (w,x,y,z) = block{n}(n,w,x,y,z)")

print("""
    print(f"w={w}, x={x}, y={y}, z={z}")
    return z

if __name__ == "__main__":
    print('running')
    run(sys.argv[1])
""")

'''
block 1 - w=n, x=1, y=n+7, z=n+7
'''
