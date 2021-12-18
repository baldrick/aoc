import a

def explode(input):
    print(f"---- exploding {input} ----")
    sfn = a.SnailfishNumber.create(input)
    try:
        sfn.explodeIfDeeplyNested()
    except a.ExplosionComplete:
        pass
    print(f"Explosion complete: {sfn}")
    return f"{sfn}"

def add(in1, in2):
    print("adding {a} and {b}")
    sfnA = a.SnailfishNumber.create(in1)
    sfnB = a.SnailfishNumber.create(in2)
    sfnSum = sfnA.add(sfnB)
    return f"{sfnSum}"

def reduce(input, steps):
    print(f"---- reducing {input} ----")
    sfn = a.SnailfishNumber.create(input)
    for step in range(0, steps):
        print(f"********** STEP {step} **********")
        try:
            sfn.reduce()
        except a.ExplosionComplete:
            pass
    print(f"{steps} steps of reduction complete: {sfn}")
    return f"{sfn}"

def fullReduce(input):
    sfn = a.SnailfishNumber.create(input)
    sfn.fullReduce()
    return f"{sfn}"

def test_explode():
    assert(explode("[[[[[9,8],1],2],3],4]") == "[[[[0,9],2],3],4]"), "single explode 1 failed"
    assert(explode("[7,[6,[5,[4,[3,2]]]]]") == "[7,[6,[5,[7,0]]]]"), "single explode 2 failed"
    assert(explode("[[6,[5,[4,[3,2]]]],1]") == "[[6,[5,[7,0]]],3]"), "single explode 3 failed"
    assert(explode("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]") == "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"), "single explode 4 failed"
    assert(explode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]") == "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"), "single explode 5 failed"

def test_add():
    assert(add("[[[[4,3],4],4],[7,[[8,4],9]]]]", "[1,1]") == "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"), "add failed"

def test_reduce():
    assert(reduce("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 1) == "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]"), "reduce failed after 1 step"
    assert(reduce("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 2) == "[[[[0,7],4],[15,[0,13]]],[1,1]]"), "reduce failed after 2 steps"
    assert(reduce("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 3) == "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"), "reduce failed after 3 steps"
    assert(reduce("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 4) == "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"), "reduce failed after 4 steps"
    assert(reduce("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 5) == "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"), "reduce failed after 5 steps"

def test_fullreduce():
    assert(fullReduce("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]") == "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"), "full reduce failed"

def test_main():
    assert(a.main("test1") == "[[[[1,1],[2,2]],[3,3]],[4,4]]"), "test 1 failed"
    assert(a.main("test2") == "[[[[3,0],[5,3]],[4,4]],[5,5]]"), "test 2 failed"
    assert(a.main("test3") == "[[[[5,0],[7,4]],[5,5]],[6,6]]"), "test 3 failed"
    assert(a.main("test4") == "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"), "test 4 failed"

def test_magnitude():
    assert(a.magnitude("[9,1]") == 29), "magnitude test 1 failed"
    assert(a.magnitude("[[1,2],[[3,4],5]]") == 143), "magnitude test 2 failed"
    assert(a.magnitude("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]") == 1384), "magnitude test 3 failed"
    assert(a.magnitude("[[[[1,1],[2,2]],[3,3]],[4,4]]") == 445), "magnitude test 4 failed"
    assert(a.magnitude("[[[[3,0],[5,3]],[4,4]],[5,5]]") == 791), "magnitude test 5 failed"
    assert(a.magnitude("[[[[5,0],[7,4]],[5,5]],[6,6]]") == 1137), "magnitude test 6 failed"
    assert(a.magnitude("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]") == 3488), "magnitude test 7 failed"

    sfn = a.main("test-magnitude")
    assert(a.SnailfishNumber.create(sfn).magnitude() == 4140), "test-magnitude failed"

def test_largestMagnitudeFromTwo():
    assert(a.largestFromTwo("test-magnitude") == 3993), "test largest magnitude failed"
    