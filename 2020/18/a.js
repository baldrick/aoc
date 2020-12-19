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
            if (line[0] == '/' && line[1] == '/') continue
            input[l] = line
            l++
        }
        cb(input)
    });
}

function skipSpaces(line, meta) {
    while (line[meta.n] == ' ') { meta.n++ }
    return meta
}

function getNumber(line, meta) {
    var start = meta.n
    switch (line[meta.n]) {
        case '0':
        case '1':
        case '2':
        case '3':
        case '4':
        case '5':
        case '6':
        case '7':
        case '8':
        case '9':
            while ('0123456789'.indexOf(line[meta.n]) != -1) { meta.n++ }
            //console.log('parsing number from ' + line.substring(start, n))
            var v = parseInt(line.substring(start, meta.n))
            //console.log(v)
            return { v: v, n: meta.n }
        default:
            console.log('Unhandled item: ' + line[meta.n] + ' at ' + meta.n + ' in ' + line)
            process.exit(1)
    }
}

function getLeft(line, meta) {
    meta = skipSpaces(line, meta)
    if (meta.n >= line.length) return meta
    if (line[meta.n] == '(') return evaluate(line, { n: meta.n+1 })
    return getNumber(line, meta)
}

function getOperator(line, meta) {
    meta = skipSpaces(line, meta)
    if (meta.n >= line.length) {
        console.log('Found end of line before operator for ' + line)
        process.exit(1)
    }
    //console.log(line[meta.n])
    return { op: line[meta.n], n: meta.n+1 }
}

function getRight(line, meta) {
    meta = skipSpaces(line, meta)
    if (meta.n >= line.length) {
        console.log('Found end of line before right hand operand for ' + line)
        process.exit(1)
    }
    if (line[meta.n] == '(') return evaluate(line, { n: meta.n+1 })
    return getNumber(line, meta)
}

function evaluate(line, meta) {
    var left = getLeft(line, meta)
    return evaluateWithLeft(line, left)
}

function evaluateWithLeft(line, left) {
    if (left.n >= line.length) return left
    var operator = getOperator(line, left)
    if (operator.n >= line.length || operator.op == ')') {
        left.n++
        return left
    }
    var right = getRight(line, operator)
    return evaluateWithLeft(line, applyOperator(left, operator, right))
}

function applyOperator(left, operator, right) {
    var v
    switch (operator.op) {
        case '+':
            v = left.v + right.v
            break
        case '-':
            v = left.v - right.v
            break
        case '*':
            v = left.v * right.v
            break
        case '/':
            v = left.v / right.v
            break
        default:
            console.log('Unhandled operator ' + operator.op)
    }
    //console.log(left.v + operator.op + right.v + '=' + v)
    return { v: v, n: right.n }
}

function doit(filename) {
    readInput(filename, function(input) {
        var total = 0
        for (var line of input) {
            var v = evaluate(line, { n: 0 })
            total += v.v
            console.log(line + ' = ' + v.v)
            if (v.v == NaN) {
                process.exit(1)
            }
        }
        console.log('Total: ' + total)
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
