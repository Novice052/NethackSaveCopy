# NethackSaveCopy

A command-line utility for backing up and restoring NetHack save files.

## Build

```bash
go build -o NethackSaveCopy .
```

## Usage

```bash
NethackSaveCopy [command]
```

## Commands

### copy

Creates a backup from an active save.

```bash
NethackSaveCopy copy <character-name>
```

Example:

```bash
NethackSaveCopy copy Alice
```

This copies:

`<home>\\AppData\\Local\\NetHack\\3.6\\Alice.NetHack-saved-game`

to:

`<home>\\AppData\\Local\\NetHack\\3.6\\Alice.NetHack-saved-game-bak`

### restore

Restores an active save from a backup.

```bash
NethackSaveCopy restore <character-name>
```

Example:

```bash
NethackSaveCopy restore Alice
```

This copies:

`<home>\\AppData\\Local\\NetHack\\3.6\\Alice.NetHack-saved-game-bak`

to:

`<home>\\AppData\\Local\\NetHack\\3.6\\Alice.NetHack-saved-game`

### list

Lists detected active saves and backup saves.

```bash
NethackSaveCopy list
```

## Notes

- This tool targets the default Windows NetHack 3.6 save location.
- Character names are passed as the command argument.



