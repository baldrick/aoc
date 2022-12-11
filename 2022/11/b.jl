function main()
    f = open(ARGS[1], "r")
    monkeys = []
    while !eof(f)
        push!(monkeys, read_monkey(f))
        if !eof(f)
            blank = readline(f)
        end
    end
    superModulo = calculateSuperModulo(monkeys)
    display(monkeys)
    # Reduce to 20 rounds for part A
    for round in 1:10000
        for monkey in monkeys
            turn(monkey, monkeys, superModulo)
        end
    end
    reportInspections(monkeys)
end

struct Operation
    op::String
    n::Int
    self::Bool
end

mutable struct Monkey
    n::Int
    items::Array{Int}
    operation::Operation
    divisibleBy::Int
    iftrue::Int
    iffalse::Int
    inspections::Int
end

#=
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3
=#
function read_monkey(f)
    return Monkey(
        decodeMonkey(readline(f)),
        decodeStartingItems(readline(f)),
        decodeOperation(readline(f)),
        decodeTest(readline(f)),
        decodeIfTrue(readline(f)),
        decodeIfFalse(readline(f)),
        0
    )
end

function decodeMonkey(s)
    r = r"Monkey ([0-9]+):"
    m = match(r, s)
    return parse(Int, m[1])
end

function decodeStartingItems(s)
    r = r" *Starting items: (.*)"
    m = match(r, s)
    items = []
    for item in split(m[1], ", ")
        push!(items, parse(Int, item))
    end
    return items
end

function decodeOperation(s)
    r = r" *Operation: new = old (.*) (.*)"
    m = match(r, s)
    self = m[2] == "old"
    return Operation(m[1], self ? 0 : parse(Int, m[2]), self)
end

function decodeTest(s)
    r = r" *Test: divisible by ([0-9]+)"
    m = match(r, s)
    return parse(Int, m[1])
end

function decodeIfTrue(s)
    r = r" *If true: throw to monkey ([0-9]+)"
    m = match(r, s)
    return parse(Int, m[1])
end

function decodeIfFalse(s)
    r = r" *If false: throw to monkey ([0-9]+)"
    m = match(r, s)
    return parse(Int, m[1])
end

function calculateSuperModulo(monkeys)
    sm = 1
    for monkey in monkeys
        sm *= monkey.divisibleBy
    end
    return sm
end

#=
Monkey inspects an item with a worry level of 79.
Worry level is multiplied by 19 to 1501.
Monkey gets bored with item. Worry level is divided by 3 to 500. (except in part B)
Current worry level is not divisible by 23.
Item with worry level 500 is thrown to monkey 3.
=#
function turn(monkey, monkeys, superModulo)
    for item in monkey.items
        monkey.inspections += 1
        worry = inspect(item, monkey.operation)
        # Uncomment for part A
        #worry = floor(worry/3)
        worry = worry % superModulo
        if worry % monkey.divisibleBy == 0
            push!(monkeys[monkey.iftrue+1].items, worry)
        else
            push!(monkeys[monkey.iffalse+1].items, worry)
        end
    end
    monkey.items = []
end

function inspect(item, operation)
    n = operation.self ? item : operation.n
    if operation.op == "+"
        return item + n
    elseif operation.op == "*"
        return item * n
    end
    println("Unhandled operation $(operation.op)")
    exit(-1)
end

function reportInspections(monkeys)
    inspections = []
    for monkey in monkeys
        println("Monkey $(monkey.n) inspected items $(monkey.inspections) times.")
        push!(inspections, monkey.inspections)
    end
    inspections = sort(inspections, rev=true)
    println("$(inspections[1]) * $(inspections[2]) = $(inspections[1] * inspections[2])")
end

main()
