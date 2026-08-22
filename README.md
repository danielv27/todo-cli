# todo-cli

[![CI](https://github.com/danielv27/todo-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/danielv27/todo-cli/actions/workflows/ci.yml)
[![license](https://img.shields.io/badge/license-MIT-5B5BD6?style=flat-square)](https://opensource.org/licenses/MIT)

A fast, no-nonsense TODO manager that lives in a single markdown file. No database, no sync service, no app — just `- [ ]` checkboxes you can grep, diff, and edit by hand whenever you want. Pairs naturally with Obsidian, Logseq, or a plain `todo.md`, but doesn't care which.

## ✨ Features

- ➕ **Add** tasks from the command line, no editor required
- ✅ **Check off / un-check** tasks by ID or by typing a bit of the text
- 🗑️ **Remove** tasks the same way
- 📋 **List** everything with color-coded status (green ✔ done, red ✘ pending)
- 🧹 **Clear** the whole list (with a confirmation prompt, so you don't nuke it by accident)
- 🪶 Zero config to get started, one env var if you want a different file

## How it works

**IDs.** Every task gets a 4-character ID, derived from a SHA-1 hash of the task text plus the current line count in the file (used as a salt so identical text on different lines doesn't collide):

```
- [ ] [id:9f3a] Buy groceries
```

**Matching.** Every command that targets an existing task (`done`, `undone`, `remove`) accepts either the ID or a plain substring of the task text:

```sh
todo done 9f3a          # exact, by ID
todo done groceries     # by substring — works if it's unique
```

If your substring matches more than one task, nothing gets touched — the CLI tells you to be more specific or use the ID instead. No silent wrong-task edits.

## Install

```sh
git clone https://github.com/danielv27/todo-cli.git && cd todo-cli
go build -o todo
mv ./todo ~/.local/bin/   # or /usr/local/bin, or anywhere on your PATH
```

## Usage

```sh
todo add Buy groceries and pick up the mail
# → - [ ] [id:9f3a] Buy groceries and pick up the mail

todo list
todo done 9f3a
todo done groceries      # substring match works too
todo undone 9f3a
todo remove 9f3a
todo clear               # asks for confirmation first
```

## Configuration

Defaults to `~/todo.md`. Point it elsewhere with:

```sh
export TODO_CLI_FILE="$HOME/Documents/Obsidian Vault/todo.md"
```

## Requirements

Go 1.24+

## Contributing

PRs welcome.

## License

MIT
