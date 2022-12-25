function main()
    f = open(ARGS[1], "r")
    total = 0
    while !eof(f)
        line = readline(f)
        dec = decode(line)
        println("$line --> $dec")
        total += dec
    end
    println("total = $total = $(encode(total))")
end

function decode(line)
    # 2, 1, 0, -, =
    dec = 0
    multiplier = 1
    for i in length(line):-1:1
        if line[i] == '='
            dec -= multiplier * 2
        elseif line[i] == '-'
            dec -= multiplier
        elseif line[i] == '1'
            dec += multiplier
        elseif line[i] == '2'
            dec += 2 * multiplier
        end
        multiplier *= 5
    end
    return dec
end

snafu_digits = Dict(-2=>"=", -1=>"-", 0=>"0", 1=>"1", 2=>"2")

# I resorted to reddit for this as I was coming up with overly complex solutions.
# I still can't remember enough school maths to understand why ceil(log(2n)/log(5))
# gives the number of digits required for the base-5 number...
function encode(dec)
    digits = []
    num_digits = ceil(log(2*dec) / log(5))
    for i in num_digits-1:-1:0
        digit = Int(round(dec / 5^i))
        dec -= digit * 5^i
        push!(digits, snafu_digits[digit])
    end
    return join(digits)
end

if length(ARGS) > 1 && ARGS[2] == "conv"
    if length(ARGS) > 2
        println(encode(parse(Int, ARGS[3])))
    else
        f = open(ARGS[1], "r")
        while !eof(f)
            line = readline(f)
            if line[1] == '#'
                continue
            end
            dec, snafu = split(line)
            got = decode(snafu)
            if got != parse(Int, dec)
                println("ERROR: want '$dec' from '$snafu', got '$got'")
            end
            println("Encoding $dec")
            got = encode(parse(Int, dec))
            if got != snafu
                println("ERROR: want '$snafu' from '$dec' got '$got'")
            end
        end
    end
else
    main()
end
