fs = require('fs')
os = require('os')
assert = require('assert')

function readInput(filename, cb) {
    var input = []
    fs.readFile(filename, 'ascii', function(err, data) {
        if (err) {
            return console.log(err);
        }
        lines = data.split(os.EOL);
        l = 0
        for (var line of lines) {
            input[l] = line
            l++
        }
        cb(input)
    });
}
function addContains(contains, container, containedColour, count) {
    if (contains[container] == undefined) {
        contains[container] = []
    }
    len = contains[container].length
    contains[container][len] = { 'count': count, 'colour': containedColour }
}

// <list of colours> no other bags | N <colour> bag<,list of colours>.
// <colour> bags contain <list of colours>
function addContainsColours(contains, containingColour, restOfLine) {
    //console.log('addContainedByColours for ' + restOfLine)
    containedColours = restOfLine.split(',')
    for (var containedColour of containedColours) {
        addContainsColour(contains, containingColour, containedColour)
    }
}

function addContainsColour(contains, containingColour, containedColourAndCount) {
    countAndBag = containedColourAndCount.split(' bag', 1)[0]
    sCount = countAndBag.trim().split(' ', 1)[0]
    count = parseInt(sCount)
    containedColour = countAndBag.split(sCount)[1].trim()
    //console.log('parsed ' + sCount + ' to ' + count + ' for ' + containedColour + ' from ' + countAndBag)
    addContains(contains, containingColour, containedColour, count)
}

/*
map colour -> list of containers: N + colours
*/
function decode(line, contains) {
    if (line.trim().length == 0) {
        return
    }
    brokenLine = line.split('bags contain')
    containingColour = brokenLine[0].trim()
    switch (brokenLine[1].trim()) {
        case 'no other bags.':
            addContains(contains, containingColour, 'no other bags', 0)
            break;
        default:
            addContainsColours(contains, containingColour, brokenLine[1])
            break;
    }
}

function bagsContain(contains, colour) {
    count = 0
    contained = contains[colour]
    if (contained == undefined || contained.length == 0) {
        console.log('contained is undefined, returning ' + count)
        return count
    }
    //console.log(contained)
    for (var countAndColour of contained) {
        if (contains[countAndColour.colour] == undefined) {
            console.log(countAndColour.colour + ' contains nothing!')
        } else {
            count += countAndColour.count
            count += countAndColour.count * bagsContain(contains, countAndColour.colour)
            console.log(countAndColour.colour + ' contains ' + countAndColour.count + ' for a total of ' + count)
        }
    }
    return count
}

function doit(filename) {
    var contains = []
    readInput(filename, function(input) {
        for (var line of input) {
            decode(line, contains)
        }
        cumulativeContents = bagsContain(contains, 'shiny gold')
        console.log('a shiny gold bag contains ' + cumulativeContents + ' bags')

    })
}

function main() {
    doit('input.txt')
}

function test() {
    doit('testb.txt')
}

if (process.argv.length > 2 && process.argv[2] == '-t') {
    test()
} else {
    main()
}
