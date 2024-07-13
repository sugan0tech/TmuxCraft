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
