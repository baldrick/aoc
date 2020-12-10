var fs = require('fs')
var os = require('os')

function readInput(filename, cb) {
    var input = [];
    input[0] = 0
    fs.readFile(filename, 'ascii', function(err, data) {
        if (err) {
            return console.log(err);
        }
        var lines = data.split(os.EOL);
        var l = 1
        for (var line of lines) {
            input[l] = parseInt(line)
            l++
        }
        input[l] = input[l-1] + 3
        cb(input)
    });
}

function visit(paths, input, node) {
    if (node >= input.length - 1) return true
    for (var i = 1;  input[node+i] - input[node] <= 3;  i++) {
        if (paths[node+i] == 0 && visit(paths, input, node+i)) {
            paths[node] = paths[node] + 1
        }
        paths[node] += paths[node+i]
    }
    return false
}

function doit(filename) {
    readInput(filename, function(input) {
        paths = []
        for (i = 0;  i < input.length;  i++) {
            paths[i] = 0
        }
        visit(paths, input, 0)
        console.log('There are ' + paths[0] + ' ways to connect the adapters')
    })
}

function main() {
    doit('input.txt')
}

function test() {
    testfile = 'test.txt'
    if (process.argv.length > 3) testfile = process.argv[3]
    doit(testfile)
}

if (process.argv.length > 2 && process.argv[2] == '-t') {
    test()
} else {
    main()
}
