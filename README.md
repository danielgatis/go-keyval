# go-keyval

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-keyval?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-keyval)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-keyval/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-keyval)
[![Release](https://img.shields.io/github/release/danielgatis/go-keyval.svg?style=flat-square)](https://github.com/danielgatis/go-keyval/releases/latest)

A Variadic helper to mimic named arguments.

## Install

```bash
go get -u github.com/danielgatis/go-keyval
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-keyval"
```

### Example

An example described below is one of the use cases.

```go
package main

import (
  "fmt"
  "time"

  "github.com/danielgatis/go-keyval"
)

func main() {
  now := time.Now().UTC()
  email := "johndoe@gmail.com"
  query := "select * from users where email = :email and created_at < :created_at"

  execNamed(query, "email", email, "created_at", now)
}

func execNamed(query string, args... interface{}) {
  m := keyval.ToMap(args...)
  
  fmt.Printf("%s", m["email"])
  fmt.Printf("%s", m["created_at"])

  // db stuff here ...
}
```

## License

Copyright (c) 2020-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)
