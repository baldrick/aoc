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

function calculateBusesInService(line) {
    var b = 0
    var busesInService = []
    for (var bus of line.split(',')) {
        if (bus != 'x') {
            busesInService[b] = parseInt(bus)
            b++
        }
    }
    return busesInService
}

function calculateLeaveTimes(buses, earliest) {
    var b = 0
    var leaveTimes = []
    for (var bus of buses) {
        leaveTimes[b] = Math.ceil(earliest / bus) * bus
        b++
    }
    return leaveTimes
}

function doit(filename) {
    readInput(filename, function(input) {
        var earliest = parseInt(input[0])
        var busesInService = calculateBusesInService(input[1])
        var nextLeaveTime = calculateLeaveTimes(busesInService, earliest)
        var closest = Number.MAX_SAFE_INTEGER
        var closestBus = -1
        for (var bus = 0;  bus < busesInService.length;  bus++) {
            if (nextLeaveTime[bus] < closest) {
                closest = nextLeaveTime[bus]
                closestBus = bus
            }
        }
        if (closestBus == -1) {
            console.log('oops!')
        } else {
            console.log('We will have to wait ' + busesInService[closestBus] * (nextLeaveTime[closestBus] - earliest))
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
