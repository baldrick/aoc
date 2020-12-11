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

const EMPTY = 'L'
const OCCUPIED = '#'
const FLOOR = '.'

function readRow(line) {
    row = []
    col = 0
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
    occupied = 0
    for (row = 0;  row < layout.length;  row++) {
        s = ""
        for (col = 0;  col < layout[row].length;  col++) {
            s += layout[row][col]
            if (layout[row][col] == OCCUPIED) {
                occupied++
            }
        }
        console.log(s)
    }
    console.log(occupied + ' seats are occupied')
}

function inBounds(layout, r, c) {
    return r >= 0 && c >= 0 && r < layout.length && c < layout[0].length
}

function visibleOccupiedDirection(layout, r, c, dr, dc) {
    cr = r + dr
    cc = c + dc
    while (inBounds(layout, cr, cc)) {
        if (layout[cr][cc] == OCCUPIED) {
            return 1
        } else if (layout[cr][cc] == EMPTY) {
            return 0
        }
        cr += dr
        cc += dc
    }
    return 0
}

function visibleOccupied(layout, r, c) {
    deltas = [ -1, 0, 1 ]
    neighbours = 0
    for (var dr of deltas) {
        for (var dc of deltas) {
            if (dr != 0 || dc != 0) {
                neighbours += visibleOccupiedDirection(layout, r, c, dr, dc)
            }
        }
    }
    return neighbours
}

function processLayout(layout, r, c) {
    neighbours = visibleOccupied(layout, r, c)
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
            if (neighbours >= 5) {
                return EMPTY
            } else {
                return OCCUPIED
            }
    }
}

function shuffleSeats(layout) {
    newLayout = []
    for (shuffleRow = 0;  shuffleRow < layout.length;  shuffleRow++) {
        newLayout[shuffleRow] = []
    }
    changed = 0
    for (shuffleRow = 0;  shuffleRow < layout.length;  shuffleRow++) {
        for (shuffleCol = 0;  shuffleCol < layout[shuffleRow].length;  shuffleCol++) {
            newLayout[shuffleRow][shuffleCol] = processLayout(layout, shuffleRow, shuffleCol)
            if (newLayout[shuffleRow][shuffleCol] != layout[shuffleRow][shuffleCol]) {
                changed++
            }
        }
    }
    return { layout: newLayout, changed: changed }
}

function doit(filename) {
    readInput(filename, function(input) {
        layout = []
        xrow = 0
        for (var line of input) {
            layout[xrow] = readRow(line)
            console.log('set layout[' + xrow + '] to ' + layout[xrow])
            xrow++
        }
        dump('initial', layout)
        newLayout = { layout: layout, changed: 1}
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
