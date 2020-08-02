package upcloud

import (
	"context"
	"net/http"
)

// PricingService handles communication with the pricing related methods of the UpCloud API
// https://developers.upcloud.com/1.3/4-pricing/
type PricingService service

// UnitPrice represents the cost per unit
type UnitPrice struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
}

// ZonePricing represents the pricing for services in the named zones
type ZonePricing struct {
	Name                   string    `json:"name"`
	Firewall               UnitPrice `json:"firewall"`
	IORequestBackup        UnitPrice `json:"io_request_backup"`
	IORequestHDD           UnitPrice `json:"io_request_hdd"`
	IORequestMaxIOPS       UnitPrice `json:"io_request_maxiops"`
	IPv4Address            UnitPrice `json:"ipv4_address"`
	IPv6Address            UnitPrice `json:"ipv6_address"`
	PublicIPv4BandwidthIn  UnitPrice `json:"public_ipv4_bandwidth_in"`
	PublicIpv4BandwidthOut UnitPrice `json:"public_ipv4_bandwidth_out"`
	PublicIPv6BandwidthIn  UnitPrice `json:"public_ipv6_bandwidth_in"`
	PublicIPv6BandwidthOut UnitPrice `json:"public_ipv6_bandwidth_out"`
	ServerCore             UnitPrice `json:"server_code"`
	ServerMemory           UnitPrice `json:"server_memory"`
	StorageBackup          UnitPrice `json:"storage_backup"`
	StorageHDD             UnitPrice `json:"storage_hdd"`
	StorageMaxIOPS         UnitPrice `json:"storage_maxiops"`
	StorageTemplate        UnitPrice `json:"storage_template"`
	// Server Plans
	// Simple backup plans
}

// PriceList represents the list of prices for each zone
type PriceList struct {
	ZonePrice []ZonePricing `json:"zone"`
}

// PriceListResponse represents the response from the Pricing.ListPrices API methods.
type PriceListResponse struct {
	PriceList *PriceList `json:"prices"`
}

// ListPrices returns the prices per zone
// https://developers.upcloud.com/1.3/4-pricing/#list-prices
func (s *PricingService) ListPrices(ctx context.Context) (*PriceList, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "price", nil)
	if err != nil {
		return nil, nil, err
	}

	priceList := new(PriceList)
	resp, err := s.client.Do(ctx, req, &PriceListResponse{PriceList: priceList})
	if err != nil {
		return nil, resp, err
	}

	return priceList, resp, nil
}
