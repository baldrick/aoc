fs = require('fs');
os = require('os');
assert = require('assert');

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

function getRow(line) {
    colcode = line.substring(0, 7)
    return decode(colcode, 'F', 'B', 0, 127)
}

function getCol(line) {
    rowcode = line.substring(7, 10)
    return decode(rowcode, 'L', 'R', 0, 7)
}

function avgpos(min, max) {
    return Math.round((max - min) / 2)
}

function decode(code, takeMinCode, takeMaxCode, min, max) {
    for (var c of code.split('')) {
        omin = min
        omax = max
        switch (c) {
            case takeMinCode:
                max = max - avgpos(min, max)
                break
            case takeMaxCode:
                min = min + avgpos(min, max)
                break
            default:
                console.log('Illegal code ' + c + ' in ' + code)
                break
        }
        //console.log('Took ' + c + ' from ' + omin + '-' + omax + ' to get ' + min + '-' + max)
    }
    if (min != max) {
        console.log(code + ' did not get to a single point: ' + min + ', ' + max)
    }
    return min
}

function report(seats) {
    maxTaken = 0
    for (seat = 0;  seat < seats.length;  seat++) {
        if (seats[seat]) {
            console.log('Seat ' + seat + ' is free')
        } else {
            maxTaken = seat
        }
    }
    console.log('Max seat taken is ' + maxTaken)
}

function getSeat(line) {
    return getRow(line) * 8 + getCol(line)
}

function main() {
    readInput('input.txt', function(input) {
        seats = []
        maxSeat = 128 * 8 + 8
        for (seat = 0;  seat < maxSeat;  seat++) {
            seats[seat] = true
        }
        for (let line of input) {
            seats[getSeat(line)] = false
        }
        report(seats)
    })
}

main()
//test()

function test() {
    assert(getSeat('FBFBBFFRLR') == 357)
    assert(getSeat('BFFFBBFRRR') == 567)
    assert(getSeat('FFFBBBFRRR') == 119)
    assert(getSeat('BBFFBBFRLL') == 820)
}
