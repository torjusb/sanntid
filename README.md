# Sanntid

Go package for retrieving realtime arrival data from the [Ruter API](http://labs.trafikanten.no/how-to-use-the-api.aspx).

## Installation

Get the package.

```shell
go get github.com/michaelenger/sanntid
```

## Usage

Import the package into your Go program.

```go
import (
	"github.com/michaelenger/sanntid"
)
```

Use the `GetArrivals` function to retrieve arrivals for a location based on the location ID.

```go
data, err := sanntid.GetArrivals(3010536)
```

### API

#### `GetArrivals(int) ([]Arrival, error)`

Retrieve the arrivals for a specific location based on the location ID, which can be found in the [Ruter API](http://labs.trafikanten.no/how-to-use-the-api.aspx). It returns a slice of `Arrival` type or an `error` if an error occurred.

```go
data, err := sanntid.GetArrivals(3010536)

if err == nil {
	for i := 0; i < len(data); i++ {
		fmt.Printf(
			"%s %s: %s\n",
			data[i].Line.Name,
			data[i].Line.Destination,
			data[i].ExpectedArrivalTime,
		)
	}
}
```

### Types

```go
type Arrival struct {
	Line Line
	ExpectedArrivalTime string
	Platform string
}
```

```go
type Line struct {
	Name string
	Destination string
}
```

## License

Copyright (c) 2015 Michael Enger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
