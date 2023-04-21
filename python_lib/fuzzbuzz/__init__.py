from typing import Optional


def fuzzbuzz(number: int) -> Optional[str]:
    if number % 3 == 0 and number % 5 == 0:
        return 'FuzzBuzz'
    elif number % 3 == 0:
        return 'Fuzz'
    elif number % 5 == 0:
        return 'Buzz'

    return None


__all__ = ('fuzzbuzz',)
