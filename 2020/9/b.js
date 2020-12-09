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

function findMinMax(numbers, start, end) {
    min = Number.MAX_SAFE_INTEGER
    max = Number.MIN_SAFE_INTEGER
    for (n = start;  n <= end;  n++) {
        if (numbers[n] < min) {
            min = numbers[n]
        }
        if (numbers[n] > max) {
            max = numbers[n]
        }
    }
    console.log('min = ' + min + ', max = ' + max + ', sum = ' + (min+max))
}

// brute force ;-)
function findContinguousSum(numbers, target) {
    start = 0
    end = 1
    for (n = start;  n < numbers.length;  n++) {
        total = numbers[n]
        for (m = n + 1;  m <= numbers.length && total < target;  m++) {
            total += numbers[m]
        }
        if (total == target) {
            m--
            console.log(n + '-' + m+ ' sum to ' + target)
            findMinMax(numbers, n, m)
            return
        }
    }
}

function processInput(input, target) {
    numbers = parseToNumber(input)
    findContinguousSum(numbers, target)
}

function doit(filename, target) {
    readInput(filename, function(input) {
        processInput(input, target)
    })
}

function main() {
    doit('input.txt', 248131121)
}

function test() {
    doit('test.txt', 127)
}

if (process.argv.length > 2 && process.argv[2] == '-t') {
    test()
} else {
    main()
}
