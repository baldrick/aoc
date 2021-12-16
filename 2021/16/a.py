import aoc

input = list(aoc.getInput()[0])

def toDecimal(bin):
    n = 1
    sum = 0
    while len(bin) > 0:
        sum += n * bin.pop()
        n *= 2
    return sum

class Bitstream:
    def __init__(self, input, bits = []):
        self.input = input
        self.binary = bits

    def __len__(self):
        return len(self.binary) + 4 * len(self.input)
       
    def read(self, bitCount):
        while bitCount > len(self.binary) and len(self.input) > 0:
            n = int(self.input.pop(0), 16)
            for bit in range(3, -1, -1):
                self.binary.append((n >> bit) & 1)
        s = slice(bitCount)
        ret = self.binary[s]
        self.binary = self.binary[slice(bitCount, 999)]
        return ret

class Header:
    def __init__(self, bs):
        self.version = toDecimal(bs.read(3))
        self.type = toDecimal(bs.read(3))
    
    def __repr__(self):
        return f"(header: v{self.version} type {self.type})"

class LiteralValue:
    def __init__(self, bs):
        n = []
        more = 1 # make sure we run the loop at least once
        while more == 1:
            more = toDecimal(bs.read(1))
            n += bs.read(4)
        self.value = toDecimal(n)
    
    def __repr__(self):
        return f"(value: {self.value})"

class Operator:
    def __init__(self, bs):
        self.subPackets = []
        self.subPacketCount = 0
        self.subPacketBitLength = 0
        lengthType = toDecimal(bs.read(1))
        if lengthType == 1:
            self.subPacketCount = toDecimal(bs.read(11))
            #print(f"reading {self.subPacketCount} packets", flush=True)
            for _ in range(0, self.subPacketCount):
                self.subPackets.append(Packet(bs))
        else:
            self.subPacketBitLength = toDecimal(bs.read(15))
            #print(f"reading {self.subPacketBitLength} bits-worth of packets")
            packetBits = bs.read(self.subPacketBitLength)
            self.subPackets = Packets(Bitstream("", packetBits)).packets

    def __repr__(self):
        if self.subPacketCount > 0:
            return f"operator: {self.subPacketCount} packets read:\n{self.subPackets}"
        else:
            return f"operator: {self.subPacketBitLength} bits-worth of packets read:\n{self.subPackets}"

    def sumVersions(self):
        sum = 0
        for p in self.subPackets:
            sum += p.sumVersions()
        return sum

class Packets:
    def __init__(self, bs):
        self.packets = []
        while len(bs) > 0:
            self.packets.append(Packet(bs))
    
    def __repr__(self):
        s = ""
        for p in self.packets:
            s += f"{p}\n"
        return s

class Packet:
    def __init__(self, bs):
        self.h = Header(bs)
        match(self.h.type):
            case 4:
                self.literalValue = LiteralValue(bs)
            case _:
                self.operator = Operator(bs)
    
    def __repr__(self):
        s = f"{self.h}"
        match(self.h.type):
            case 4:
                s += f"{self.literalValue}"
            case _:
                s += f"{self.operator}"
        return s

    def sumVersions(self):
        match(self.h.type):
            case 4:
                return self.h.version
            case _:
                return self.h.version + self.operator.sumVersions()
        
    def value(self):
        match(self.h.type):
            case 4:
                return self.literalValue.value
            case 0:
                # return sum of sub-packets
                sum = 0
                for p in self.operator.subPackets:
                    sum += p.value()
                return sum
            case 1:
                # return product of sub-packets
                product = 1
                for p in self.operator.subPackets:
                    product *= p.value()
                return product
            case 2:
                # return min value of sub-packets
                min = 99999999999999999999999999999999999999
                for p in self.operator.subPackets:
                    v = p.value()
                    if v < min:
                        min = v
                return min
            case 3:
                # return max value of sub-packets
                max = 0
                for p in self.operator.subPackets:
                    v = p.value()
                    if v > max:
                        max = v
                return max
            case 5:
                # return 1 if first sub-packet > second
                if len(self.operator.subPackets) != 2:
                    print(f"error: {self.operator} does not contain 2 sub-packets")
                if self.operator.subPackets[0].value() > self.operator.subPackets[1].value():
                    return 1
                return 0
            case 6:
                # return 1 if first sub-packet < second
                if len(self.operator.subPackets) != 2:
                    print(f"error: {self.operator} does not contain 2 sub-packets")
                if self.operator.subPackets[0].value() < self.operator.subPackets[1].value():
                    return 1
                return 0
            case 7:
                # return 1 if first sub-packet == second
                if len(self.operator.subPackets) != 2:
                    print(f"error: {self.operator} does not contain 2 sub-packets")
                if self.operator.subPackets[0].value() == self.operator.subPackets[1].value():
                    return 1
                return 0


print(input)
bs = Bitstream(input)
packet = Packet(bs)
print(packet)
print(packet.sumVersions())
print(packet.value())