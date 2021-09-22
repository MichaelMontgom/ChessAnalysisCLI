from utilities import WrongValueError
from player_profile import Player_Profile
if __name__ == '__main__':
    while True:
        print(f'Welcome to the Chess.com Analysis CLI!')

        while True:
            try:
                choice = int(input(f'1.Get Opening Move Preferences\n'))
                if choice not in [1]:
                    raise WrongValueError
            except WrongValueError:
                print(f'That was not a choice! Try again! \n')
                continue
            else:
                break

        if choice == 1:
            player = Player_Profile(input(f'Enter the username: '))
            player.get_moves()


