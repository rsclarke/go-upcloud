package upcloud

import (
	"context"
	"net/http"
)

// PlansService handles communication ith the plans related methods of the UpCloud API
// https://developers.upcloud.com/1.3/7-plans/
type PlansService service

// Plan represents the configuration of an UpCloud Plan
type Plan struct {
	CoreNumber       int    `json:"core_number"`
	MemoryAmount     int    `json:"memory_amount"`
	Name             string `json:"name"`
	PublicTrafficOut int    `json:"public_traffic_out"`
	StorageSize      int    `json:"storage_size"`
	StorageTier      string `json:"storage_tier"`
}

// PlanList represents the list of zones
type PlanList struct {
	Plans []Plan `json:"plan"`
}

// PlanListResponse represents the response from the ListPlans API call
type PlanListResponse struct {
	PlanList *PlanList `json:"plans"`
}

// ListAvailablePlans returns the description of each zone.
// https://developers.upcloud.com/1.3/5-zones/#list-available-zones
func (s *PlansService) ListAvailablePlans(ctx context.Context) (*PlanList, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "plan", nil)
	if err != nil {
		return nil, nil, err
	}

	planList := new(PlanList)
	resp, err := s.client.Do(ctx, req, &PlanListResponse{PlanList: planList})
	if err != nil {
		return nil, resp, err
	}

	return planList, resp, nil
}
