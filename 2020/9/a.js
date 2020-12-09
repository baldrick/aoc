fs = require('fs')
os = require('os')
assert = require('assert')

function readInput(filename, cb) {
    var input = [];
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

function parseToNumber(input) {
    numbers = []
    index = 0
    for (var i of input) {
        numbers[index] = parseInt(i)
        index++
    }
    return numbers
}

// brute force ;-)
function canFindSum(numbers, start, end, target) {
    for (n = start;  n < end;  n++) {
        for (m = n + 1;  m <= end;  m++) {
            if (numbers[n] + numbers[m] == target) {
                return true
            }
        }
    }
    return false
}

function processInput(input, preamble, memorySize) {
    numbers = parseToNumber(input)
    start = 0
    end = preamble - 1
    for (test = preamble;  test < input.length;  test++) {
        if (!canFindSum(numbers, start, end, numbers[test])) {
            console.log('cannot sum ' + numbers[test])
            return
        }
        start++
        end++
    }
}

function doit(filename, preamble) {
    readInput(filename, function(input) {
        processInput(input, preamble, preamble)
    })
}

function main() {
    doit('input.txt', 25)
}

function test() {
    doit('test.txt', 5)
}

if (process.argv.length > 2 && process.argv[2] == '-t') {
    test()
} else {
    main()
}
