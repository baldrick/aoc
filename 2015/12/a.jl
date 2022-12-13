function main()
    f = open(ARGS[1], "r")
    while !eof(f)
        line = readline(f)
        println("=== $line ===")
        total = get_total(line)
        println(total)
        total -= reds(line)
        println("ex reds = $total")
    end
end

function get_total(line)
    line = replace(line, r"[,:{}\"\[\]]" => " ")
    elements = split(line, " ")
    total = 0
    for el in elements
        r = r"([\-]*[0-9]+).*"
        m = match(r, el)
        if m == nothing || m[1] == "-"
            continue
        end
        println("el=$el, m[1]=$(m[1])")
        total += parse(Int, m[1])
    end
    return total
end

function reds(line)
    r = r".*{(.*)}.*"
    m = match(r, line)
    if m == nothing
        return 0
    end
    exc_internal_objects = replace(line, r"{.*}" => "")
    red = r".*\":\"red\".*"
    if occursin(r, exc_internal_objects)
        m = match(line, r".*{(.*)}.*")
        return -get_total(m[1])
    end
NEED TO TAKE INTO ACCOUNT {...}, {...} etc. structures
    m = match(line, r".*{(.*)}.*")
    return reds(m[1])
end

main()
