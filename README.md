# Go Utils

## Overview

`go-utils` is a collection of DiPhyx utility functions.

## Installation

To install `go-utils`, use the following command:

```sh
go get github.com/diphyx/go-utils
```

## Usage

Import the package in your Go code:

```go
import "github.com/diphyx/go-utils"
```

### Encryption

The `Encryption` struct provides methods for encoding, decoding, encrypting, and decrypting strings.

#### Example

```go
package main

import (
    "fmt"
    "github.com/diphyx/go-utils"
)

func main() {
    encryption, err := utils.NewEncryption("ABCDEFGHIJKLMNOPQRSTUVWX12345678", "1234567890123456")
    if err != nil {
        fmt.Println("Error creating encryption:", err)

        return
    }

    encoded, err := encryption.Encode("Sample")
    if err != nil {
        fmt.Println("Error encoding:", err)

        return
    }

    decoded, err := encryption.Decode(encoded)
    if err != nil {
        fmt.Println("Error decoding:", err)

        return
    }

    fmt.Println("Encoded:", encoded)
    fmt.Println("Decoded:", decoded)

    encrypted, err := encryption.Encrypt("Sample")
    if err != nil {
        fmt.Println("Error encrypting:", err)

        return
    }

    decrypted, err := encryption.Decrypt(encrypted)
    if err != nil {
        fmt.Println("Error decrypting:", err)

        return
    }

    fmt.Println("Encrypted:", encrypted)
    fmt.Println("Decrypted:", decrypted)
}
```

### IP Conversion

The `IpToNumber` and `NumberToIp` functions convert IPv4 addresses to their integer representations and vice versa.

#### Example

```go
package main

import (
    "fmt"
    "github.com/diphyx/go-utils"
)

func main() {
    ip := "192.168.0.1"
    number, err := utils.IpToNumber(ip)
    if err != nil {
        fmt.Println("Error converting IP to number:", err)

        return
    }

    fmt.Println("IP to Number:", number)

    ip, err = utils.NumberToIp(number)
    if err != nil {
        fmt.Println("Error converting number to IP:", err)

        return
    }

    fmt.Println("Number to IP:", ip)
}
```

### Secret Generation

The `NewSecret` function generates a new secret string with the given prefix.

#### Example

```go
package main

import (
    "fmt"
    "github.com/diphyx/go-utils"
)

func main() {
    secret, err := utils.NewSecret("prefix_")
    if err != nil {
        fmt.Println("Error generating secret:", err)

        return
    }

    fmt.Println("Generated Secret:", secret)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
