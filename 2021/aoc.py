import sys


class colors:  # You may need to change color settings
    RED = '\033[31m'
    ENDC = '\033[m'
    GREEN = '\033[32m'
    YELLOW = '\033[33m'
    BLUE = '\033[34m'


def getInput():
    filename = sys.argv[1]
    with open(filename) as f:
        lines = f.readlines()
    return lines
