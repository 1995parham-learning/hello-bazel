# Hello Bazel

A collection of simple Bazel examples demonstrating how to build projects in different programming languages using modern Bzlmod.

## Contents

1. **Go Example** (`//hello-go:hello-go`)
   - Web server using Fuego framework
   - Internal package structure
   - Unit tests with go_test

2. **Python Example** (`//hello-py:hello`)
   - Python binary example
   - Demonstrates py_binary rule

## Prerequisites

- [Bazel](https://bazel.build/) 7.x or later (uses Bzlmod)
- Go toolchain (auto-downloaded by Bazel)
- Python 3.x (for Python examples)

## Building

```bash
# Build all targets
bazel build //...

# Build specific targets
bazel build //hello-go:hello-go
bazel build //hello-py:hello
```

## Running

```bash
# Run Go server (starts on :9999)
bazel run //hello-go:hello-go

# Run Python example
bazel run //hello-py:hello
```

## Testing

```bash
# Run all tests
bazel test //...

# Run Go tests only
bazel test //hello-go/...

# Run with verbose output
bazel test //... --test_output=all
```

## Development

### Regenerate BUILD Files

```bash
# Update BUILD files from Go source
bazel run //:gazelle

# Update dependencies from go.mod
bazel run //:gazelle-update-repos
```

### IDE Setup (VS Code / GoLand)

For Go IDE integration with Bazel, set the `GOPACKAGESDRIVER` environment variable:

```bash
export GOPACKAGESDRIVER="$(pwd)/tools/gopackagesdriver.sh"
```

See [rules_go Editor Setup](https://github.com/bazelbuild/rules_go/wiki/Editor-setup) for details.

## Static Analysis

This project uses `nogo` for static analysis during builds:

- Go vet checks (enabled by default)
- Nilness analysis (detect nil pointer dereferences)
- Shadow analysis (detect shadowed variables)
- Unused result analysis (detect ignored return values)

Analysis runs automatically on every build - no separate step needed.

## CI/CD

GitHub Actions workflow runs on push and PR:

- Builds all targets with optimizations
- Runs all tests
- Uses Bazel caching for fast builds

## Project Structure

```
hello-bazel/
├── BUILD.bazel              # Root gazelle targets
├── MODULE.bazel             # Bazel module configuration
├── .bazelrc                 # Bazel settings
├── hello-go/
│   ├── BUILD.bazel          # Go binary and nogo rules
│   ├── main.go              # Server entry point
│   ├── go.mod               # Go module
│   └── internal/handler/    # Internal packages
│       ├── BUILD.bazel
│       ├── hello.go
│       └── hello_test.go
├── hello-py/
│   ├── BUILD.bazel
│   └── main.py
└── tools/
    └── gopackagesdriver.sh  # IDE integration
```
