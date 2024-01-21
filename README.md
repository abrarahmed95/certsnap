# CertSnap

CertSnap is a command-line interface (CLI) tool designed to check the expiration status of SSL certificates for specified domains.

## Usage

```bash
certsnap [command]
```

### Available Commands

- `check`: Check the expiration status of specified domains
- `help`: Get help about any command

### Flags

- `-h, --help`: Display help for CertSnap

Use `certsnap [command] --help` for more detailed information about a specific command.

### Build Instructions

To build, use the following commands:

```bash
make build
```

### Cleaning Up

To clean up build artifacts, run:

```bash
make build
```

This will remove the build directory and any generated files.
