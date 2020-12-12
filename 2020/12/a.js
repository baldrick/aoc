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

function processMovement(ship, line) {
    var cmd = line[0]
    var size = parseInt(line.substring(1))
    switch (cmd) {
        case 'N':
            ship.y += size
            break
        case 'E':
            ship.x += size
            break
        case 'W':
            ship.x -= size
            break
        case 'S':
            ship.y -= size
            break 
        case 'L':
            ship.direction = (ship.direction - size) % 360
            if (ship.direction < 0) ship.direction += 360
            break
        case 'R':
            ship.direction = (ship.direction + size) % 360
            if (ship.direction < 0) ship.direction += 360
            break
        case 'F':
            switch (ship.direction) {
                case 0:
                    ship.y += size
                    break
                case 90:
                    ship.x += size
                    break
                case 180:
                    ship.y -= size
                    break
                case 270:
                    ship.x -= size
                    break
                default:
                    console.log('Unhandled direction ' + ship.direction)
                    break
            }
            break
        default:
            console.log('Unhandled cmd ' + cmd)
            break
    }
}

function doit(filename) {
    readInput(filename, function(input) {
        var ship = { direction: 90, x: 0, y: 0 }
        for (var line of input) {
            processMovement(ship, line)
        }
        var m = Math.abs(ship.x) + Math.abs(ship.y)
        console.log('Manhattan distance = ' + m)
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
