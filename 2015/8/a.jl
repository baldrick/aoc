function main(f)
    codeLen = 0
    memLen = 0
    encodedLenDiff = 0
    while ! eof(f)
        line = readline(f)
        encodedLenDiff += encode(line)
        orig = line
        origLen = length(line)
        codeLen += length(line)
        line = replace(line, "\\\\" => ".")
        line = replace(line, "\\\"" => ".")
        line = replace(line, r"\\x.." => ".")
        line = replace(line, r"\"" => "")
        memLen += length(line)
    end
    return encodedLenDiff
end

function encode(line)
    # \" -> \\\" or .eq.
    # \\ -> \\\\ or .bs.
    # \x -> \\x or ...
    # " -> \"" or .q.
    orig = line
    origLen = length(line)
    line = replace(line, "\\\\" => ".bs.")
    line = replace(line, "\\\"" => ".eq.")
    line = replace(line, "\\x" => ".h.")
    line = replace(line, "\"" => ".q.")
    l = length(line)
    d = l - origLen
    println("$orig ($origLen) encodes to $line, length $l, diff $d")
    return d
end

f = open(ARGS[1], "r")
r = main(f)
println("$r")