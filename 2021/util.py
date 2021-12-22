import sys

def getInput():
    return readFile(sys.argv[1])

def readFile(filename):
    with open(filename) as f:
        lines = f.readlines()
    stripped = []
    for line in lines:
        stripped.append(line.strip('\n'))
    return stripped
