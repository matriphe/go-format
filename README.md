# Go Format

A helper to format date time in Go.

I'm fed up with repetitive formatting weird Go date format 🙄, so I decided to create a package to not repeat myself.

## Install

As usual, run: `go get github.com/matriphe/go-format`.

## Usage

### Date and Time

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/matriphe/go-format/datetime"
)

func main()  {
	theTime := time.Date(2023, 3, 5, 18, 45, 21, 0, time.UTC)
	
	fmt.Println(theTime.Format(datetime.ISODate)) // print '2023-03-05'
	fmt.Println(datetime.ToISODate(theTime)) // print '2023-03-05'
	_, _ = datetime.FromISODate("2023-03-05") // parse to time.Time

	fmt.Println(theTime.Format(datetime.ISOTime)) // print '18:45:21'
	fmt.Println(datetime.ToISOTime(theTime)) // print '18:45:21'
	_, _ = datetime.FromISOTime("18:45:21") // parse to time.Time

	fmt.Println(theTime.Format(datetime.ISOHourMin)) // print '18:45'
	fmt.Println(datetime.ToISOHourMin(theTime)) // print '18:45'
	_, _ = datetime.FromISOHourMin("18:45") // parse to time.Time

	fmt.Println(theTime.Format(datetime.ISODateTime)) // print '2023-03-05 18:45:21'
	fmt.Println(datetime.ToISODateTime(theTime)) // print '2023-03-05 18:45:21'
	_, _ = datetime.FromISODateTime("2023-03-05 18:45:21") // parse to time.Time
}
```

# License

This package is released under [MIT license](LICENSE).