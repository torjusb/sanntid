package sanntid

type Line struct {
	Name        string
	Destination string
}

type Arrival struct {
	Line                Line
	ExpectedArrivalTime string
	Platform            string
}

func GetArrivals(locationId int) ([]Arrival, error) {
	var arrivals []Arrival

	data, err := requestArrivalData(locationId)

	if err == nil {
		for i := 0; i < len(data); i++ {
			line := Line{
				data[i].MonitoredVehicleJourney.PublishedLineName,
				data[i].MonitoredVehicleJourney.DestinationName,
			}
			arrival := Arrival{
				line,
				data[i].MonitoredVehicleJourney.MonitoredCall.ExpectedArrivalTime,
				data[i].MonitoredVehicleJourney.MonitoredCall.DeparturePlatformName,
			}

			arrivals = append(arrivals, arrival)
		}
	}

	return arrivals, err
}
