import util

class xyz:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z
        self.transforms = {}
    
    def __repr__(self):
        return f"{self.x},{self.y},{self.z}"

    def perturb(self, x, y, z):
        return xyz(self.x+x, self.y+y, self.z+z)
    
    def overlaps(self, other):
        return self.x == other.x and self.y == other.y and self.z == other.z
    
    def __hash__(self):
        return hash((self.x, self.y, self.z))

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y and self.z == other.z

    def __lt__(self, other):
        if self.x < other.x: return True
        if self.x > other.x: return False
        if self.y < other.y: return True
        if self.y > other.y: return False
        return self.z < other.z

    def __sub__(self, other):
        return xyz(self.x - other.x, self.y - other.y, self.z - other.z)

    def __mul__(self, other):
        return xyz(self.x * other.x, self.y * other.y, self.z * other.z)

    def __add__(self, other):
        return xyz(self.x + other.x, self.y + other.y, self.z + other.z)

    def rotateAroundX(self, rightAngleCount):
        match rightAngleCount % 4:
            case 0: return self
            case 1: return xyz(self.x, -self.z, self.y)
            case 2: return xyz(self.x, -self.y, -self.z)
            case 3: return xyz(self.x, self.z, -self.y)

    def rotateAroundY(self, rightAngleCount):
        match rightAngleCount % 4:
            case 0: return self
            case 1: return xyz(-self.z, self.y, self.x)
            case 2: return xyz(-self.x, self.y, -self.z)
            case 3: return xyz(self.z, self.y, -self.x)

    def rotateAroundZ(self, rightAngleCount):
        match rightAngleCount % 4:
            case 0: return self
            case 1: return xyz(-self.y, self.x, self.z)
            case 2: return xyz(-self.x, -self.y, self.z)
            case 3: return xyz(self.y, -self.x, self.z)
    
    def transform(self, xr, yr, zr):
        id = (xr, yr, zr)
        if id not in self.transforms:
            t = self.rotateAroundX(xr).rotateAroundY(yr).rotateAroundZ(zr)
            self.transforms[id] = t
        return self.transforms[id]

class Cube:
    def __init__(self, id):
        self.id = id
        self.beacons = set()
        self.distance = 0
        self.offset = xyz(0, 0, 0)

    def __repr__(self):
        return f"id:{self.id}, offset:{self.offset}, beacons:{self.beacons}"
    
    def __iter__(self):
        for beacon in self.beacons:
            yield beacon
        
    def __getitem__(self, item):
        return self.beacons[item]

    def __eq__(self, other):
        return self.beacons == other.beacons

    def __lt__(self, other):
        return self.beacons < other.beacons

    def __hash__(self):
        return hash(len(self.beacons))

    def addBeacon(self, beacon):
        self.beacons.add(beacon)
    
    def transform(self, offset, xr, yr, zr):
        beacons = []
        for beacon in self.beacons:
            beacons.append(beacon.rotateAroundX(xr).rotateAroundY(yr).rotateAroundZ(zr) + offset)
        return beacons

    def addShiftedBeacons(self, other, xr, yr, zr, offset):
        #print(f"adding {other} beacons to {self} rotated by {xr}, {yr}, {zr} shifted by {offset}")
        print(f"adding {len(other.beacons)} beacons rotated by {xr}, {yr}, {zr} shifted by {offset}")
        for beacon in other:
            self.addBeacon(beacon.rotateAroundX(xr).rotateAroundY(yr).rotateAroundZ(zr) + offset)

    def addShiftedBeacons2(self, other):
        print(f"adding {len(other)} beacons")
        for beacon in other:
            self.addBeacon(beacon)

    def overlap(self, other, xr, yr, zr, offset):
        overlap = 0
        for refBeacon in self.beacons:
            for beacon in other.beacons:
                if refBeacon == beacon.rotateAroundX(xr).rotateAroundY(yr).rotateAroundZ(zr) + offset:
                    overlap += 1
        return overlap
    
    def overlap2(self, other):
        overlap = 0
        for refBeacon in self.beacons:
            for beacon in other:
                if refBeacon == beacon:
                    overlap += 1
        return overlap

    def overlap3(self, other, targetCount):
        overlap = 0
        for refBeacon in self.beacons:
            if refBeacon in other:
                overlap += 1
                if overlap >= targetCount:
                    return overlap

            # for beacon in other:
            #     if refBeacon == beacon:
            #         overlap += 1
            #         if overlap >= targetCount:
            #             return overlap
        return overlap
    
    def distance(self, other):
        print(f"calculating distance from {self} to {other}")
        return abs(self.offset.x - other.offset.x) + abs(self.offset.y - other.offset.y) + abs(self.offset.z - other.offset.z)

