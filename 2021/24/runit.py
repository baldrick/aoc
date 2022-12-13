import generated

count = 0
for n in range(99999999999999, 99999999999997, -1): #11111111111111, -1):
    s = str(n)
    if s.find("0") == -1: continue
    z = generated.run(s)
    if z == 0:
        print(f"{n} produces zero result")
        break
    if count % 100000 == 0:
        print(f"done {count}, now at {n}")
    count += 1
