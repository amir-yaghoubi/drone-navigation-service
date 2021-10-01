package dns

import (
	"context"
)

type DroneNavigationServicer interface {
	CalculateLocation(ctx context.Context, req DnsRequest) (*DnsResponse, error)
}
