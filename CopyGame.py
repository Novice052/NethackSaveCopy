# This was my first crack at the program written in Python.  Keeping it for historical reference.
#
# Requirements:
# Find the save folder
# Find the save file to be copied
#   copy orig to bak folder
#   copy bak to orig folder

import shutil, os, re

# TODO: use REGIX to find playground folder
playground = 'C:\\Users\\immer\\AppData\\Local\\VirtualStore\\Program Files (x86)\\Nethack'
# TODO: use REGIX to find backup folder OR create if not there
backup = 'C:\\Users\\immer\\AppData\\Local\\VirtualStore\\Program Files (x86)\\Nethack\\Backup'
os.chdir(playground)
# TODO: check to see if file exists already, if not, don't copy.
# TODO: use REGIX to find and copy all save files
try:
    shutil.copy(playground + '\\immer-Dirk.NetHack-saved-game', backup)
except IOError:
    print('Couldn\'t find an original save file.  Looking for a backup...' )
# TODO: use REGIX to find all save files
try:
    shutil.copy(backup + '\\immer-Dirk.NetHack-saved-game', playground)
except IOError:
    print('No backup file found.')
