# Tmuxcraft

Tmuxcraft is a Go-based alternative to tmuxifier, using YAML configuration files for defining tmux sessions and layouts.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Sample Configuration](#sample-configuration)
- [Releases](#releases)
- [Builds](#builds)
- [Commands](#commands)
- [Contributing](#contributing)
- [Running Tests](#running-tests)
- [Running Linter](#running-linter)
- [Code Formatting](#code-formatting)
- [License](#license)
- [Author](#author)

## Overview

Tmuxcraft allows you to easily manage your tmux sessions and layouts through simple YAML configuration files. It's designed to be lightweight and user-friendly, offering a range of commands to streamline your tmux workflow.

---

## Features

- Define tmux sessions and layouts in YAML.
- Load and list session layouts.
- Create new session layouts with your preferred editor.
- Generate shell scripts from session layouts for easy execution.
- Cross-platform support: Linux, macOS, and Windows builds are available.

---

## Installation

### From Source

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

### From Pre-Built Releases

1. Go to the [Releases](https://github.com/sugan0tech/tmuxcraft/releases) page.
2. Download the appropriate binary for your operating system:
   - `tmuxcraft-linux-amd64` (Linux, 64-bit)
   - `tmuxcraft-darwin-amd64` (macOS, Intel)
   - `tmuxcraft-darwin-arm64` (macOS, M1/M2)
   - `tmuxcraft-windows-amd64.exe` (Windows, 64-bit)

3. Move the binary to a directory in your `PATH` (e.g., `/usr/local/bin` or `~/.local/bin`).

4. Make the binary executable (if necessary):
    ```sh
    chmod +x tmuxcraft
    ```

---

## Usage

To use Tmuxcraft, you can run the following commands:

```sh
tmuxcraft <command> [<args>]
```

---

## Sample Configuration

Here’s an example YAML configuration for a session:

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

Generated shell scripts can be configured using the `-O {path}` flag. Example usage:

```sh
tmuxcraft -ls test -gs -O {path}
```

Generated Script Example:

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

---

## Releases

Tmuxcraft binaries are available for download in the [Releases](https://github.com/sugan0tech/tmuxcraft/releases) section. Each release includes:

- Binaries for Linux, macOS (Intel and ARM), and Windows.
- Change logs for each version.

Example tag: `v1.0.0-alpha.1`

### Release Automation

Releases are automatically created using GitHub Actions whenever a new tag is pushed. The build workflow ensures cross-platform binaries are generated for each release.

---

## Builds

Tmuxcraft is built for the following platforms:

| Platform      | Architecture | File Name                  |
|---------------|--------------|----------------------------|
| Linux         | amd64        | `tmuxcraft-linux-amd64`    |
| macOS (Intel) | amd64        | `tmuxcraft-darwin-amd64`   |
| macOS (M1/M2) | arm64        | `tmuxcraft-darwin-arm64`   |
| Windows       | amd64        | `tmuxcraft-windows-amd64.exe` |

Builds are created using Go’s cross-compilation capabilities.

---

## Commands

List available commands:
```sh
tmuxcraft --help
```

---

## Contributing

We welcome contributions to Tmuxcraft! To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch:
    ```sh
    git checkout -b feature-branch
    ```
3. Make your changes.
4. Commit your changes:
    ```sh
    git commit -m "Add some feature"
    ```
5. Push to the branch:
    ```sh
    git push origin feature-branch
    ```
6. Open a pull request.

Please ensure that your code adheres to the existing coding standards and includes appropriate tests.

---

## Running Tests

To run tests, use the following command:

```sh
go test ./...
```

---

## Running Linter

We use golangci-lint for linting. Ensure you have it installed, then run:

```sh
golangci-lint run ./...
```

---

## Code Formatting

Ensure your code is properly formatted by running:

```sh
gofmt -d .
```

---

## License

Tmuxcraft is released under the MIT License. See [LICENSE](./LICENSE) for more information.

---

## Author

[sugan0tech](https://github.com/sugan0tech)

