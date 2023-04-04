package data

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Devices struct {
	Devices []Device `json:"result_list"`
}

type Device struct {
	DeviceID    string      `json:"device_id"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ActiveState string      `json:"active_state"`
	DisplayName string      `json:"display_name"`
	Online      bool        `json:"online"`
	DevicePoint DevicePoint `json:"latest_device_point"`
}

type DevicePoint struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Altitude float64 `json:"altitude"`
	Angle    int     `json:"angle"`
	Speed    float64 `json:"speed"`
}

func (d *Devices) GetAllDevices() (*Devices, error) {
	res, err := http.Get("")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var devices Devices
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return nil, err
	}

	return &devices, nil
}
