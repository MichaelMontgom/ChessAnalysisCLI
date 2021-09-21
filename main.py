from chessdotcom import get_player_profile, get_player_game_archives, get_player_games_by_month_pgn
from utilities import *

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
            player_name = input(f'Please enter in the players profile name: ')
            print(get_player_games_by_month_pgn(player_name, 2021, 9).json)
