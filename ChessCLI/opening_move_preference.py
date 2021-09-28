from player_profile import PlayerProfile
import re

def opening_move_preference():
    profile_name = input(f'Please enter the name of the profile: ')
    player = PlayerProfile(profile_name)

    moves = player.get_moves_for_year()

    move_dict_white = []
    move_dict_black = []

    for month in moves:

        move_dict_white.append(re.findall('1\. ... ', month['text']))
        move_dict_black.append((re.findall('1\.\.\. ... ', month['text'])))


    print(move_dict_white)
    print(move_dict_black)




