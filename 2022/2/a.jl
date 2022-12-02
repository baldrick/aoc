f = open(ARGS[1], "r")

opponentCode = Dict("A" => "rock", "B" => "paper", "C" => "scissors")
myCode = Dict("X" => "rock", "Y" => "paper", "Z" => "scissors")

function score(opponentRPS, myRPS)
    if opponentRPS == myRPS
        return 3
    elseif opponentRPS == "rock" && myRPS == "paper"
        return 6
    elseif opponentRPS == "paper" && myRPS == "scissors"
        return 6
    elseif opponentRPS == "scissors" && myRPS == "rock"
        return 6
    end
    return 0
end

function shapeScore(shape)
    if shape == "rock"
        return 1
    elseif shape == "paper"
        return 2
    end
    return 3
end

totalScore = 0
while ! eof(f)
    line = readline(f)
    shapes = split(line, " ")
    opponent = opponentCode[shapes[1]] # A=rock, B=paper, C=scissors
    me = myCode[shapes[2]] # X=rock, Y=paper, Z=scissors
    # rock>scissors
    # paper>rock
    # scissors>paper
    roundScore = score(opponent, me)
    ss = shapeScore(me)
    println("$opponent plays $me for score of $roundScore + $ss")
    global totalScore += roundScore + ss
end  

println("$totalScore")
