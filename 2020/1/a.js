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
        remainder = target - number
        if (seenNumbers[remainder]) {
            console.log(remainder + ' x ' + number + ' = ' + (remainder * number));
        }
        seenNumbers[number] = true
    }
});