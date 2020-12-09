fs = require('fs');
os = require('os');

fs.readFile('2.txt', 'ascii', function(err, data) {
    if (err) {
        return console.log(err);
    }
    lines = data.split(os.EOL);
    valid = 0
    for (let line of lines) {
        if (line.length > 0) {
            fields = line.split(" ");
            minmax = fields[0].split("-");
            min = parseInt(minmax[0]);
            max = parseInt(minmax[1]);
            letter = fields[1][0];
            pwd = fields[2];
            if (passwordIsValid(min, max, letter, pwd)) {
                valid++
            }
        }
    }
    console.log('There are ' + valid + ' valid passwords');
});

function passwordIsValid(min, max, letter, pwd) {
    occurrences = countOccurrences(letter, pwd);
    return occurrences >= min && occurrences <= max;
}

function countOccurrences(letter, pwd) {
    count = 0;  
    for (let c of pwd.split("")) {
        if (c == letter) {
            count++;
        }
    }
    return count;
}

/*
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
Each line gives the password policy and then the password. 
The password policy indicates the lowest and highest number of times a given letter must appear for the 
password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. 
The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. 
The first and third passwords are valid: they contain one a or nine c, both within the limits of their 
respective policies.
*/