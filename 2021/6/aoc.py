import sys

def getInput():
    filename = sys.argv[1]
    with open(filename) as f:
        lines = f.readlines()
    return lines
