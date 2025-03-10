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
    encryption, encryptionError := utils.NewEncryption("ABCDEFGHIJKLMNOPQRSTUVWX12345678", "1234567890123456")
    if encryptionError != nil {
        fmt.Println("Error creating encryption:", encryptionError)

        return
    }

    encoded, encodeError := encryption.Encode("Sample")
    if encodeError != nil {
        fmt.Println("Error encoding:", encodeError)

        return
    }

    decoded, decodeError := encryption.Decode(encoded)
    if decodeError != nil {
        fmt.Println("Error decoding:", decodeError)

        return
    }

    fmt.Println("Encoded:", encoded)
    fmt.Println("Decoded:", decoded)

    encrypted, encryptError := encryption.Encrypt("Sample")
    if encryptError != nil {
        fmt.Println("Error encrypting:", encryptError)

        return
    }

    decrypted, decryptError := encryption.Decrypt(encrypted)
    if decryptError != nil {
        fmt.Println("Error decrypting:", decryptError)

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
    number, convertToNumberError := utils.IpToNumber(ip)
    if convertToNumberError != nil {
        fmt.Println("Error converting IP to number:", convertToNumberError)

        return
    }

    fmt.Println("IP to Number:", number)

    ip, convertToIpError = utils.NumberToIp(number)
    if convertToIpError != nil {
        fmt.Println("Error converting number to IP:", convertToIpError)

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
    secret, secretError := utils.NewSecret("prefix_")
    if secretError != nil {
        fmt.Println("Error generating secret:", secretError)

        return
    }

    fmt.Println("Generated Secret:", secret)
}
```

### YAML Template Rendering

The `RenderYamlTemplate` function renders a YAML template with the given variables.

#### Example

```go
package main

import (
    "fmt"
    "github.com/diphyx/go-utils"
)

func main() {
    template := "PLACEHOLDER:\n    default: placeholder\n    required: true\n---\nname: {{ PLACEHOLDER }}"

    _, parseError := utils.ParseYamlTemplate(template)
    if parseError != nil {
        fmt.Println("Error parsing YAML template:", parseError)

        return
    }

    variables := map[string]string{
        "PLACEHOLDER": "placeholder",
    }

    rendered, renderError := utils.RenderYamlTemplate(template, variables)
    if renderError != nil {
        fmt.Println("Error rendering YAML template:", renderError)

        return
    }

    fmt.Println("Rendered YAML Template:", rendered)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
