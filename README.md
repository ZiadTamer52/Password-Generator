# Password Generator in Go

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)  
A simple and secure password generator written in Go (Golang). This tool allows you to generate random passwords of customizable length and complexity, including uppercase letters, lowercase letters, numbers, and special characters.

---

## Features

- Generate passwords with customizable length.
- Include or exclude:
  - Uppercase letters (`A-Z`)
  - Lowercase letters (`a-z`)
  - Numbers (`0-9`)
  - Special characters (`!@#$%^&*`, etc.)
- Secure random number generation using Go's `crypto/rand` package.
- Easy-to-use command-line interface.

---

## Installation

1. **Ensure you have Go installed** (version 1.20 or higher).  
   Download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

2. **Clone this repository**:
   ```bash
   git clone https://github.com/yourusername/password-generator-go.git
   cd password-generator-go

---
1-Build the project:
   ```bash
   go build
 ```
2-Run the executable:
   ```bash
   ./password-generator-go
 ```
---
## Usage
Run the program with command-line flags to customize the password generation:
   ```bash
   ./password-generator-go -length=16 -uppercase=true -lowercase=true -numbers=true -special=true
 ```
---
## Example
Generate a 16-character password with all character types enabled:
   ```bash
./password-generator-go -length=16 -uppercase=true -lowercase=true -numbers=true -special=true
```
## Output:
   ```bash
Generated Password: aB3$dE7!gH2@jK9#
```
