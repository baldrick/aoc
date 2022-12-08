function main()
    g = readGrid()
    lx = length(g)
    ly = length(g[1])
    seen = falses(lx,ly)
    display(g)

    #println("doing 1:$lx")
    for x in 1:lx
        h = -1
        for y in 1:ly
            #println("$x,$y=$(g[x][y]), h=$h")
            if g[x][y] > h
                h = g[x][y]
                seen[x,y] = true
            end
        end
        #display(seen)
        h = -1
        for y in ly:-1:1
            #println("$x,$y=$(g[x][y]), h=$h")
            if g[x][y] > h
                h = g[x][y]
                seen[x,y] = true
            end
        end
        #display(seen)
    end

    #println("doing $lx:1, 1:$ly")
    for y in 1:ly
        h = -1
        for x in 1:lx
            #println("$x,$y=$(g[x][y]), h=$h")
            if g[x][y] > h
                h = g[x][y]
                seen[x,y] = true
            end
        end
        #display(seen)
        h = -1
        for x in lx:-1:1
            #println("$x,$y=$(g[x][y]), h=$h")
            if g[x][y] > h
                h = g[x][y]
                seen[x,y] = true
            end
        end
        #display(seen)
    end
    c = 0
    for x in 1:lx
        for y in 1:ly
            if seen[x,y]
                c += 1
            end
        end
    end
    println(c)
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

main()
