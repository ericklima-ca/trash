# Trash 🗑️
`trash` is a command-line interface (CLI) tool designed to prevent irreversible file deletions. It achieves this by moving files to a designated trash folder instead of permanently deleting them, allowing for future cleanup or recovery.

## Warning
**UNDER DEVELOPMENT**

## Installation
You can install `trash` using the following command:

```sh
go install github.com/ericklima-ca/trash
```

## Usage
To use `trash`, you can run the following command:

```sh
$ trash remove file.txt
```

`file.txt` will be moved to the `~/.trash` directory inside of a timestamp folder instead of being deleted.

```sh
user@ in ~/.trash
$ tree
.
└── 20230927005243
    └── file.txt

2 directories, 1 file
```

### Available Commands
- **clean**: Cleans the trash directory.
- **completion**: Generates autocompletion scripts for various shells.
- **help**: Provides help and information about available commands.
- **init**: Initializes the trash directory (default location: `$HOME/.trash`).
- **remove**: Moves files to the trash directory.
- **stats**: Displays statistics about the contents of the trash directory, similar to 'du -sh'.

### Flags
- `-h, --help`: Display help for the `trash` command.
- `-v, --version`: Display the version information for `trash`.

## Creating an Alias for Easy Usage (Optional)
To simplify the usage of `trash` and make it a seamless replacement for the `rm` command, you can create an alias in your shell configuration file (e.g., `.zshrc` for Zsh). This alias will allow you to use `trash` instead of `rm`, ensuring that deleted files go to the trash directory by default.

```.zshrc
# ~/.zshrc
...
alias rm="trash remove"
...
```

## Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:
```sh
echo "autoload -U compinit; compinit" >> ~/.zshrc
```
To load completions in your current shell session:
```sh
source <(trash completion zsh)
```

To load completions for every new session, execute once:

#### Linux:
```sh
trash completion zsh > "${fpath[1]}/_trash"
```

#### macOS:
```sh
trash completion zsh > $(brew --prefix)/share/zsh/site-functions/_trash
```
You will need to start a new shell for this setup to take effect.
