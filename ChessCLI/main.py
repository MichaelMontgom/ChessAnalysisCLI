from utilities import WrongValueError
from player_profile import PlayerProfile
from datetime import date
if __name__ == '__main__':
    print()
    player = PlayerProfile('DarthBraves')
    print(player.get_moves_for_year())



