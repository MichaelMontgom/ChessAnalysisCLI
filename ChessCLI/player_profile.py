import chessdotcom.types
from chessdotcom import *


class PlayerProfile:
    def __init__(self, username):
        self.username = username
        try:
            self.profile = get_player_profile(username).json
        except ChessDotComError:
            print(f'There was an error retrieving the user by that username')

    def get_moves(self):
        try:
            moves = get_player_games_by_month_pgn(self.username, 2021, 9)
            return moves.__dict__
        except ChessDotComError:
            print(f'There was an error retrieving {self.username} previous games')

    def get_current(self):
        try:
            games = get_player_current_games(self.username)
            return games.__dict__
        except ChessDotComError:
            print(f'There was an issue retrieving {self.username} current games')

    def get_game_archive(self):
        try:
            games = get_player_game_archives(self.username)

            return games.__dict__
        except ChessDotComError:
            print(f'There was an issue retrieving {self.username} archived games')
