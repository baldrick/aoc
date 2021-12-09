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

const LETTER = 'LETTER'
const OTHER_RULES = 'OTHER RULES'

function readRules(input) {
    var rules = []
    for (var line = 0;  line < input.length;  line++) {
        if (input[line].length == 0) return { rules: rules, line: line+1 /* skip the blank line */ }
        var fields = input[line].split(' ')
        var ruleNumber = parseInt(fields[0].split(':')[0])
        var rule
        if (fields[1][0] == '"') {
            rule = { type: LETTER, letter: fields[1][1] }
        } else {
            var ruleSet = []
            var n = 0
            var s = 0
            ruleSet[0] = []
            for (var field = 1;  field < fields.length;  field++) {
                if (fields[field] == '|') {
                    n++
                    ruleSet[n] = []
                    s = 0
                } else {
                    ruleSet[n][s] = parseInt(fields[field])
                    s++
                }
            }
            rule = { type: OTHER_RULES, rules: ruleSet, validInputs: [] }
        }
        rules[ruleNumber] = rule
    }
}

function dumpRules(rules) {
    for (var rule = 0;  rule < rules.length;  rule++) {
        console.log(rule + ': ' + JSON.stringify(rules[rule]))
    }
}

function calculateAllValidInputs(rules, rule) {
    switch (rules[rule].type) {
        case LETTER:
            return rules[rule].letter
        case OTHER_RULES:
            // e.g. subRuleArray = [1, 3] | [3, 4]
            // = [ab] | [b[a|b]]
            // = ab | ba | bb
            return calculateSubruleValidInputs()
            var subruleArray = rules[rule].rules
            var subruleValidInputs = []
            for (var san = 0;  san < subruleArray.length;  san++) {
                for (var subrule = 0;  subrule < subruleArray[san].length;  subrule++) {
                     var validInputs = calculateAllValidInputs(rules, subrule)
                     for (var vi of validInputs) {
                        subruleValidInputs[s] = 
                     }
                }
            }
            return subruleValidInputs
        default:
            console.log('Invalid rule type ' + rules[rule].type + ' for rule #' + rule)
            process.exit(1)
    }
}

function validate(rules, line) {
    console.log(line)
}

// validInputs[0] = map of valid inputs, e.g. validInputs[0]['ababab'] = true

function doit(filename) {
    readInput(filename, function(input) {
        var rules = readRules(input)
        dumpRules(rules.rules)
        validInputs = calculateAllValidInputs(rules.rules, 0)
        //for (var line = rules.line;  line < input.length;  line++) {
        //    validate(rules, input[line])
        //}
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
