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

function processInput(input) {
    console.log('processing ' + input)
    sp = 0
    acc = 0
    executed = []
    while (sp < input.length) {
        if (executed[sp]) {
            console.log('loop found acc = ' + acc)
            return false
        }
        if (input[sp].length == 0) {
            break
        }
        executed[sp] = true
        fields = input[sp].split(' ')
        cmd = fields[0]
        arg = fields[1]
        switch (cmd) {
            case 'acc':
                acc += parseInt(arg)
                sp++
                break
            case 'jmp':
                sp += parseInt(arg)
                break
            case 'nop':
                sp++
                break
            default:
                console.log('unhandled command -' + cmd + '- for ' + input[sp])
                break
        }
    }
    console.log('acc = ' + acc)
    return true
}

function countOccurrences(input, find) {
    return (input.join().match(new RegExp(find, 'g')) || []).length
}

function testMutation(input, from, to) {
    count = countOccurrences(input, from)
    for (mutation = 0;  mutation < count;  mutation++) {
        mutated = input.map((x) => x)
        found = 0
        for (i = 0;  i < mutated.length;  i++) {
            fields = mutated[i].split(' ')
            if (fields[0] == from) {
                found++
                if (found > mutation) {
                    mutated[i] = to + ' ' + fields[1]
                    break
                }
            }
        }
        console.log('processing mutation #' + mutation + ' of ' + from)
        if (processInput(mutated)) {
            return true
        }
    }
    return false
}

function doit(filename) {
    readInput(filename, function(input) {
        if (!testMutation(input, 'nop', 'jmp')) {
            testMutation(input, 'jmp', 'nop')
        }
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
