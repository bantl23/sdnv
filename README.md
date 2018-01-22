SDNV
====

[![Build Status](https://travis-ci.org/bantl23/sdnv.svg?branch=master)](https://travis-ci.org/bantl23/sdnv)
[![Coverage Status](https://coveralls.io/repos/bantl23/sdnv/badge.svg?branch=master&service=github)](https://coveralls.io/github/bantl23/sdnv?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/bantl23/sdnv)](https://goreportcard.com/report/github.com/bantl23/sdnv)

This library implements the Self-Delimiting Numeric Values (SDNV) protocol.

The algorithm for implementation and testing was obtained from:

https://tools.ietf.org/html/rfc6256

## Usage

```
import (
  "fmt"
  "github.com/bantl23/sdnv"
)

func main() {
  s := sdnv.NewSdnv(0)

  // returns byte array
  data := s.Marshal()

  // sets sdnv.Value
  // sets sdnv.EncLen
  // return err if any
  err := s.Unmarshal(data)
  if err != nil {
    fmt.Println(err)
  }
}
```
