package cervell

import (
	"context"
)

// Placemark is a placemark for geocoding.
type Placemark struct {
	Name               string  `json:"name"`
	Thoroughfare       string  `json:"thoroughfare"`
	Locality           string  `json:"locality"`
	Country            string  `json:"country"`
	AdministrativeArea string  `json:"administrativeArea"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
}

// GeocodeAddress returns a list of placemarks for the provided address.
func (cl *Client) GeocodeAddress(ctx context.Context, address string) ([]*Placemark, error) {
	r, err := postCall[[]*Placemark](ctx, cl, "/geocoding/forward", &struct {
		Address string `json:"address"`
	}{
		Address: address,
	})
	if err != nil {
		return nil, err
	}
	return *r, nil
}

// GeocodeReverse returns a list of placemarks for the provided latitude and longitude.
func (cl *Client) GeocodeReverse(ctx context.Context, latitude, longitude float64) ([]*Placemark, error) {
	r, err := postCall[[]*Placemark](ctx, cl, "/geocoding/reverse", &struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		Latitude:  latitude,
		Longitude: longitude,
	})
	if err != nil {
		return nil, err
	}
	return *r, nil
}
