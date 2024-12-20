# Parse

Parse is a functional parsing library written in Go. It enables the building of complex parsers from small, simple and (most importantly) unit testable functions that process the input.

The core of the library is a partial rewrite of [Adrian Hesketh's library](https://github.com/a-h/parse/) with reduced verbosity, bug fixes and then extended as required for [Notedown](https://github.com/notedownorg/notedown).

## Packages

- [`core`](./core) contains all the base parsers for parsing documents.
- [`time`](./time) contains all parsers related to time, dates and durations.
- [`test`](./test) contains helper functions for unit testing your own parsers.

The packages are designed to be composable via dot import. Dot imports are generally discouraged in Golang except in the case of reducing verbosity for DSL-like APIs which is typical here.

```go

import (
    . "github.com/liamawhite/parse/core"
    // . "github.com/liamawhite/parse/time" any other packages you may need
)

var BlankLine = Times(2, NewLine) // \n\n

```


## Implementing Your Own Parsers

To implement a parser implement the `Parser[T]` type alias, a function that takes an `Input` and returns `(T, bool, error)`. Each parser should attempt to parse the `Input` and roll back if it is unable to find what it is looking for.

You can find examples in the [`time`](./time) package.
