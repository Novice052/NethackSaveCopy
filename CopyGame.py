# Requirements:
# Find the save folder
# Find the save file to be copied
#   copy orig to bak folder
#   copy bak to orig folder

import shutil, os

# TODO: use REGIX to find playground folder
playground = 'C:\\Users\\immer\\AppData\\Local\\VirtualStore\\Program Files (x86)\\Nethack'
# TODO: use REGIX to find backup folder OR create if not there
backup = 'C:\\Users\\immer\\AppData\\Local\\VirtualStore\\Program Files (x86)\\Nethack\\Backup'
os.chdir(playground)
# TODO: cehck to see if file exists already, if not, don't copy.
# TODO: use REGIX to find and copy all save files
shutil.copy(playground + '\\immer-Dirk.NetHack-saved-game', backup)
# TODO: use REGIX to find all save files
shutil.copy(backup + '\\immer-Dirk.NetHack-saved-game', playground)
