package dns

import (
	"context"
	"math"
)

// Creates a new instance of DroneNavigationService
func NewDNS(sectorID int) *DroneNavigationService {
	return &DroneNavigationService{sectorID}
}

type DroneNavigationService struct {
	sectorID int
}

func (dns *DroneNavigationService) CalculateLocation(ctx context.Context, req DnsRequest) (*DnsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	secID := float64(dns.sectorID)
	loc := (req.X * secID) + (req.Y * secID) + (req.Z * secID) + req.Velocity

	// reducing precision to 2 floating points
	loc = math.Round(loc*100) / 100

	return &DnsResponse{loc}, nil
}
