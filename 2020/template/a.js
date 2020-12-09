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
