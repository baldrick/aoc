using MD5

key=ARGS[1]
n=0

while true
    h=bytes2hex(md5("$key$n"))
    if startswith(h,"000000")
        println("n=$n")
        exit()
    end
    global n += 1
end
