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

function dump(numbers, n) {
    return 'numbers[' + n + '] = { oldest: ' + numbers[n].oldest + ', newest: ' + numbers[n].newest + ' }'
}

function doit(filename) {
    readInput(filename, function(input) {
        var starting = input[0].split(',')
        var numbers = []
        var turn = 1
        var lastSpoken
        var seenCount = []
        for (var s of starting) {
            lastSpoken = parseInt(s)
            numbers[lastSpoken] = { oldest: -1, newest: turn }
            seenCount[lastSpoken] = 1
            //console.log(turn + ': ' + lastSpoken + '; ' + dump(numbers, lastSpoken))
            turn++
        }
        var nextNumber
        var nextOldest
        while (turn <= 30000000) {
            if (turn % 100000 == 0) {
                console.log(turn)
            }
            if (numbers[lastSpoken].oldest == -1) {
                //console.log(lastSpoken + ' is the first time that has been said; ' + dump(numbers, lastSpoken))
                nextNumber = 0
                nextOldest = numbers[nextNumber].newest
            } else {
                //console.log(lastSpoken + ' has been spoken before; ' + dump(numbers, lastSpoken))
                nextNumber = numbers[lastSpoken].newest - numbers[lastSpoken].oldest
                if (nextNumber in numbers) {
                    nextOldest = numbers[nextNumber].newest
                } else {
                    nextOldest = -1
                }
            }
            numbers[nextNumber] = { oldest: nextOldest, newest: turn}
            lastSpoken = nextNumber
            //console.log(turn + ': ' + nextNumber + '; ' + dump(numbers, nextNumber))
            turn++
        }
        console.log(lastSpoken)
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
