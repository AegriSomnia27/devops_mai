def fuzzbuzz(number: int) -> str | None:
    if number % 3 == 0 and number % 5 == 0:
        return 'FuzzBuzz'
    elif number % 3 == 0:
        return 'Fuzz'
    elif number % 5 == 0:
        return 'Buzz'

    return None


__all__ = ('fuzzbuzz',)
