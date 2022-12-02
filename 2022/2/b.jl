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

function cheatScore(opponentShape, code) # returns round score, shape score
    if code == "X" # Lose
        if opponentShape == "rock"
            return 0, shapeScore("scissors")
        elseif opponentShape == "paper"
            return 0, shapeScore("rock")
        end
        return 0, shapeScore("paper")
    elseif code == "Y" # Draw
        return 3, shapeScore(opponentShape)
    end
    # Win
    if opponentShape == "rock"
        return 6, shapeScore("paper")
    elseif opponentShape == "paper"
        return 6, shapeScore("scissors")
    end
    return 6, shapeScore("rock")
end

totalScore = 0
while ! eof(f)
    line = readline(f)
    shapes = split(line, " ")
    opponent = opponentCode[shapes[1]]
    roundScore, ss = cheatScore(opponent, shapes[2])
    global totalScore += roundScore + ss
end  

println("$totalScore")
