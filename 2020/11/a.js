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
            input[l] = line
            l++
        }
        cb(input)
    });
}

const EMPTY = 'L'
const OCCUPIED = '#'
const FLOOR = '.'

function readRow(line) {
    var row = []
    var col = 0
    console.log('Adding ' + line)
    for (var c of line.split('')) {
        if (c == 'L') {
            row[col] = EMPTY
        } else {
            row[col] = FLOOR
        }
        col++
    }
    return row
}

function dump(msg, layout) {
    console.log('----------------------------')
    console.log(msg)
    console.log('----------------------------')
    console.log('layout is ' + layout.length + ' x ' + layout[0].length)
    var occupied = 0
    for (var row = 0;  row < layout.length;  row++) {
        var s = ""
        for (var col = 0;  col < layout[row].length;  col++) {
            s += layout[row][col]
            if (layout[row][col] == OCCUPIED) {
                occupied++
            }
        }
        console.log(s)
    }
    console.log(occupied + ' seats are occupied')
}

function adjacentOccupied(layout, r, c) {
    var neighbours = 0
    for (var row = r - 1;  row <= r + 1;  row++) {
        for (var col = c - 1;  col <= c + 1;  col++) {
            if ((row == r && col == c)
                || row < 0 || row >= layout.length
                || col < 0 || col >= layout[0].length) {
                // don't check current seat or ones that are out of bounds
            } else {
                if (layout[row][col] == OCCUPIED) {
                    neighbours++
                }
            }
        }
    }
    return neighbours
}

/*
If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
*/
function processLayout(layout, r, c) {
    var neighbours = adjacentOccupied(layout, r, c)
    switch (layout[r][c]) {
        case FLOOR:
            return FLOOR
        case EMPTY:
            if (neighbours == 0) {
                return OCCUPIED
            } else {
                return EMPTY
            }
        case OCCUPIED:
            if (neighbours >= 4) {
                return EMPTY
            } else {
                return OCCUPIED
            }
    }
}

function shuffleSeats(layout) {
    var newLayout = []
    for (var row  = 0;  row < layout.length;  row++) {
        newLayout[row] = []
    }
    var changed = 0
    for (var row = 0;  row < layout.length;  row++) {
        for (var col = 0;  col < layout[row].length;  col++) {
            newLayout[row][col] = processLayout(layout, row, col)
            if (newLayout[row][col] != layout[row][col]) {
                changed++
            }
        }
    }
    return { layout: newLayout, changed: changed }
}

function doit(filename) {
    readInput(filename, function(input) {
        var layout = []
        var row = 0
        for (var line of input) {
            layout[row] = readRow(line)
            console.log('set layout[' + row + '] to ' + layout[row])
            row++
        }
        dump('initial', layout)
        var newLayout = { layout: layout, changed: 1}
        while (newLayout.changed != 0) {
            newLayout = shuffleSeats(newLayout.layout)
            console.log('Shuffle changed ' + newLayout.changed + ' seats')
        }
        dump('complete', newLayout.layout)
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
