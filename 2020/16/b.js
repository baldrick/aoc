'use strict';

var fs = require('fs')
var os = require('os')

function readInput(filename, cb) {
    var input = [];
    fs.readFile(filename, 'ascii', function(err, data) {
        if (err) {
            return console.log(err);
        }
        var lines = data.split(os.EOL);
        var l = 0
        for (var line of lines) {
            input[l] = line
            l++
        }
        cb(input)
    });
}

function storeClassInfo(line) {
    var nameAndNumbers = line.split(':')
    var name = nameAndNumbers[0]
    var numberRanges = nameAndNumbers[1].split(' or ')
    setValid(name, numberRanges[0], numberRanges[1])
}

function getMinMax(nr) {
    var minMax = nr.split('-')
    var min = parseInt(minMax[0])
    var max = parseInt(minMax[1])
    for (var i = min;  i <= max;  i++) {
        validNumbers[i] = true
    }
    return { min: min, max: max }
}

function setValid(name, nr1, nr2) {
    validNumbersWithInfo[name] = { one: getMinMax(nr1), two: getMinMax(nr2) }
}

function valid(n) {
    return n in validNumbers
}

function processTicket(ticket) {
    var numbers = ticket.split(',')
    var validTicket = true
    for (var sn of numbers) {
        var n = parseInt(sn)
        if (!valid(n)) {
            validTicket = false
        }
    }
    if (validTicket) {
        tickets[tickets.length] = parseTicket(ticket)
    }
}

function parseTicket(ticket) {
    var numbers = ticket.split(',')
    var parsedTicket = []
    for (var i = 0;  i < numbers.length;  i++) {
        parsedTicket[i] = parseInt(numbers[i])
    }
    return parsedTicket
}

function processLine(line, state) {
    switch (state) {
        case 0:
            // classes of info
            storeClassInfo(line)
            break
        case 1:
            // my ticket
            if (line != 'your ticket:') {
                yourTicket = parseTicket(line)
            }
            break
        case 2:
            // nearby tickets
            if (line != 'nearby tickets:') {
                processTicket(line)
            }
            break
        default:
            console.log('Unknown state ' + state)
            break
    }
}

var validNumbers = []
var validNumbersWithInfo = []
var tickets = []
var yourTicket

function dumpValidTickets() {
    console.log('There are ' + tickets.length + ' valid tickets')
    for (var t of tickets) {
        console.log(t)
    }
}

function dumpValidNumbers() {
    for (var name in validNumbersWithInfo) {
        var v = validNumbersWithInfo[name]
        console.log(name + ': ' + v.one.min + '-' + v.one.max + ' or ' + v.two.min + '-' + v.two.max)
    }
}

function inRange(tn, vr) {
    //console.log('tn ' + tn + '; vr ' + vr.min + '-' + vr.max)
    return tn >= vr.min && tn <= vr.max
}

function fieldValid(ticketNumber, validRanges) {
    var inRangeOne = inRange(ticketNumber, validRanges.one)
    var inRangeTwo = inRange(ticketNumber, validRanges.two)
    return inRangeOne || inRangeTwo
}

function calculateFieldOrder() {
    var combinations = []
    var columns = tickets[0].length
    for (var name in validNumbersWithInfo) {
        combinations[name] = []
        for (var column = 0;  column < columns;  column++) {
            var allTicketsValidForThisColumn = true
            for (var t of tickets) {
                if (!fieldValid(t[column], validNumbersWithInfo[name])) {
                    console.log('column ' + column + ' not valid for ' + name)
                    allTicketsValidForThisColumn = false
                    break
                }
            }
            combinations[name][column] = allTicketsValidForThisColumn
        }
    }

    for (var c in combinations) {
        var oneValid = findNameWithOneValidColumn(combinations)
        columnMap[oneValid.name] = oneValid.column
        columnToNameMap[oneValid.column] = oneValid.name
        console.log(oneValid.name + ' is only valid for one column (' + oneValid.column + ')')    
    }

    console.log(yourTicket)
    var m = 1
    for (var c in columnMap) {
        //console.log(c + ' maps to column ' + columnMap[c])
        if (c.startsWith('departure')) {
            console.log('*' + yourTicket[columnMap[c]])
            m = m * parseInt(yourTicket[columnMap[c]])
        }
    }
    console.log(m)
}

var columnMap = []
var columnToNameMap = []

function nameTaken(name) {
    return name in columnMap
}

function columnTaken(column) {
    return column in columnToNameMap
}

function findNameWithOneValidColumn(combinations) {
    for (var name in combinations) {
        if (nameTaken(name, combinations)) continue
        var validCount = 0
        var validColumn = -1
        for (var column in combinations[name]) {
            if (columnTaken(column)) continue
            if (combinations[name][column]) {
                validCount++
                validColumn = column
                if (validCount > 1) break
            }
        }
        if (validCount == 1) return { name: name, column: validColumn }
    }
    console.log('Oh dear, no columns are valid for only one column')
    return { name: '', column: -1 }
}
/*
class
a: 1-3 or 5-9
b: 4-5 or 6-10

create data structure

class   a        b
column
1       valid    valid
2                valid
3       valid    valid

a - 1 or 3
b - 1, 2 or 3


*/
function doit(filename) {
    readInput(filename, function(input) {
        var state = 0
        for (var line of input) {
            if (line.length == 0) {
                state++
            } else {
                processLine(line, state)
            }
        }
        dumpValidTickets()
        dumpValidNumbers()
        calculateFieldOrder()
    })
}

function main() {
    doit('input.txt')
}

function test(testfile) {
    doit(testfile)
}

if (process.argv.length > 2 && process.argv[2] == '-t') {
    if (process.argv.length > 3) {
        test(process.argv[3])
    } else {
        test('testb.txt')
    }
} else {
    main()
}
