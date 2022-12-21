mutable struct Dependency
    dependentOn::Array{String}
    dependentOnMe::Array{String}
    operation::String
    inputValues::Array{Int}
    n::Union{Int, Nothing}
end

function main()
    f = open(ARGS[1], "r")
    monkeys = Dict()
    while !eof(f)
        monkeyName, n, dependencies, op = decode(readline(f))
        if length(dependencies) > 0
            if !haskey(monkeys, monkeyName)
                monkeys[monkeyName] = Dependency([], [], op, [0,0], nothing)
            end
            for dependency in dependencies
                push!(monkeys[monkeyName].dependentOn, dependency)
            end
        else
            monkeys[monkeyName] = Dependency([], [], "", [0,0], n)
        end
    end
    for monkey in monkeys
        monkey.second.dependentOnMe = findWhoCares(monkeys, monkey.first)
    end
    root = monkeys["root"]
    while root.n == nothing
        println("root has $(root.n)")
        for monkey in monkeys
            yell(monkeys, monkey.first)
            if root.n != nothing
                println("root.n = $(root.n) - stopping")
                break
            end
        end
    end
    display(monkeys)
    println("Root yelled $(root.n)")
end

#=
sjmn: drzm * dbpl
sllz: 4
=#
function decode(line)
    #println("decoding $line")
    r = r"([a-z]+): (.*)"
    m = match(r, line)
    monkey = m[1]
    if tryparse(Int, m[2]) == nothing
        r = r"([a-z]+) ([-+*/]+) ([a-z]+)"
        #println("attempting to match $(m[2])")
        m = match(r, m[2])
        dep1 = m[1]
        dep2 = m[3]
        op = m[2]
        return monkey, 0, [dep1, dep2], op
    else
        return monkey, parse(Int, m[2]), [], ""
    end
end

# Find the monkeys who want to know when "monkey" shouts its number.
function findWhoCares(monkeys, monkey)
    whoCares = []
    #println("looking for who cares about $monkey")
    for m in monkeys
        #println("checking $(m.first) dependencies which are $(m.second.dependentOn)")
        for dep in m.second.dependentOn
            if dep == monkey
                push!(whoCares, m.first)
            end
        end
    end
    return whoCares
end

function yell(monkeys, name)
    m = monkeys[name]
    if length(m.dependentOn) == 0
        #println("$name is yelling $(m.n) since it has no dependencies")
        for d in m.dependentOnMe
            tellMonkey(monkeys, name, d, m.n)
        end
        m.dependentOnMe = []
    #else
    #    println("$name is quiet since it has dependencies: $(m.dependentOn)")
    end
end

function tellMonkey(monkeys, from, to, n)
    println("Telling $to that $from yelled $n")
    tm = monkeys[to]
    for i in 1:length(tm.dependentOn)
        if tm.dependentOn[i] == from
            tm.inputValues[i] = n
            tm.dependentOn[i] = ""
            break
        end
    end
    if length(tm.dependentOn[1]) + length(tm.dependentOn[2]) == 0
        tm.dependentOn = []
        if tm.operation == "+"
            tm.n = tm.inputValues[1] + tm.inputValues[2]
        elseif tm.operation == "-"
            tm.n = tm.inputValues[1] - tm.inputValues[2]
        elseif tm.operation == "*"
            tm.n = tm.inputValues[1] * tm.inputValues[2]
        elseif tm.operation == "/"
            tm.n = tm.inputValues[1] / tm.inputValues[2]
        end
    end
end

main()
