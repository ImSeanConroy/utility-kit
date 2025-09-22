# Utility Kit

A collection of my most commonly used command-line utilities — things I reach for weekly like generating secure passwords or picking meal ideas — and it will continue to grow as I add more tools over time.

## Table of Contents

- [Getting Started](#getting-started)
- [Usage](#usage)
- [License](#license)
- [Support](#support)

## Getting Started

### Prerequisites

Before getting started, ensure you have the following installed:

* [Go](https://go.dev/dl/) (1.18+ recommended)

### Installation

The following contains instruction for getting this application running locally:

1. Clone the repository:
```bash
git clone https://github.com/imseanconroy/utility-kit.git
cd utility-kit
```

2. Build the binary:
```bash
go build -o kit
```

3. (Optional) Install globally:
```bash
mv kit /usr/local/bin/
```

Now you can run `kit` from anywhere in your terminal.

## Usage

Generate a password:

```bash
kit password
```

Generate a meal suggestion:

```bash
kit meals
```

## License

This project is distributed under the MIT License – see the [LICENSE](LICENSE) file for information.

## Support

If you are having problems, please let me know by [raising a new issue](https://github.com/imseanconroy/utility-kit/issues/new/choose).
