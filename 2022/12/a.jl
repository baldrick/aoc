using Graphs

function main()
    (g, width) = create_graph()
    f = open(ARGS[1], "r")
    y = 0
    nodes = []
    startPos = zeros(Int, 0)
    endPos = 0
    while !eof(f)
        line = readline(f)
        for c in line
            push!(nodes, c == 'S' ? 'a' : c == 'E' ? 'z' : c)
            if c == 'S' || c == 'a'
                push!(startPos, length(nodes))
            end
            if c == 'E'
                endPos = length(nodes)
            end
        end
    end

    for node in 1:length(nodes)
        maybe_add_edge(g, width, nodes, node, -1, 0)
        maybe_add_edge(g, width, nodes, node, 1, 0)
        maybe_add_edge(g, width, nodes, node, 0, -width)
        maybe_add_edge(g, width, nodes, node, 0, width)
    end
    display(startPos)
    #h = SimpleDiGraphFromIterator(edges(g))
    #println(collect(edges(h)))
    r = dijkstra_shortest_paths(g, startPos)
    display(r.dists[endPos])
end

function create_graph()
    f = open(ARGS[1], "r")
    line = readline(f)
    x = length(line)
    y = 1
    while !eof(f)
        readline(f)
        y += 1
    end
    #println("Creating graph $x by $y")
    return (SimpleDiGraph{Int64}(x*y), x)
end

function maybe_add_edge(g, width, nodes, node, dx, dy)
    if node+dx <= 0 || node+dy <= 0 || floor((node-1)/width) != floor((node+dx-1)/width) || node+dy > length(nodes)
        #println("$(node+dx),$(node+dy) is out of bounds")
        return
    end
    #println("Checking $node to $(node+dx+dy)")
    if close_enough(nodes[node], nodes[node+dx+dy])
        #println("Adding edge from $node to $(node+dx+dy)")
        add_edge!(g, node, node+dx+dy)
    end
end

function close_enough(from, to)
    ce = to - from <= 1
    #println("$to is $(to-from) away from $from - $(ce ? "" : "not ") close enough")
    return ce
end

main()
