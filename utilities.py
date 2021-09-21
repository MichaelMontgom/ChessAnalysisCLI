import os


class Error(Exception):
    pass


class WrongValueError(Error):
    pass


# clear_console = lambda: os.system('cls' if os.name in ('nt', 'dos') else 'clear')

def clear_console():
    """Clear console depending on operating system"""

    if os.name in ('nt', 'dos'):
        os.system('cls')
    else:
        os.system('clear')