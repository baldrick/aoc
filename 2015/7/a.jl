function main()
    f = open(ARGS[1], "r")
    values = Dict()
    futures = Dict()
    while ! eof(f)
        line = readline(f)
        (lvalue, rvalue, output, completed) = exec(values, line)
        if completed
            executeFutures(values, futures, output)
        else
            futureOp(values, futures, lvalue, line)
            futureOp(values, futures, rvalue, line)
        end
    end
    println(futures)
    return values
end

function interpret(s)
    r = r"(.*) -> (.*)"
    m = match(r, s)
    return (m[1], m[2])
end

# s can be a number of different things:
#   a number, e.g. 123, return 123, nothing, nothing
#   NOT x, return nothing, NOT, x
#   a AND|OR|LSHIFT|RSHIFT b, return a, $gate, b
function decode(s)
    if startswith(s, "NOT")
        return (nothing, "NOT", s[5:end])
    end
    r = r"(.*) (AND|OR|LSHIFT|RSHIFT) (.*)"
    m = match(r, s)
    if m != nothing
        return (m[1], m[2], m[3])
    end
    return s, nothing, nothing
end

function canOperate(values, operand)
    if isNumber(operand)
        return true
    end
    return haskey(values, operand)
end

function getValue(values, operand)
    if isNumber(operand)
        return parse(UInt16, operand)
    end
    return values[operand]
end

function exec(values, line)
    println("EXECUTING $line")
    (input, output) = interpret(line)
    (lvalue, op, rvalue) = decode(input)
    if op == nothing
        n = r"([0-9]+)"
        m = match(n, lvalue)
        if m == nothing
            if canOperate(values, lvalue)
                values[output] = getValue(values, lvalue)
                return nothing, nothing, output, true
            end
        else
            values[output] = parse(UInt16, m[1])
            return nothing, nothing, output, true
        end
        return lvalue, nothing, output, false
    elseif op == "NOT"
        if canOperate(values, rvalue)
            values[output] = ~getValue(values, rvalue)
            return nothing, nothing, output, true
        end
        return rvalue, nothing, output, false
    elseif op == "RSHIFT"
        if canOperate(values, lvalue)
            values[output] = getValue(values, lvalue) >> parse(UInt16, rvalue)
            return nothing, nothing, output, true
        end
        return lvalue, nothing, output, false
    elseif op == "LSHIFT"
        if canOperate(values, lvalue)
            values[output] = getValue(values, lvalue) << parse(UInt16, rvalue)
            return nothing, nothing, output, true
        end
        return lvalue, nothing, output, false
    elseif op == "AND"
        if canOperate(values, lvalue) && canOperate(values, rvalue)
            values[output] = getValue(values, lvalue) & getValue(values, rvalue)
            return nothing, nothing, output, true
        end
        return lvalue, rvalue, output, false
    elseif op == "OR"
        if canOperate(values, lvalue) && canOperate(values, rvalue)
            values[output] = getValue(values, lvalue) | getValue(values, rvalue)
            return nothing, nothing, output, true
        end
        return lvalue, rvalue, output, false
    else
        println("ERROR processing $line: $lvalue $op $rvalue -> $output")
        exit(1)
    end
end

function futureOp(values, futures, awaitingValue, line)
    if awaitingValue == nothing || haskey(values, awaitingValue) || isNumber(awaitingValue)
        return
    end
    if haskey(futures, awaitingValue)
        push!(futures[awaitingValue], line)
    else
        futures[awaitingValue] = [line]
    end
    println("Added future for $awaitingValue: $line")
end

function isNumber(v)
    r=r"[0-9]+"
    return match(r, v) != nothing
end

function executeFutures(values, futures, gotValue)
    if ! haskey(futures, gotValue)
        return
    end
    f=futures[gotValue]
    println("executing futures for $gotValue: $f")
    println("all futures: $futures")
    for item in 1:length(futures[gotValue])
        line = futures[gotValue][item]
        if line == ""
            continue
        end
        (l, r, output, completed) = exec(values, line)
        println("future exec for $line returned $l, $r, $output, $completed")
        futures[gotValue][item] = ""
        if completed
            executeFutures(values, futures, output)
        end
    end
    f=futures[gotValue]
    println("deleting completed commands from $f")
    deleteat!(futures[gotValue], findall(line -> line == "", futures[gotValue]))
    f=futures[gotValue]
    println("deleting completed commands gave $f")

    println("deleting empty futures from $futures")
    filter!(f -> length(last(f)) != 0, futures)
    println("deleting empty futures gave $futures")
end

values = main()
for v in values
    (a,b) = v
    println("$a = $b")
end


