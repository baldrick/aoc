f = open("e:/dev/aoc/2022/1/puzzle", "r")
elves = []
payload = 0
while ! eof(f)
  line = readline(f)
  if line == ""
    push!(elves, payload)
    global payload = 0
  else
    val = parse(Int64, line)
    global payload += val
  end
end
if payload > 0
  push!(elves, payload)
end
close(f)
sort!(elves, rev=true)
println("$elves")
max = elves[1] + elves[2] + elves[3]
println("max = $max")