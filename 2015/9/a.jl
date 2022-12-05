function main()
    f = open(ARGS[1], "r")
    g = simple_inclist(5)
    dists = []
    while ! eof(f)
        line = readline(f)
        (l1, l2, d) = decode(line)
        add_edge!(g, l1, l2)
        push!(dists, d)
    end
    r = dijkstra_shortest_paths(g, dists, 1)
    println("$r")
end

function decode(line)
    # London to Belfast = 518
    r = r"([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)"
    m = match(r, line)
    return m[1], m[2], parse(Int, m[3])
end

main()
