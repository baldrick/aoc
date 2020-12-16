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
    for (var nr of numberRanges) {
        var minMax = nr.split('-')
        var min = parseInt(minMax[0])
        var max = parseInt(minMax[1])
        for (var i = min;  i <= max;  i++) {
            validNumbers[i] = true
        }
    }
}

function valid(n) {
    return n in validNumbers
}

function processTicket(ticket) {
    var numbers = ticket.split(',')
    for (var sn of numbers) {
        var n = parseInt(sn)
        if (!valid(n)) {
            console.log(n + ' is invalid on ticket ' + ticket)
            errorRate += n
        }
    }
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
                // ignore my ticket for now
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

var errorRate = 0
var validNumbers = []

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
        console.log('Error rate: ' + errorRate)
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
        test('test.txt')
    }
} else {
    main()
}
