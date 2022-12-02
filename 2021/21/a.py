
class Player:
    def __init__(self, name, start):
        self.name = name
        self.pos = start
        self.score = 0

    def play(self, die):
        roll = die.roll(3)
        self.pos = (self.pos + roll) % 10
        if self.pos == 0: self.pos = 10
        self.score += self.pos
        print(f"Player {self.name} rolls {roll} and moves to space {self.pos} for a total score of {self.score}")
        return self.score >= 1000

class Die:
    def __init__(self, start, end):
        self.start = start
        self.end = end
        self.current = start
        self.rolls = 0

    def roll(self, times):
        sum = 0
        for _ in range(0, times):
            self.rolls += 1
            sum += self.current
            self.current += 1
            if self.current > self.end:
                self.current = self.start
        return sum

d = Die(1,100)
p1 = Player("1", 5)
p2 = Player("2", 8)

while True:
    if p1.play(d):
        print(f"Player 1 wins with {p1.score} points after {d.rolls} rolls, p2 had {p2.score} points, p2*d = {d.rolls*p2.score}")
        break
    if p2.play(d):
        print(f"Player 2 wins with {p2.score} points after {d.rolls} rolls, p1 had {p1.score} points, p1*d = {d.rolls*p1.score}")
        break
