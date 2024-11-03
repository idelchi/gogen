# gogen

A Go tool for cryptographic key generation and password hashing.

## Installation

### From source

```sh
go install github.com/idelchi/gogen/cmd/gogen@latest
```

## Usage

```sh
gogen [flags] command [flags]
```

### Commands

- `key`: Generate cryptographic keys
- `hash`: Hash passwords using bcrypt

### Available flags

```sh
--show, -s:    Show the configuration and exit
--version:     Show version information
```

### Key generation

Generate cryptographic keys with configurable length:

```sh
# Generate a 32-byte key (default)
gogen key

# Generate a 64-byte key
gogen key -l 64

# Supported lengths: 32-512 bytes (must be multiple of 4)
```

### Password hashing

Hash passwords using bcrypt with configurable cost:

```sh
# Hash a password with default cost (12)
gogen hash mypassword

# Hash with custom cost (4-31)
gogen hash -c 14 mypassword

# Run password hashing benchmark
gogen hash -b mypassword
```

### Environment variables

All flags can be set through environment variables with the `GOGEN` prefix:

- `GOGEN_SHOW`: Show configuration
- `GOGEN_LENGTH`: Key length
- `GOGEN_COST`: Bcrypt cost
- `GOGEN_BENCHMARK`: Run benchmark

For more details on usage and configuration:

```sh
gogen --help
gogen key --help
gogen hash --help
```
