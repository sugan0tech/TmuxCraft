# Tmuxcraft

Tmuxcraft is a Go-based alternative to tmuxifier, using YAML configuration files for defining tmux sessions and layouts. 

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Author](#author)

## Overview

Tmuxcraft allows you to easily manage your tmux sessions and layouts through simple YAML configuration files. It's designed to be lightweight and user-friendly, offering a range of commands to streamline your tmux workflow.

## Features

- Define tmux sessions and layouts in YAML
- Load and list session layouts
- Create new session layouts with your preferred editor
- Generate shell scripts from session layouts for easy execution

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/sugan0tech/tmuxcraft.git
    ```

2. Navigate to the project directory:
    ```sh
    cd tmuxcraft
    ```

3. Build the project:
    ```sh
    go build -o tmuxcraft ./cmd/tmuxcraft
    ```

4. Ensure the binary is in your `PATH`:
    ```sh
    export PATH=$PATH:$(pwd)
    ```

## Usage

To use Tmuxcraft, you can run the following commands:

```sh
tmuxcraft <command> [<args>]
```

Sample Config file
```yaml
session_name: test
path: ~/Documents/GitHub/tmuxcraft
windows:
  - name: nvim
    path: ~/Documents
    command: nvim
    panes:
      - command: tty-clock -t
        split: h
        size: 20
      - command: bash
        size: 50
        split: v
  - name: resource
    command: htop
    panes: []
  - name: test
    panes: []
```

Executable ready to go, shell scripts. By default will be generated in pwd. Can be configured with `-O {path}` flag

```sh 
tmuxcraft -ls test -gs -O {path}
```

```sh 
#!/bin/bash
tmux new-session -d -s test -c ~/Documents/GitHub/tmuxcraft
tmux rename-window -t test: nvim
sleep 0.1
tmux send-keys -t test:nvim.1 "nvim" C-m
tmux split-window -t test:nvim.1 -c ~/Documents/GitHub/tmuxcraft -h -l 40
sleep 0.1
tmux send-keys -t test:nvim.2 "tty-clock -t" C-m
tmux split-window -t test:nvim.2 -c ~/Documents/GitHub/tmuxcraft -v -l 20
sleep 0.1
tmux send-keys -t test:nvim.3 "bash" C-m
tmux new-window -t test -c ~/Documents/GitHub/tmuxcraft -n resource
sleep 0.1
tmux send-keys -t test:resource.1 "htop" C-m
tmux new-window -t test -c ~/Documents/GitHub/tmuxcraft -n test
```


## Contributing

We welcome contributions to Tmuxcraft! To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Make your changes.
4. Commit your changes (git commit -m 'Add some feature').
5. Push to the branch (git push origin feature-branch).
6. Open a pull request.
7. bash

Please ensure that your code adheres to the existing coding standards and includes appropriate tests.

## Running Tests

To run tests, use the following command:

```sh
go test ./...
```

## Running Linter

We use golangci-lint for linting. Ensure you have it installed, then run:

```sh
golangci-lint run ./...
```

## Code Formatting

Ensure your code is properly formatted by running:

```sh
gofmt -d .
```

## License

Tmuxcraft is released under the MIT License. See LICENSE for more information.

## Author

This project is maintained by sugan0tech.
