package http

import (
	"drone_navigation_service/pkg/dns"
	"strconv"
)

type DnsRequest struct {
	X        string `json:"x" binding:"required,numeric" default:"123.12"`
	Y        string `json:"y" binding:"required,numeric" default:"456.56"`
	Z        string `json:"z" binding:"required,numeric" default:"789.89"`
	Velocity string `json:"vel" binding:"required,numeric" default:"20.0"`
}

// Transform Handler DnsRequest to domain dns request
func ToDomainDnsRequest(r DnsRequest) (*dns.DnsRequest, error) {
	x, err := strconv.ParseFloat(r.X, 64)
	if err != nil {
		return nil, err
	}

	y, err := strconv.ParseFloat(r.Y, 64)
	if err != nil {
		return nil, err
	}

	z, err := strconv.ParseFloat(r.Z, 64)
	if err != nil {
		return nil, err
	}

	vel, err := strconv.ParseFloat(r.Velocity, 64)
	if err != nil {
		return nil, err
	}

	return &dns.DnsRequest{X: x, Y: y, Z: z, Velocity: vel}, nil
}
