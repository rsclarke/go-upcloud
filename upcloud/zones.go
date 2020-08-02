package upcloud

import (
	"context"
	"net/http"
)

// ZonesService handles communication with the zones related methods of the UpCloud API
// https://developers.upcloud.com/1.3/5-zones/
type ZonesService service

// Zone represents the description and properties of an Upcloud Zone
type Zone struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Public      string `json:"public"` // yes/no
}

// ZoneList represents the list of zones
type ZoneList struct {
	Zones []Zone `json:"zone"`
}

// ZoneListResponse represents the response from the ListZones API call
type ZoneListResponse struct {
	ZoneList *ZoneList `json:"zones"`
}

// ListAvailableZones returns the description of each zone.
// https://developers.upcloud.com/1.3/5-zones/#list-available-zones
func (s *ZonesService) ListAvailableZones(ctx context.Context) (*ZoneList, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "zone", nil)
	if err != nil {
		return nil, nil, err
	}

	zoneList := new(ZoneList)
	resp, err := s.client.Do(ctx, req, &ZoneListResponse{ZoneList: zoneList})
	if err != nil {
		return nil, resp, err
	}

	return zoneList, resp, nil
}
