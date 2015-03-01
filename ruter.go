package sanntid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// sanntidDirection defines the direction of the vehicle. It is either,
// 0 (undefined (?)), 1 or 2.
type sanntidDirection int

type sanntidMonitoredCall struct {
	ExpectedArrivalTime   string
	DeparturePlatformName string
	DestinationDisplay    string
}

type sanntidMonitoredVehicleJourney struct {
	DestinationName   string
	MonitoredCall     sanntidMonitoredCall
	PublishedLineName string
	VehicleMode       int
	DirectionRef      sanntidDirection `json:",string"`
}

// ArrivalData cointains the parsed data returned from a request to
// Ruter's API.
type sanntidArrivalData struct {
	MonitoredVehicleJourney sanntidMonitoredVehicleJourney
}

// Ruter
type Ruter struct{}

// RequestArrivalData retrieves information about the upcoming arrivals for
// a given location based on its locationId.
func (ruter *Ruter) requestArrivalData(locationID int) ([]sanntidArrivalData, error) {
	url := fmt.Sprintf("http://reisapi.ruter.no/stopvisit/getdepartures/%d", locationID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ruter.parseArrivalData(body), err
}

func (ruter *Ruter) parseArrivalData(content []byte) []sanntidArrivalData {
	var data []sanntidArrivalData

	json.Unmarshal(content, &data)

	return data
}
