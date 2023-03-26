# Go Format

A helper to format and convert many things in Go.

I'm fed up with the repetitive formatting of many things, especially the weird Go date format 🙄, so I decided to create a package to not repeat myself.

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

### Wind

This package converts wind [direction in degree to cardinal directional](http://snowfence.umn.edu/Components/winddirectionanddegrees.htm).

The reference of this function is from [Stackoverflow](https://stackoverflow.com/a/7490772/3460840).

```go
package main

import (
	"fmt"
	
	"github.com/matriphe/go-format/wind"
)

func main()  {
	fmt.Println(wind.Direction(90)) // print 'E'
	fmt.Println(wind.Direction(13)) // print 'NNE'
}
```

# License

This package is released under [MIT license](LICENSE).