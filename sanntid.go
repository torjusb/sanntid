package sanntid

const (
	// DirAny will give you Line in any direction.
	DirAny sanntidDirection = iota

	// DirUp will give you Line in only one direction.
	DirUp

	// DirDown will give you Line in only one direction, reverse of DirUp.
	DirDown
)

type Line struct {
	Name        string
	Destination string
	Direction   sanntidDirection
}

type Arrival struct {
	Line                Line
	ExpectedArrivalTime string
	Platform            string
}

func GetArrivals(locationId int, direction sanntidDirection) ([]Arrival, error) {
	var arrivals []Arrival

	data, err := requestArrivalData(locationId)

	if err == nil {
		for i := 0; i < len(data); i++ {
			lineDir := data[i].MonitoredVehicleJourney.DirectionRef
			if direction == DirAny || direction == lineDir {
				line := Line{
					data[i].MonitoredVehicleJourney.PublishedLineName,
					data[i].MonitoredVehicleJourney.DestinationName,
					lineDir,
				}
				arrival := Arrival{
					line,
					data[i].MonitoredVehicleJourney.MonitoredCall.ExpectedArrivalTime,
					data[i].MonitoredVehicleJourney.MonitoredCall.DeparturePlatformName,
				}

				arrivals = append(arrivals, arrival)
			}
		}
	}

	return arrivals, err
}
