package dns

import (
	"context"
	"math"
)

// Creates a new instance of DroneNavigationService
func NewDNS(sectionID int) *DroneNavigationService {
	return &DroneNavigationService{sectionID}
}

type DroneNavigationService struct {
	sectionID int
}

func (dns *DroneNavigationService) CalculateLocation(ctx context.Context, req DnsRequest) (*DnsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	secID := float64(dns.sectionID)
	loc := (req.X * secID) + (req.Y * secID) + (req.Z * secID) + req.Velocity

	// reducing precision to 2 floating points
	loc = math.Round(loc*100) / 100

	return &DnsResponse{loc}, nil
}
