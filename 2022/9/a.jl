mutable struct pos
    x::Int
    y::Int
end

function main()
    grid = falses(1000, 1000)
    head = pos(500,500)
    tail = pos(500,500)
    visit(grid, tail)
    f = open(ARGS[1], "r")
    while ! eof(f)
        line = readline(f)
        (direction, amount) = decode(line)
        move(grid, head, tail, direction, amount)
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

function move(grid, head, tail, direction, amount)
    println("moving $direction $amount")
    for m in 1:amount
        moveOne(grid, head, tail, direction)
    end
end

function moveOne(grid, head, tail, direction)
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
    maybeMoveTail(grid, head, tail)
end

function maybeMoveTail(grid, head, tail)
    dx = head.x - tail.x
    dy = head.y - tail.y
    if abs(dx) <= 1 && abs(dy) <= 1
    # if dx == 0 && dy == 0
    #     # Head & tail overlap, do nothing.
    #     return
    # elseif (abs(dx) <= 1 && dy == 0) || (abs(dy) <= 1 && dx == 0)
    #     # Head & tail are touching in direct line.
        return
    end
    mdx = dx > 0 ? 1 : dx < 0 ? -1 : 0
    mdy = dy > 0 ? 1 : dy < 0 ? -1 : 0
    #println("head at $head, tail at $tail, moving tail by $mdx,$mdy")
    tail.x += mdx
    tail.y += mdy
    visit(grid, tail)
end

function visit(grid, tail)
    grid[tail.x, tail.y] = true
end

main()
