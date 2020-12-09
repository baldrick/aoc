fs = require('fs');
os = require('os');

function readMap(mapfile, cb) {
    var map = [];
    fs.readFile(mapfile, 'ascii', function(err, data) {
        if (err) {
            return console.log(err);
        }
        lines = data.split(os.EOL);
        row = 0
        for (let line of lines) {
            if (line.length > 0) {
                col = 0
                map[row] = []
                for (let ch of line) {
                    map[row][col] = (ch == '#')
                    col++
                }
                row++
            }
        }
        cb(map)
    });
}

function traverse(map, rowdelta, coldelta) {
    col = 0
    row = 0
    trees = 0
    while (row < map.length) {
        if (map[row][col]) trees++;
        row += rowdelta
        col = (col + coldelta) % map[0].length;
    }
    return trees
}

readMap('map.txt', function(map) {
    console.log('map is ' + map[0].length + ' wide by ' + map.length + ' long')
    trees = traverse(map, 1, 1)
    trees *= traverse(map, 1, 3)
    trees *= traverse(map, 1, 5)
    trees *= traverse(map, 1, 7)
    trees *= traverse(map, 2, 1)
    console.log('tree multiple is ' + trees)
})
/*
Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
*/