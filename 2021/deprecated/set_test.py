import set

def test_add():
    s = set.Set()
    assert(s.add(1)), "Add 1 for the first time should return True"
    assert(not s.add(1)), "Add 1 for the second time should return False"
    assert(s.add("foo")), "Should be able to add strings to Set, should return True"
    assert(not s.add("foo")), "Should only be able to add a given string once to Set, expected False"

def test_len():
    s = set.Set()
    s.add(1)
    s.add("two")
    assert(len(s)) == 2, f"Set should contain 2 items at this stage, not {len(s)}"

def test_in():
    s = set.Set()
    s.add(1)
    s.add("foo")
    assert(1 in s), "Should find 1 in s: {s}"
    assert(2 not in s), "Should not find 2 in s: {s}"
    assert("foo" in s), "Should find foo in s: {s}"
    assert("bar" not in s), "Should not find foo in s: {s}"
