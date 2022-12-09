mutable struct pos
    x::Int
    y::Int
end

function main()
    grid = falses(1000, 1000)
    head = pos(500,500)
    tails = []
    for t in 1:9
        push!(tails, pos(500,500))
    end
    visit(grid, head) # they're all in one place...
    f = open(ARGS[1], "r")
    while ! eof(f)
        line = readline(f)
        (direction, amount) = decode(line)
        move(grid, head, tails, direction, amount)
    end
    visited = 0
    for v in grid
        if v
            visited += 1
        end
    end
    println(visited)
end

function decode(line)
    s = split(line)
    return s[1], parse(Int, s[2])
end

function move(grid, head, tails, direction, amount)
    println("moving $direction $amount")
    for m in 1:amount
        moveOne(grid, head, tails, direction)
    end
end

function moveOne(grid, head, tails, direction)
    dx = 0
    dy = 0
    if direction == "R"
        dx = 1
    elseif direction == "L"
        dx = -1
    elseif direction == "U"
        dy = 1
    elseif direction == "D"
        dy = -1
    end
    head.x += dx
    head.y += dy
    maybeMoveTails(grid, head, tails)
end

function maybeMoveTails(grid, head, tails)
    prev = head
    for tail in tails
        maybeMoveTail(grid, prev, tail)
        prev = tail
    end
    visit(grid, tails[9])
end

function maybeMoveTail(grid, head, tail)
    dx = head.x - tail.x
    dy = head.y - tail.y
    if abs(dx) <= 1 && abs(dy) <= 1
        return
    end
    mdx = dx > 0 ? 1 : dx < 0 ? -1 : 0
    mdy = dy > 0 ? 1 : dy < 0 ? -1 : 0
    #println("head at $head, tail at $tail, moving tail by $mdx,$mdy")
    tail.x += mdx
    tail.y += mdy
end

function visit(grid, tail)
    grid[tail.x, tail.y] = true
end

main()
