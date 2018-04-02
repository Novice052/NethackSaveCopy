# Requirements:
# Find the save folder
# Find the save file to be copied
#   copy orig to bak folder
#   copy bak to orig folder

import shutil, os

playground = 'C:\\Users\\immer\\AppData\\Local\\VirtualStore\\Program Files (x86)\\Nethack'
backup = 'C:\\Users\\immer\\AppData\\Local\\VirtualStore\\Program Files (x86)\\Nethack\\Backup'
os.chdir(playground)
shutil.copy(playground + '\\immer-Dirk.NetHack-saved-game', backup)
shutil.copy(backup + '\\immer-Dirk.NetHack-saved-game', playground)
