import util

# distribution of three rolls
# key = sum of rolls, value = how many times that occurs
sum = {
    3:1,
    4:2,
    5:4,
    6:4,
    7:4,
    8:2,
    9:1
}

# points scored from board positions
scoreFrom = {

}

class Player:
    def __init__(self, name, start):
        self.name = name
        self.pos = start
        self.scores = {0:0}

    def play(self):
        for score, occurrences in self.scores:
            for roll, times in sum:
                newPos = self.pos + roll
                if newPos > 10: newPos = 1
                score += newPos
            

start on 5
land on
8 1 times
9 2 times
10 4 times
1 4 times
2 4 times
3 4 times
4 2 times
5 1 time

start on 8
land on
1 1 time   total score 9
2 2 times              10
3 4 times
4 4 times
5 4 times
6 2 times
7 1 time               15

start on 1
land on
4 1 time
5 2
6 4
7 4
8 4
9 2
10 1


1
10
..
..
..
n

take m steps such that sum of (landings ~%10) = 21