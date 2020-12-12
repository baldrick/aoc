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

function move(thing, cmd, size) {
    switch (cmd) {
        case 'N':
            thing.y += size
            break
        case 'E':
            thing.x += size
            break
        case 'W':
            thing.x -= size
            break
        case 'S':
            thing.y -= size
            break 
        default:
            console.log('Unhandled cmd ' + cmd)
            break
    }
}

function rotate(ship, waypoint, cmd) {
    /* when ship is at 0,0
    10, 1 rotate 90 left => -1, 10   = swap, -x
    10, 1 rotate 90 right => 1, -10  = swap, -y

    -1, 10 rotate 90 left => -10, -1 = swap, -x
    -1, 10 rotate 90 right => 10, 1  = swap, -y

    -10, -1 rotate 90 left => 1, -10 = swap, -x
    -10, -1 rotate 90 right => -1, 10= swap, -y

    1, -10 rotate 90 left 10, 1      = swap, -x
    1, -10 rotate 90 right -10, -1   = swap, -y
    */

    var relativeWaypoint = { x: waypoint.x - ship.x, y: waypoint.y - ship.y }

    var t = relativeWaypoint.x
    relativeWaypoint.x = relativeWaypoint.y
    relativeWaypoint.y = t
    if (cmd == 'L') {
        relativeWaypoint.x *= -1
    } else {
        relativeWaypoint.y *= -1
    }

    waypoint.x = relativeWaypoint.x + ship.x
    waypoint.y = relativeWaypoint.y + ship.y
}

function processMovement(ship, waypoint, line) {
    var cmd = line[0]
    var size = parseInt(line.substring(1))
    switch (cmd) {
        case 'N':
        case 'E':
        case 'W':
        case 'S':
            move(waypoint, cmd, size)
            break 
        case 'L':
            while (size > 0) {
                rotate(ship, waypoint, cmd)
                size -= 90
            }
            break
        case 'R':
            while (size > 0) {
                rotate(ship, waypoint, cmd)
                size -= 90
            }
            break
        case 'F':
            var dx = size * (waypoint.x - ship.x)
            var dy = size * (waypoint.y - ship.y)
            move(ship, 'N', dy)
            move(ship, 'E', dx)
            move(waypoint, 'N', dy)
            move(waypoint, 'E', dx)
            break
        default:
            console.log('Unhandled cmd ' + cmd)
            break
    }
}

function dump(line, ship, waypoint) {
    console.log(line + ' -> ship ' + ship.x + ',' + ship.y + '; waypoint ' + waypoint.x + ',' + waypoint.y)
}

function doit(filename) {
    readInput(filename, function(input) {
        var waypoint = { x: 10, y: 1 }
        var ship = { x: 0, y: 0 }
        dump('start', ship, waypoint)
        for (var line of input) {
            processMovement(ship, waypoint, line)
            dump(line, ship, waypoint)
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
