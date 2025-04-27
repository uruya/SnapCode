# SnapCode

SnapCode is a simple CLI tool that converts code snippets into stylish screenshots automatically.

## Features

- Input a code snippet via command line
- Generate a beautiful screenshot from the code
- Supports dark theme (default)

## Getting Started

### Build

```bash
git clone git@github.com:uruya/SnapCode.git
cd SnapCode
go build -o snapcode ./cmd/cli
./snapcode 'console.log("Hello, SnapCode!");'
```
