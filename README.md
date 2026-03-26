# handleErr

A lightweight Go utility for structured error handling. It simplifies the process of separating internal debug information from user-facing error messages.

---

## Features

- **Simplified Checks:** Use `CheckErr` to handle logic flow and messaging in one go.
- **Environment Aware:** Toggle detailed stack traces for development or clean messages for production.
- **Contextual Errors:** Includes support for inner errors and context strings.
- **Idiomatic:** Designed to fit naturally into standard Go control flows.

---

## Installation

```bash

go get github.com/4CHILL3S101/handleErr

or if specific version. you can do

go get [github.com/4CHILL3S101/handleErr@v1.0.0](https://github.com/4CHILL3S101/handleErr@v1.0.0)

```


## Usage Example
```bash
Basic Error Handling
The CheckErr function returns a boolean indicating if an error occurred, and a string formatted for the user.


package main

import (
    "errors"
    "fmt"
    "[github.com/4CHILL3S101/handleErr](https://github.com/4CHILL3S101/handleErr)"
)

func main() {
    err := errors.New("database connection failed")

    // CheckErr(err, userMessage, isDevelopment)
    if isErr, msg := handleErr.CheckErr(err, "Unable to load user data", true); isErr {
        fmt.Println(msg) 
        return           
    }

    fmt.Println("Success!") 
}

```


## Development vs. Production
```bash
Development Mode (true): Prints the stack trace and internal error details to the console for easier debugging.

Production Mode (false): Silently handles the error and only returns the high-level message for the user.
```



## Versioning

This library follows Semantic Versioning: vMAJOR.MINOR.PATCH.

To pin a specific version:
```bash
go get [github.com/4CHILL3S101/handleErr@v1.0.0](https://github.com/4CHILL3S101/handleErr@v1.0.0)
```


