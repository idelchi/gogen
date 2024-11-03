# gogen

A tool for generating cryptographic keys and password hashes.

## Installation

### From source

```sh
go install github.com/idelchi/gogen/cmd/gogen@latest
```

### From installation script

````sh
curl -sSL https://raw.githubusercontent.com/idelchi/gogen/refs/heads/main/install.sh | sh -s -- -d ~/.local/bin
```

## Usage

```sh
gogen [flags] command [flags]
````

### Global Flags

```sh
-s, --show      Show the configuration and exit
-h, --help      Help for gogen
-v, --version   Version for gogen
```

### Commands

#### `key` - Generate a cryptographic key

Generate keys of configurable length:

```sh
gogen key [flags]

Flags:
  -h, --help         Help for key command
  -l, --length int   Length of the key to generate (default 32)
```

Examples:

```sh
# Generate a 32-byte key (default)
gogen key

# Generate a 64-byte key
gogen key -l 64

# Key length must be between 32-512 bytes and a multiple of 4
```

#### `hash` - Hash a password

Hash passwords using bcrypt with configurable cost and benchmarking capabilities:

```sh
gogen hash [flags] password

Flags:
  -b, --benchmark   Run a benchmark on the password hash
  -c, --cost int    Cost of the password hash (4-31) (default 12)
  -h, --help        Help for hash command
```

Examples:

```sh
# Hash a password with default cost (12)
gogen hash password

# Hash with custom cost (4-31)
gogen hash -c 14 password

# Run benchmark to measure hashing performance across costs
gogen hash -b password
```

### Environment Variables

All flags can be configured through environment variables using the `GOGEN` prefix:

- `GOGEN_SHOW`: Show configuration
- `GOGEN_LENGTH`: Key generation length
- `GOGEN_COST`: Bcrypt cost factor
- `GOGEN_BENCHMARK`: Enable benchmark mode

For detailed help on any command:

```sh
gogen --help
gogen <command> --help
```
