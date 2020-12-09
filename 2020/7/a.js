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
function addContainedBy(containedBy, container, containedColour, count) {
    if (containedBy[containedColour] == undefined) {
        containedBy[containedColour] = []
    }
    len = containedBy[containedColour].length
    containedBy[containedColour][len] = { 'count': count, 'colour': container }
}

// <list of colours> no other bags | N <colour> bag<,list of colours>.
// <colour> bags contain <list of colours>
function addContainedByColours(containedBy, containingColour, restOfLine) {
    //console.log('addContainedByColours for ' + restOfLine)
    containedColours = restOfLine.split(',')
    for (var containedColour of containedColours) {
        addContainedByColour(containedBy, containingColour, containedColour)
    }
}

function addContainedByColour(containedBy, containingColour, containedColourAndCount) {
    countAndBag = containedColourAndCount.split(' bag', 1)[0]
    sCount = countAndBag.trim().split(' ', 1)[0]
    count = parseInt(sCount)
    containedColour = countAndBag.split(sCount)[1].trim()
    //console.log('parsed ' + sCount + ' to ' + count + ' for ' + containedColour + ' from ' + countAndBag)
    addContainedBy(containedBy, containingColour, containedColour, count)
}

/*
map colour -> list of containers: N + colours
*/
function decode(line, containedBy) {
    if (line.trim().length == 0) {
        return
    }
    brokenLine = line.split('bags contain')
    containingColour = brokenLine[0].trim()
    //console.log('containingColour = ' + containingColour + ', rest of line = ' + brokenLine[1])
    switch (brokenLine[1].trim()) {
        case 'no other bags.':
            addContainedBy(containedBy, containingColour, 'no other bags', 0)
            break;
        default:
            addContainedByColours(containedBy, containingColour, brokenLine[1])
            break;
    }
}

function bagsCanBeContainedBy(containedBy, colour, possibleContainers) {
    containers = containedBy[colour]
    if (containers == undefined || containers.length == 0) {
        return
    }
    //console.log(containers)
    for (var countAndColour of containers) {
        //console.log(countAndColour)
        if (containedBy[countAndColour.colour] == undefined) {
            console.log(countAndColour.colour + ' cannot be contained by any other colour')
        }
        possibleContainers[countAndColour.colour] = countAndColour.count
        bagsCanBeContainedBy(containedBy, countAndColour.colour, possibleContainers)
    }
}

function doit(filename) {
    var containedBy = []
    readInput(filename, function(input) {
        for (var line of input) {
            decode(line, containedBy)
        }
        //console.log('=====================')
        //console.log(containedBy)
        //console.log('=====================')
        possibleContainers = {}
        bagsCanBeContainedBy(containedBy, 'shiny gold', possibleContainers)
        //console.log('=====================')
        console.log(possibleContainers)
        count = 0
        Object.keys(possibleContainers).forEach(function(key, index) {
            count++
          }, possibleContainers);
        console.log(count + ' possible colours can contain shiny gold')
    })
}

function main() {
    doit('input.txt')
}

function test() {
    doit('test.txt')
}

if (process.argv.length > 2 && process.argv[2] == '-t') {
    test()
} else {
    main()
}
