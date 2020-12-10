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

function doit(filename) {
    readInput(filename, function(input) {
        prevJolt = 0
        diffs = []
        for (var i = 0;  i < 5;  i++) {
            diffs[i] = 0
        }
        diffs[3] = 1
        for (var line of input) {
            jolt = parseInt(line)
            diffs[jolt - prevJolt] = diffs[jolt - prevJolt] + 1
            prevJolt = jolt
        }
        console.log('diffs[1] = ' + diffs[1] + ', diffs[3] = ' + diffs[3] + ' *= ' + (diffs[1] * diffs[3]))
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
