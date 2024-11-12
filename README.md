# gogen

[![Go Reference](https://pkg.go.dev/badge/github.com/idelchi/gogen.svg)](https://pkg.go.dev/github.com/idelchi/gogen)
[![Go Report Card](https://goreportcard.com/badge/github.com/idelchi/gogen)](https://goreportcard.com/report/github.com/idelchi/gogen)
[![Build Status](https://github.com/idelchi/gogen/actions/workflows/github-actions.yml/badge.svg)](https://github.com/idelchi/gogen/actions/workflows/github-actions.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`gogen` is a tool for generating cryptographic keys, passwords and password hashes.

## Installation

### From source

```sh
go install github.com/idelchi/gogen@latest
```

### From installation script

```sh
curl -sSL https://raw.githubusercontent.com/idelchi/gogen/refs/heads/main/install.sh | sh -s -- -d ~/.local/bin
```

## Usage

```sh
gogen [flags] command [flags]
```

### Global Flags and Environment Variables

| Flag            | Environment Variable | Description                     | Default |
| --------------- | -------------------- | ------------------------------- | ------- |
| `-s, --show`    | `GOGEN_SHOW`         | Show the configuration and exit | `false` |
| `-h, --help`    | -                    | Help for gogen                  | -       |
| `-v, --version` | -                    | Version for gogen               | -       |

### Commands

#### `key` - Generate a cryptographic key

Generate keys of configurable length.

**Flags and Environment Variables:**

| Flag           | Environment Variable | Description                   | Default | Valid Range             |
| -------------- | -------------------- | ----------------------------- | ------- | ----------------------- |
| `-l, --length` | `GOGEN_LENGTH`       | Length of the key to generate | 32      | 32-512 (multiple of 32) |
| `-h, --help`   | -                    | Help for key command          | -       | -                       |

Examples:

```sh
# Generate a 32-byte key (default)
gogen key

# Generate a 64-byte key
gogen key -l 64

# Key length must be between 32-512 bytes and a multiple of 4
```

#### `password` - Generate a password

Generate secure passwords of configurable length.

**Flags and Environment Variables:**

| Flag           | Environment Variable | Description                        | Default | Valid Range |
| -------------- | -------------------- | ---------------------------------- | ------- | ----------- |
| `-l, --length` | `GOGEN_LENGTH`       | Length of the password to generate | 16      | -           |
| `-h, --help`   | -                    | Help for password command          | -       | -           |

Examples:

```sh
# Generate a 16-character password (default)
gogen password

# Generate a 20-character password
gogen password -l 20

# Using the shorter alias
gogen pw
```

#### `hash` - Hash a password

Hash passwords using bcrypt with configurable cost and benchmarking capabilities.

**Flags and Environment Variables:**

| Flag              | Environment Variable | Description                          | Default | Valid Range |
| ----------------- | -------------------- | ------------------------------------ | ------- | ----------- |
| `-c, --cost`      | `GOGEN_COST`         | Cost of the password hash            | 12      | 4-31        |
| `-b, --benchmark` | `GOGEN_BENCHMARK`    | Run a benchmark on the password hash | `false` | -           |
| `-h, --help`      | -                    | Help for hash command                | -       | -           |

Examples:

```sh
# Hash a password with default cost (12)
gogen hash password

# Hash with custom cost (4-31)
gogen hash -c 14 password

# Run benchmark to measure hashing performance across costs
gogen hash -b password
```

For detailed help on any command:

```sh
gogen --help
gogen <command> --help
```
