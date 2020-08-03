package upcloud

import (
	"context"
	"net/http"
)

// TimezonesService handles communication with timezones related methods of the UpCloud API
// https://developers.upcloud.com/1.3/6-timezones/
type TimezonesService service

// TimezoneList represents the list of timezones
type TimezoneList struct {
	Timezones []string `json:"timezone"`
}

// TimezoneListResponse represents the response from the ListTimezones API call
type TimezoneListResponse struct {
	TimezoneList *TimezoneList `json:"timezones"`
}

// ListTimezones returns the description of each zone.
// https://developers.upcloud.com/1.3/6-timezones/#list-timezones
func (s *TimezonesService) ListTimezones(ctx context.Context) (*TimezoneList, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "timezone", nil)
	if err != nil {
		return nil, nil, err
	}

	timezoneList := new(TimezoneList)
	resp, err := s.client.Do(ctx, req, &TimezoneListResponse{TimezoneList: timezoneList})
	if err != nil {
		return nil, resp, err
	}

	return timezoneList, resp, nil
}
