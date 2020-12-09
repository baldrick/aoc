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

function processGroup(group) {
    console.log('processing group ' + group)
    var q = []
    for (a2z = 0;  a2z < 26;  a2z++) {
        q[a2z] = false
    }
    for (var c of group.split("")) {
        console.log('processing ' + c)
        if ((c.charCodeAt(0) >= 'a'.charCodeAt(0)) && (c.charCodeAt(0) <= 'z'.charCodeAt(0))) {
            console.log('setting index ' + (c.charCodeAt(0) - 'a'.charCodeAt(0)) + ' to true')
            q[c.charCodeAt(0) - 'a'.charCodeAt(0)] = true
        }
    }
    c = 0
    for (a2z = 0;  a2z < 26;  a2z++) {
        if (q[a2z]) c++
    }
    return c
}

function doit(filename) {
    readInput(filename, function(input) {
        group = ""
        count = 0
        for (var line of input) {
            if (line.length > 0) {
                group += line
            } else {
                count += processGroup(group)
                group = ""
            }
        }
        if (group.length > 0) {
            count += processGroup(group)
        }
        console.log("Total is " + count)
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
