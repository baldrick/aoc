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
    var ingredients = aori[0].trim().split(' ').map(a => a.trim())
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

function dumpAllergens(a2i) {
    for (var allergen in a2i) {
        console.log(allergen + ': ' + JSON.stringify(a2i[allergen]))
    }
}

function removeFood(a2i, allergen) {
    var knownFood = a2i[allergen][0]
    console.log('Removing food ' + knownFood)
    for (var a in a2i) {
        console.log('Checking ' + a + ' (' + JSON.stringify(a2i[a]) + ') for ' + knownFood)
        if (a != allergen) {
            var i = a2i[a].indexOf(knownFood)
            console.log(knownFood + ' found in ' + a2i[a] + ' at ' + i)
            if (i != -1) {
                console.log('Removing ' + knownFood + ' from ' + a2i[a])
                a2i[a].splice(i, 1)
            }
        }
    }
}

function identifyAllergens(a2i) {
    for (var allergen in a2i) {
        if (a2i[allergen].length == 1) {
            // we know which food contains this allergen, remove it from all other foods
            removeFood(a2i, allergen)
        }
    }
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

        // Could make this more general but can't be ar*ed
        identifyAllergens(a2i)
        identifyAllergens(a2i)
        identifyAllergens(a2i)

        console.log('Ingredient for allergens')

        var sortedAllergens = []
        var n = 0
        for (var allergen in a2i) {
            sortedAllergens[n] = allergen
            n++
        }
        sortedAllergens.sort()

        var ingredients = []
        for (n = 0;  n < sortedAllergens.length;  n++) {
            console.log(sortedAllergens[n] + ': ' + a2i[sortedAllergens[n]])
            ingredients[n] = a2i[sortedAllergens[n]]
        }
        console.log('Sorted ingredients: ' + ingredients.join(','))
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
