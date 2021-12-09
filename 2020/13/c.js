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

function getBuses(busInput) {
    var buses = []
    var b = 0
    for (var bus of busInput) {
        if (bus == 'x') {
            buses[b] = { bus: 0, offset: b }
        } else {
            buses[b] = { bus: parseInt(bus), offset: b }
        }
        b++
    } 
    return buses
}

function isInteger(f) {
    return Math.ceil(f) == f
}

function dumpBuses(buses) {
    for (var bus of buses) {
        console.log('bus: ' + bus.bus + ', offset: ' + bus.offset)
    }
}

function isDepartureTime(t, refBus, bus) {
    return isInteger((t-refBus.offset+bus.offset)/bus.bus)
}

function getAGoldStar(line, expected) {
    var buses = getBuses(line)

    /*
    first departs at t which is multiple of bus#
    second departs at t+1 which is multiple of bus# AND (multiple of t)+1

    bus 1 departs at t = bus1 * x
    bus 2 departs at t+1 => t = bus2 * y - 1


    17,x,13,19=3417

    t=3417
    3417/17 = 201
    3419/13 = 263
    3420/19 = 180

    

    */

    var bus
    //dumpBuses(buses)
    var firstBus = buses[0] // assumes first bus listed isn't x - true for all our inputs
    //console.log('first bus = ' + firstBus.bus)
    buses.sort(function(a, b) { return b.bus - a.bus });
    var t = 100000000000000
    //dumpBuses(buses)
    var refBus = buses[0]
    //console.log('first bus = ' + firstBus.bus)
    while (true) {
        t += buses[0].bus
        if (t % 1e7 == 0) console.log('checking t = ' + t)
        for (bus = 1;  bus < buses.length;  bus++) {
            if (!isDepartureTime(t, refBus, buses[bus])) break
        }
        if (bus == buses.length) {
            console.log('success!')
            return t - refBus.offset
        }
        if (t > 1.1*expected) return 0
        //if (t > 50000) return t
    }
/*

    while (true) {
        if (m % 1e6 == 0) {
            console.log('Testing m = ' + m)
        }
        t = firstBus.bus * m
        //console.log('testing m = ' + m + ', t: ' + t)
        for (bus = 0;  bus < buses.length;  bus++) {
            //console.log('bus: ' + buses[bus].bus + ', offset: ' + buses[bus].offset)
            if (buses[bus].bus == 0 || buses[bus].bus == firstBus.bus) {
                //console.log('continuing')
                continue
            }
            if (!isDepartureTime(t, buses[bus])) break
            //console.log('INTEGER t: ' + t + ' bus=' + buses[bus].bus + ' offset=' + buses[bus].offset + ' t/(t-offset) = ' + (t+buses[bus].offset)/t)
        }
        if (bus == buses.length) {
            //console.log('found it..........')
            return t
        }
        m++ // can we jump quicker?
    }
    */
}

function doit(filename) {
    readInput(filename, function(input) {
        for (var line of input) {
            var possibleTest=line.split('=')
            var buses=possibleTest[0].split(',')
            if (buses.length == 1) continue // skip first line of input
            var expected
            if (possibleTest.length > 1) {
                var expected = parseInt(possibleTest[1])
            }
            var t=getAGoldStar(buses, expected)
            if (possibleTest.length > 1) {
                console.log('t:' + t + ', expected ' + expected + ' test ' + (expected == t ? 'passed' : 'failed'))
            } else {
                console.log('t=' + t)
            }
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
        test('testb.txt')
    }
} else {
    main()
}
