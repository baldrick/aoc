function main()
    g = readGrid()
    lx = length(g)
    ly = length(g[1])
    display(g)

    max = 0
    for x in 1:lx
        for y in 1:ly
            s = scenicScore(g, x, y)
            if s > max
                max = s
            end
        end
    end
    println(max)
end

function readGrid()
    g = []
    f = open(ARGS[1], "r")
    numLines = 0
    numCols = 0
    while ! eof(f)
        numLines += 1
        line = readline(f)
        for c in line
            append!(g, parse(Int, c))
        end
        numCols = length(line)
    end
    grid = reshape(g, numLines, numCols)
    return slicematrix(grid)
end

function slicematrix(A::AbstractMatrix)
    return [A[i, :] for i in 1:size(A,1)]
end

function scenicScore(g, x, y)
    v1 = 0
    h = g[x][y]
    for i in x-1:-1:1
        v1 += 1
        if g[i][y] >= h
            break
        end
    end

    v2 = 0
    for i in x+1:length(g)
        v2 += 1
        if g[i][y] >= h
            break
        end
    end

    v3 = 0
    for i in y-1:-1:1
        v3 += 1
        if g[x][i] >= h
            break
        end
    end

    v4 = 0
    for i in y+1:length(g[1])
        v4 += 1
        if g[x][i] >= h
            break
        end
    end

    #println("$x,$y=$h: $v1*$v2*$v3*$v4")
    return v1 * v2 * v3 * v4
end

main()
