using Graphs, SimpleWeightedGraphs

function main()
    f = open(ARGS[1], "r")
    g = SimpleWeightedGraph(7)
    start = -1
    dists = zeros(20,20)
    while ! eof(f)
        line = readline(f)
        (l1, l2, d) = decode(line)
        dists[l1, l2] = d
        add_edge!(g, l1, l2, d)
        if start == -1
            start = l1
        end
    end
    r = desopo_pape_shortest_paths(g, start) #prim_mst(g)
    println(r)
    return
    d = 0
    for e in r
        d += reportEdge(dists, e)
    end
    println("Total distance = $d")
end

function decode(line)
    # London to Belfast = 518
    r = r"([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)"
    m = match(r, line)
    return locationId(m[1]), locationId(m[2]), parse(Int, m[3])
end

locations = Dict()
revLoc = Dict()

function locationId(loc)
    if !haskey(locations, loc)
        id = length(locations)+1
        locations[loc] = id
        revLoc[id] = loc
    end
    return locations[loc]
end

function reportEdge(dists, edge)
    e = "$edge"
    r = r"Edge ([0-9]+) => ([0-9]+)"
    m = match(r, e)
    m1 = parse(Int, m[1])
    m2 = parse(Int, m[2])
    l1 = revLoc[m1]
    l2 = revLoc[m2]
    d = dists[m1, m2]
    println("$l1 -> $l2 = $d")
    return d
end

main()
