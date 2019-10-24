# ricebox

This is migrations source module for
[github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate) package that
reads migration files from embedded
[go.rice](https://github.com/GeertJohan/go.rice) boxes.

This source driver must be loaded manually by using `WithInstance` constructor
instead of using source specific URL.

Expected module usage:

```go
src, err := ricebox.WithInstance(rice.MustFindBox("sql"))
if err != nil {
	log.Fatal(err)
}
m, err := migrate.NewWithSourceInstance("source", src, "mysql://user:pass@tcp(localhost:3306)/db")
if err != nil {
	log.Fatal(err)
}
```

This will **NOT** work:

```go
m, err := migrate.New("ricebox://sql", "mysql://user:pass@tcp(localhost:3306)/db")
if err != nil {
	log.Fatal(err)
}
```

Rice boxes can not be loaded by string URL for the following reasons:

1. If there are no calls to `rice.FindBox()` in  the main application **rice**
command line tool will not embed directory contents into the main binary. If
source URL were used then call to FindBox() would be included inside this
package instead of main application. Application could force **rice** tool to
embed directory contents by including a dummy method invocation like
`var _ = rice.MustFindBox("dir")` but that requires manual synchronizations of
source URL and this extra call.
2. Having `rice.FindBox()` inside this module would break development builds
where directory contents are not embedded into the final binary but read
directly from the file system. Current `rice.FindBox()` implementation tries to
detect name of the package that called this method by analyzing call stack
trace. Package name is later used to search for a directory relative to the
caller source code.

Example usage:

```go
package main

import (
        "fmt"
        "log"

        rice "github.com/GeertJohan/go.rice"
        ricebox "github.com/fln/migrate-ricebox"
        "github.com/golang-migrate/migrate/v4"
        _ "github.com/golang-migrate/migrate/v4/database/mysql"
)

func main() {
        src, err := ricebox.WithInstance(rice.MustFindBox("sql"))
        if err != nil {
                log.Fatalf("opening sql box: %v", err)
        }
        m, err := migrate.NewWithSourceInstance("source", src, "mysql://user:pass@tcp(localhost:3306)/db")
        if err != nil {
                log.Fatalf("creating migrate instance: %v", err)
        }
        if err = m.Up(); err != nil {
                log.Fatalf("executing up migration: %v", err)
        }
        err1, err2 := m.Close()
        if err1 != nil {
                log.Fatalf("closing source driver: %v", err1)
        }
        if err2 != nil {
                log.Fatalf("closing db driver: %v", err2)
        }
}
```
