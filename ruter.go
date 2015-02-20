package sanntid

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type sanntidMonitoredCall struct {
	ExpectedArrivalTime string
	DeparturePlatformName string
	DestinationDisplay string
}

type sanntidMonitoredVehicleJourney struct {
	DestinationName string
	MonitoredCall sanntidMonitoredCall
	PublishedLineName string
	VehicleMode int
}

// ArrivalData cointains the parsed data returned from a request to
// Ruter's API.
type ArrivalData struct {
	MonitoredVehicleJourney sanntidMonitoredVehicleJourney
}

// RequestArrivalData retrieves information about the upcoming arrivals for
// a given location based on its locationId.
func RequestArrivalData(locationID int) ([]ArrivalData, error) {
	var data []ArrivalData

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

	json.Unmarshal(body, &data)

	return data, err
}
