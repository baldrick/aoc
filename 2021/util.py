import argparse

def getInput():
    args = parser().parse_args()
    if args.filename:
        return readFile(args.filename)
    if args.input:
        return args.input

def parser():
    parser = argparse.ArgumentParser(description='Process command line arguments.')
    parser.add_argument('--file', dest='filename', help='filename from which to load input')
    parser.add_argument('--input', dest='input', help='specify input directly from the command line')
    return parser

def readFile(filename):
    with open(filename) as f:
        lines = f.readlines()
    stripped = []
    for line in lines:
        stripped.append(line.strip('\n'))
    return stripped

def formatArray(a):
    s = ""
    for item in a:
        s += f"{item}\n"
    return s