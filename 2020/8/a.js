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
        sp = 0
        acc = 0
        executed = []
        while (true) {
            if (executed[sp]) {
                console.log('acc = ' + acc)
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
                    console.log('unhandled command ' + cmd)
                    break
            }
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