def addScanner(input, start):
    #print(f"Adding scanner from line {start}")
    numStart = input[start].find("scanner")+8
    numEnd = input[start].find(" ", numStart)
    scannerId = int(input[start][numStart:numEnd])
    scanner = Cube(scannerId)
    for line in range(start+1, len(input)):
        if len(input[line]) == 0:
            break
        coords = input[line].split(',')
        scanner.addBeacon(xyz(int(coords[0]), int(coords[1]), int(coords[2])))
    return scanner, line

def readScans(input):
    scanners = []
    for line in range(0, len(input)):
        if input[line].find('scanner') != -1:
            (scanner, line) = addScanner(input, line)
            scanners.append(scanner)
    return scanners

def mergeScanner(ref, target, overlappingBeaconCount):
    # For each target beacon, move it to match a reference beacon then move the other
    # target beacons by the same amount and see what overlap we get...  If we don't
    # get a big enough overlap, repeat for the next reference beacon.
    #print(f"merging {target}")
    foundOverlap = False
    maxOverlap = 0
    for refBeacon in ref:
        for targetBeacon in target:
            #print(f"checking {targetBeacon} from {target}")
            attemptedTransformations = set()
            for xr in range(0, 4):
                for yr in range(0, 4):
                    for zr in range(0, 4):
                        #transformedBeacon = targetBeacon.rotateAroundX(xr).rotateAroundY(yr).rotateAroundZ(zr)
                        transformedBeacon = targetBeacon.transform(xr, yr, zr)
                        if transformedBeacon not in attemptedTransformations:
                            attemptedTransformations.add(transformedBeacon)
                            offsetTargetToRef = refBeacon - transformedBeacon
                            transformedBeacons = target.transform(offsetTargetToRef, xr, yr, zr)
                            overlapped = ref.overlap3(transformedBeacons, overlappingBeaconCount)
                            if overlapped > maxOverlap:
                                maxOverlap = overlapped
                            if overlapped >= overlappingBeaconCount:
                                ref.addShiftedBeacons2(transformedBeacons)
                                foundOverlap = True
                        if foundOverlap: break
                    if foundOverlap: break
                if foundOverlap: break
            if foundOverlap: break
        if foundOverlap: break
    #if not foundOverlap:
        #print(f"failed to find overlap for {target}, max overlap {maxOverlap}")
    return foundOverlap, offsetTargetToRef

def findOverlap(input, overlappingBeaconCount):
    scanners = readScans(input)
    scannersToMerge = set()
    for n in range(1, len(scanners)):
        scannersToMerge.add(n)
    while len(scannersToMerge) > 0:
        print(f"Merging {len(scannersToMerge)} scanners", flush=True)
        retry = set()
        for scanner in scannersToMerge:
            print(f"Merging scanner #{scanner}", flush=True)
            (foundOverlap, offset) = mergeScanner(scanners[0], scanners[scanner], overlappingBeaconCount)
            if foundOverlap:
                print(f"found overlap, offset: {offset}")
                scanners[scanner].offset = offset
            else:
                retry.add(scanner)
        scannersToMerge = retry
    
    furthest = 0
    for start in range(0, len(scanners)):
        for end in range(0, len(scanners)):
            if start == end:
                continue
            s = scanners[start]
            dist = distance(s.offset, scanners[end].offset)
            if dist > furthest:
                furthest = dist
    print(f"beacon count: {len(scanners[0].beacons)}, max scanner distance: {furthest}")

def distance(a, b):
    print(f"calculating distance from {a} to {b}")
    return abs(a.x - b.x) + abs(a.y - b.y) + abs(a.z - b.z)

if __name__ == "__main__":
    input = util.getInput()
    findOverlap(input, 12)
