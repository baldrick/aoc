import aoc

lines = aoc.getInput()

def addCodes(codeToDigit, digitToCode, input):
    for digit in input.split():
        match len(digit):
            case 2:
                codeToDigit[digit] = 1
                digitToCode[1] = digit
            case 3:
                codeToDigit[digit] = 7
                digitToCode[7] = digit
            case 4:
                codeToDigit[digit] = 4
                digitToCode[4] = digit
            case 7:
                codeToDigit[digit] = 8
                digitToCode[8] = digit
            case _:
                codeToDigit[digit] = -1

def getUniqueCodes(line):
    codeToDigit = {}
    digitToCode = {}
    addCodes(codeToDigit, digitToCode, line.split("|")[0])
    return codeToDigit, digitToCode

def findOverlap(codes):
    overlap = []
    for letter in codes[0]:
        count = 0
        for code in codes:
            if letter in code:
                count += 1
        if count == len(codes):
            overlap.append(letter)
    return overlap

def decode0369(codeToDigit, digitToCode):
    for key in codeToDigit:
        match len(key):
            case 5:
                # 3 shares two code lines with 1 whereas 2 and 5 only share 1 each
                overlap = findOverlap([digitToCode[1], key])
                if len(overlap) == 2:
                    print("Setting codeToDigit[", key, "] = 3")
                    digitToCode[3] = key
                    codeToDigit[key] = 3
            case 6:
                # 0 and 9 share three code lines with 7 whereas 6 only shares 2
                overlap = findOverlap([digitToCode[7], key])
                if len(overlap) == 2:
                    digitToCode[6] = key
                    codeToDigit[key] = 6
                if len(overlap) == 3:
                    # key could be 0 or 9
                    overlap = findOverlap([digitToCode[4],key])
                    if len(overlap) == 4:
                        digitToCode[9] = key
                        codeToDigit[key] = 9
                    if len(overlap) == 3:
                        digitToCode[0] = key
                        codeToDigit[key] = 0

def decode25(codeToDigit, digitToCode):
    for key in codeToDigit:
        match len(key):
            case 5:
                overlap = findOverlap([digitToCode[6], key])
                if len(overlap) == 5:
                    digitToCode[5] = key
                    codeToDigit[key] = 5
                if len(overlap) == 4:
                    # Could be 2 or 3...  3 has been set above so we know when it's a 2
                    if codeToDigit[key] == -1:
                        digitToCode[2] = key
                        codeToDigit[key] = 2

def decode(codeToDigit, outputCode):
    for code in codeToDigit:
        if len(findOverlap([code, outputCode])) == max(len(outputCode), len(code)):
            print("decoded", outputCode, "to", codeToDigit[code])
            return codeToDigit[code]

def decodeOutput(codeToDigit, codes):
    output = []
    for code in codes.split():
        outputDigit = decode(codeToDigit, code)
        if outputDigit != None:
            output.append(outputDigit)
    return output

sum = 0
for line in lines:
    # Get unique codes
    codeToDigit, digitToCode = getUniqueCodes(line)

    # Check we have something for every number
    if len(codeToDigit) != 10:
        print("We're missing", 10 - len(codeToDigit), "digits")

    decode0369(codeToDigit, digitToCode)
    decode25(codeToDigit, digitToCode)
    print(codeToDigit)
    print(digitToCode)
    output = decodeOutput(codeToDigit, line.split("|")[1])
    print(output)
    n = ""
    for c in output:
        n += str(c)
    sum += int(n)

print(sum)