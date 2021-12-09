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

function readCards(input, player) {
    var c = []
    var i = 0
    var addCards = false
    for (var line of input) {
        if (addCards && line.length == 0) break
        if (addCards) {
            c[i] = parseInt(line)
            i++
        }
        if (line == player) {
            addCards = true
        }
    }
    return c
}

function dumpCards(i, c) {
    console.log('Player ' + i)
    for (var card of c) {
        console.log(card)
    }
}

function playing(p1, p2) {
    return p1.length > 0 && p2.length > 0
}

function playARound(p1, p2) {
    if (p1[0] > p2[0]) {
        winner(p1, p2)
    } else {
        winner(p2, p1)
    }
}

function winner(win, lose) {
    var losingCard = lose[0]
    lose.splice(0,1)
    var winningCard = win[0]
    win.splice(0,1)
    win.push(winningCard)
    win.push(losingCard)
}

function gameWinner(p) {
    var t = 0
    for (var i = 0;  i < p.length;  i++) {
        var m = p.length - i
        t += m * p[i]
        console.log(' += ' + m + ' * ' + p[m-1])
    }
    console.log('Total: ' + t)
}

function doit(filename) {
    readInput(filename, function(input) {
        var p1 = readCards(input, 'Player 1:')
        var p2 = readCards(input, 'Player 2:')
        dumpCards(1, p1)
        dumpCards(2, p2)
        var r = 0
        while (playing(p1, p2)) {
            playARound(p1, p2)
            r++
        }
        if (p1.length > 0) {
            gameWinner(p1)
        } else {
            gameWinner(p2)
        }
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
