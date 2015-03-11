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

#### `GetArrivals(int, sanntidDirection) ([]Arrival, error)`

Retrieve the arrivals for a specific location based on the location ID, which can be found in the [Ruter API](http://labs.trafikanten.no/how-to-use-the-api.aspx). It returns a slice of `Arrival` type or an `error` if an error occurred.

```go
data, err := sanntid.GetArrivals(3010536, sanntid.DirAny)

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
	Direction sanntidDirection
}
```
