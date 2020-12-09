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
            first = parseInt(minmax[0]);
            second = parseInt(minmax[1]);
            letter = fields[1][0];
            pwd = fields[2];
            if (passwordIsValid(first, second, letter, pwd)) {
                console.log(pwd + ' is valid');
                valid++;
            } else {
                console.log(pwd + ' is not valid');
            }
        }
    }
    console.log('There are ' + valid + ' valid passwords');
});

function passwordIsValid(first, second, letter, pwd) {
    fm = pwd[first - 1] == letter;
    console.log('character ' + first + ' is ' + letter + ' in ' + pwd + ' ... ' + (fm ? 'it is' : 'it is not'));
    sm = pwd[second - 1] == letter;
    console.log('checking character ' + second + ' is ' + letter + ' in ' + pwd + ' ... ' + (sm ? 'it is' : 'it is not'));
    validpwd = (fm && !sm) || (!fm && sm)
    console.log(pwd + ' ' + (valid ? 'is valid' : 'is not valid'));
    return validpwd
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
