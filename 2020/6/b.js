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

function processGroup(group, groupSize) {
    console.log('processing group ' + group)
    var q = []
    for (a2z = 0;  a2z < 26;  a2z++) {
        q[a2z] = 0
    }
    for (var c of group.split("")) {
        console.log('processing ' + c)
        if ((c.charCodeAt(0) >= 'a'.charCodeAt(0)) && (c.charCodeAt(0) <= 'z'.charCodeAt(0))) {
            console.log('incrementing index ' + (c.charCodeAt(0) - 'a'.charCodeAt(0)))
            q[c.charCodeAt(0) - 'a'.charCodeAt(0)]++
        }
    }
    c = 0
    for (a2z = 0;  a2z < 26;  a2z++) {
        if (q[a2z] == groupSize) c++
    }
    console.log('group size ' + groupSize + ' ' + c + ' answers all yes')
    return c
}

function doit(filename) {
    readInput(filename, function(input) {
        group = ""
        groupSize = 0
        count = 0
        for (var line of input) {
            if (line.length > 0) {
                group += line
                groupSize++
            } else {
                count += processGroup(group, groupSize)
                group = ""
                groupSize = 0
            }
        }
        if (group.length > 0) {
            count += processGroup(group, groupSize)
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
