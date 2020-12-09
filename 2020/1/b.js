fs = require('fs');
os = require('os');

fs.readFile('1.txt', 'ascii', function(err, data) {
    if (err) {
        return console.log(err);
    }
    lines = data.split(os.EOL);
    target = 2020
    seenNumbers = []
    for (let line of lines) {
        number = parseInt(line)
        seenNumbers[number] = true
    }
    for (var first in seenNumbers) {
        if (seenNumbers.hasOwnProperty(first)) {
            //console.log('first = ' + first);
            findTarget(first, target, seenNumbers)
        }
    }
});

function findTarget(first, target, seenNumbers) {
    secondaryTarget = target - first
    for (var second in seenNumbers) {
        if (seenNumbers.hasOwnProperty(second)) {
            remainder = secondaryTarget - second
            //console.log('second = ' + second + ', looking for ' + remainder);
            if (remainder > 0) {
                if (seenNumbers[remainder]) {
                    console.log(first + ' x ' + second + ' x ' + remainder + ' = ' + (first * second * remainder));
                }
            }
        }
    }
}