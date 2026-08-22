# todo-cli

[![CI](https://github.com/danielv27/todo-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/danielv27/todo-cli/actions/workflows/ci.yml)

A simple command-line tool written in Go to manage TODOs in a plain markdown file, using GitHub-style checkbox syntax (`- [ ]` / `- [x]`). Works great as a companion to Obsidian, Logseq, or any other markdown-based notes app — but it doesn't require one.

## Features
- Add, list, complete, un-complete, remove, and clear TODO items from the command line.
- Each TODO gets a short, stable ID so you can reference it later, or match by a unique substring of its text instead.
- Color-coded `list` output: done items in green, pending items in red.

## Setup

### 1. Clone the Repository
```sh
git clone https://github.com/danielv27/todo-cli.git
cd todo-cli
```

### 2. Initialize Go Modules (if not already done)
```sh
go mod tidy
```

### 3. Build the Binary
```sh
go build -o todo
```

### 4. Make the Binary Globally Accessible
Move the binary to a directory in your `PATH`, such as:
```sh
# For your user only (no sudo required):
mv ./todo ~/.local/bin/
# Or for all users (requires sudo):
sudo mv ./todo /usr/local/bin/
```
Make sure the directory is in your `PATH`.

After this, you can run the CLI tool from any terminal session by simply typing:
```sh
todo add "Your task here"
```

## Usage

### Add a TODO
```sh
todo add Buy groceries and pick up the mail
```
This appends a line like:
```
- [ ] [id:1a2b] Buy groceries and pick up the mail
```
to your TODO file.

### List TODOs
```sh
todo list
```

### Mark a TODO as done / not done
```sh
todo done 1a2b
todo done groceries   # matches by unique substring instead of ID
todo undone 1a2b
```

### Remove a TODO
```sh
todo remove 1a2b
todo remove groceries
```

### Clear the entire list
```sh
todo clear
```
You'll be asked to confirm before the file is wiped.

## Configuration

By default, TODOs are read from and written to `~/todo.md`. To use a different file, set the `TODO_CLI_FILE` environment variable:
```sh
export TODO_CLI_FILE="$HOME/Documents/Obsidian Vault/todo.md"
```
Add that line to your shell profile (`.bashrc`, `.zshrc`, etc.) to make it persistent.

## Requirements
- Go 1.24 or newer
- [Cobra CLI](https://github.com/spf13/cobra) (for development)

## Contributing
Pull requests and suggestions are welcome!

## License
MIT
