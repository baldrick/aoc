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

function intersect(s1, s2) {
    var overlap = []
    var o = 0
    for (var i = 0;  i < s1.length;  i++) {
        if (s2.includes(s1[i])) {
            overlap[o] = s1[i]
            o++
        } else {
            console.log(s1[i] + ' not in ' + JSON.stringify(s2))
        }
    }
    return overlap
}

function mapAllergensToIngredients(a2i, i, line) {
    var aori = line.split('(')
    var ingredients = aori[0].trim().split(' ')
    for (var ingredient of ingredients) {
        //console.log('Adding ingredient ' + ingredient)
        if (i[ingredient] === undefined) {
            i[ingredient] = 1
        } else {
            i[ingredient]++
        }
    }
    var allergens = aori[1].substring(8, aori[1].length - 1).trim().split(',')
    console.log('allergens: ' + allergens)
    for (var allergen of allergens) {
        allergen = allergen.trim()
        if (a2i.hasOwnProperty(allergen)) {
            console.log('Finding intersect of ' + ingredients + ' and ' + a2i[allergen])
            a2i[allergen] = intersect(a2i[allergen], ingredients)
        } else {
            console.log('1. a2i[' + allergen + '] = ' + ingredients)
            a2i[allergen] = ingredients
            console.log('2. a2i[' + allergen + '] = ' + a2i[allergen])
        }
    }
    for (var allergen in a2i) {
        console.log(allergen + ': ' + JSON.stringify(a2i[allergen]))
    }
}

function removeIngredients(i, remove) {
    for (var r of remove) {
        i[r] = 0
    }
}

function dumpIngredients(i) {
    var t = 0
    for (var ingredient in i) {
        if (i[ingredient] > 0) {
            console.log(ingredient + ' ' + i[ingredient])
            t += i[ingredient]
        }
    }
    console.log('Total = ' + t)
}

function doit(filename) {
    readInput(filename, function(input) {
        var a2i = []
        var i = []
        for (var line of input) {
            mapAllergensToIngredients(a2i, i, line)
        }
        for (var allergen in a2i) {
            console.log('Removing ' + a2i[allergen])
            removeIngredients(i, a2i[allergen])
        }
        console.log('Ingredients that cannot contain an allergen:')
        dumpIngredients(i)
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
