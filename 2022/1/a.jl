f = open("e:/dev/aoc/2022/1/puzzle", "r")
cmax = 0
max = 0
while ! eof(f)
  line = readline(f)
  if line == ""
    if cmax > max
        global max = cmax
    end
    global cmax = 0
  else
    val = parse(Int64, line)
    cmax += val
    println("$line = $val")
  end
end
close(f)
println("max = $max")