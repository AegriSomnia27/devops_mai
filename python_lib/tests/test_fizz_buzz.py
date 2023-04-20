from fuzzbuzz import fuzzbuzz


def test_fuzz():
    value = 3
    for _ in range(10):
        assert fuzzbuzz(value) == 'Fuzz'
        value *= 3


def test_buzz():
    value = 5
    for _ in range(10):
        assert fuzzbuzz(value) == 'Buzz'
        value *= 5


def test_fuzz_buzz():
    value = 15
    for _ in range(10):
        assert fuzzbuzz(value) == 'FuzzBuzz'
        value *= 15

