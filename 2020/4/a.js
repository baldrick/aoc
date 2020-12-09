fs = require('fs');
os = require('os');
assert = require('assert');

function readPassports(file, cb) {
    var map = [];
    fs.readFile(file, 'ascii', function(err, data) {
        if (err) {
            return console.log(err);
        }
        lines = data.split(os.EOL);
        passports = []
        passport = ''
        n = 0
        for (let line of lines) {
            if (line.length > 0) {
                passport += ' ' + line
            } else {
                passports[n] = passport
                passport = ''
                n++
            }
        }
        console.log(n + ' passports read')
        cb(passports)
    });
}
/*
byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)

byr:1949 hgt:176cm pid:531868428 hcl:#cfa07d ecl:brn iyr:2014 eyr:2024
*/
fieldMap = []
fieldMap['byr'] = 0 // 1
fieldMap['iyr'] = 1 // 2
fieldMap['eyr'] = 2 // 4
fieldMap['hgt'] = 3 // 8
fieldMap['hcl'] = 4 // 16
fieldMap['ecl'] = 5 // 32
fieldMap['pid'] = 6 // 64
fieldMap['cid'] = 7 // 128

allFields = (1 << 7) - 1

function isValid(passport) {
    fields = passport.split(' ')
    fieldsSeen = 0
    for (let field of fields) {
        if (field.trim() != 0) {
            kv = field.trim().split(':')
            //console.log('fieldsSeen=' + fieldsSeen + ' field=.' + field.trim() + '. key -' + kv[0] + '- shift ' + fieldMap[kv[0]] + '(' + (1 << fieldMap[kv[0]]) + ')')
            fieldsSeen |= 1 << fieldMap[kv[0]]
        }
    }
    pp = fieldsSeen & allFields
    passportIsValid = (pp == allFields)
    console.log('isValid: ' + passport + ' - ' + fieldsSeen + ' (' + pp + ') ' + (passportIsValid ? 'valid' : 'not valid'))
    return passportIsValid
}

function main() {
    readPassports('input.txt', function(passports) {
        valid2 = 0
        for (let p of passports) {
            if (isValid(p) && isValid2(p)) {
                console.log('both valid')
                valid2++
            }
        }
        console.log(valid2 + ' valid passports')
    })
}

/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
*/
function isValid2(passport) {
    console.log('checking validity of ' + passport)
    fields = passport.split(' ')
    valid = true
    for (let field of fields) {
        if (field.trim().length != 0) {
            kv = field.trim().split(':')
            switch (kv[0]) {
                case 'byr':
                    valid = valid && checkYear(kv[1], 1920, 2002)
                    break
                case 'iyr':
                    valid = valid && checkYear(kv[1], 2010, 2020)
                    break
                case 'eyr':
                    valid = valid && checkYear(kv[1], 2020, 2030)
                    break
                case 'hgt':
                    valid = valid && checkHeight(kv[1])
                    break;
                case 'hcl':
                    valid = valid && (kv[1].length == 7) && (kv[1].match('#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]') != null)
                    break;
                case 'ecl':
                    valid = valid && checkEyeColour(kv[1])
                    break;
                case 'pid':
                    valid = valid && (kv[1].length == 9) && (kv[1].match('[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]') != null)
                    break;
                case 'cid':
                    break;
                default:
                    console.log('cannot deal with ' + kv[0])
                    break;
            }
            console.log(kv[0] + ' = ' + kv[1] + ' is ' + valid)
        }
    }
    console.log('isValid2: ' + passport + ' is ' + valid)
    return valid
}

function checkYear(yr, min, max) {
    valid = true
    valid = valid && (yr.length == 4);
    yrn = parseInt(yr)
    valid = valid && (yrn >= min) && (yrn <= max)
    console.log(' vvv = ' + yrn + ' - ' + valid)
    return valid
}

function checkHeight(hgt) {
    l = hgt.length
    units = hgt.substring(l - 2, l)
    size = parseInt(hgt.substring(0, l - 2))
    console.log('height ' + size + ' units = ' + units)
    switch (units) {
        case 'cm':
            return (size >= 150) && (size <= 193)
            break;
        case 'in':
            return (size >= 59) && (size <= 76)
            break;
        defaul:
            return false
    }
}

function checkEyeColour(ec) {
    switch (ec) {
        case 'amb':
        case 'blu':
        case 'brn':
        case 'gry':
        case 'grn':
        case 'hzl':
        case 'oth':
            return true
        default:
            return false
    }
}

main()
//test()

function test() {
    assert(checkEyeColour('amb'))
    assert(checkEyeColour('oth'))
    assert(checkHeight('150 cm'))
    assert(!checkHeight('149cm'))
    assert(checkHeight('193cm'))
    assert(!checkHeight('194cm'))
    assert(checkHeight(' 59 in'))
    assert(!checkHeight(' 58 in'))
    assert(checkHeight(' 76 in'))
    assert(!checkHeight(' 77 in'))
    assert(checkYear('1920', 1920, 2002))
    assert(!checkYear('1919', 1920, 2002))
    assert(checkYear('2002', 1920, 2002))
    assert(!checkYear('2003', 1920, 2002))
    assert(!checkYear('02002', 1920, 2002))
}
