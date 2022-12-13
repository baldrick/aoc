# struct Opt{T}
#     x::Union{T, Nothing}
#     function Opt{T}(x = nothing) where {T}
#         T <: Packet || throw(ArgumentError("invalid type"))
#         new{T}(x)
#     end  
# end

# mutable struct Packet
#     x::Union{Int, Packet, Nothing}
#     # n::Int # n = -1 => embedded Packet
#     # p::Packet
# end

struct PacketPair
    left::Array
    right::Array
end

function main()
    f = open(ARGS[1], "r")
    pairs = []
    while !eof(f)
        push!(pairs, read_pair(f))
        blank = readline(f)
        #println("Skipped '$blank'")
    end
    #display(pairs)
    total = 0
    for i in 1:length(pairs)
        println("== Pair $i ==")
        if less_than(pairs[i].left, pairs[i].right, 1, false)
            println("Index $i - left < right")
            total += i
        end
    end
    println(total)
end

function read_pair(f)
    left = read_list(f)
    right = read_list(f)
    return PacketPair(left, right)
end

function read_list(f)
    line = getline(f)
    if length(line) > 0
        (packets, inc) = decodePacket(line, 1)
        return packets
    end
end

function decodePacket(line, start)
    println("Decoding $line from $start: $(line[start:end])")
    origStart = start
    packets = []
    while start <= length(line)
        if line[start] == '['
            (nestedPackets, inc) = decodePacket(line, start+1)
            println("Got $nestedPackets, inc start by $inc to $(start+inc) for $(line[(start+inc):end])")
            start += inc
            while start <= length(line) && line[start] == ','
                println("Skipping comma, inc start from $start to $(start+1)")
                start += 1
            end
            push!(packets, nestedPackets)
        elseif line[start] == ']'
            start += 1
            while start <= length(line) && line[start] == ','
                println("Skipping comma, inc start from $start to $(start+1)")
                start += 1
            end
            println("Found end of array, skipping by $(start-origStart) = $(line[start:end])")
            return packets, start - origStart
        elseif line[start] == ','
            start += 1
        else
            r = r"([0-9]+).*"
            m = match(r, line[start:end])
            println("Got $m from $(line[start:end])")
            start += length(m[1])
            push!(packets, parse(Int, m[1]))
        end
    end
    return packets, start - origStart
end

function getline(f)
    line = "#"
    while !eof(f) && length(line) > 0 && line[1] == '#'
        line = readline(f)
    end
    if eof(f)
        return ""
    end
    return line
end

function less_than(leftArray, rightArray, indent, treatAsSingleNumber)
    sindent = repeat('.', indent)
    println("$sindent- Compare $leftArray vs $rightArray (tasn=$treatAsSingleNumber)")
    if length(leftArray) == 0 && length(rightArray) > 0
        return true
    elseif length(rightArray) == 0 && length(leftArray) > 0
        return false
    end
    il = 1
    ir = 1
    while il <= length(leftArray) && ir <= length(rightArray)
        left = leftArray[il]
        right = rightArray[ir]
        if left isa Number && right isa Number
            println("$sindent- Compare $left vs $right")
            if left > right
                println("$sindent- $left > $right - FAIL")
                return false
            end
            if treatAsSingleNumber
                return true
            end
        elseif left isa Number
            println("$sindent- Mixed types; convert left to [$left] and retry")
            println("$sindent- Compare [$left] vs $(right)")
            if !less_than([left], right, indent+1, true)
                println("$sindent- $left > $right - FAIL")
                return false
            end
            if treatAsSingleNumber
                return true
            end
        elseif right isa Number
            println("$sindent- Mixed types; convert right to [$right] and retry")
            println("$sindent- Compare $(left) vs [$right]")
            if !less_than(left, [right], indent+1, true)
                println("$sindent- $left > $right - FAIL")
                return false
            end
            if treatAsSingleNumber
                return true
            end
        else
            if !less_than(left, right, indent+1, treatAsSingleNumber)
                return false
            end
            if treatAsSingleNumber
                return true
            end
        end
        il += 1
        ir += 1
        if ir > length(rightArray) && il <= length(leftArray)
            println("$sindent- Right ($right) ran out of items before left ($left) at $ir")
            return false
        elseif il > length(leftArray) && ir <= length(rightArray)
            println("$sindent- Left ($left) ran out of items before right ($right) at $il")
        end
    end
    return true
end

main()
