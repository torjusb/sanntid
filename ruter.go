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

type sanntidArrivalData struct {
	MonitoredVehicleJourney sanntidMonitoredVehicleJourney
}

func requestArrivalData(locationId int) ([]sanntidArrivalData, error) {
	var data []sanntidArrivalData

	url := fmt.Sprintf("http://reisapi.ruter.no/stopvisit/getdepartures/%d", locationId)

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
