from chessdotcom import *

class Player_Profile:
    def __init__(self, username):
        self.username = username
        try:
            self.profile = get_player_profile(username).json
        except ChessDotComError:
            print(f'There was an error retrieving the user by that username')


    def get_moves(self):
        try:
            moves = get_player_games_by_month_pgn(self.username, 2021, 9)
        except ChessDotComError:
            print(f'There was an error retrieving ')