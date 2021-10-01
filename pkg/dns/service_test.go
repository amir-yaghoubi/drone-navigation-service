package dns_test

import (
	"context"
	"drone_navigation_service/pkg/dns"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateLocation(t *testing.T) {
	// Setup
	svc := dns.NewDNS(1)
	cases := []struct {
		description string
		req         dns.DnsRequest
		res         *dns.DnsResponse
		err         error
	}{
		{"Zero out values", dns.DnsRequest{}, &dns.DnsResponse{0}, nil},
		{"Negative values without floating point", dns.DnsRequest{-1, -1, -1, -1}, &dns.DnsResponse{-4}, nil},
		{"Positive values without floating point", dns.DnsRequest{1, 1, 1, 1}, &dns.DnsResponse{4}, nil},
		{"Negative values with floating point", dns.DnsRequest{-0.5, -0.5, -0.5, -0.2}, &dns.DnsResponse{-1.7}, nil},
		{"Positive values with floating point", dns.DnsRequest{0.5, 0.5, 0.5, 0.2}, &dns.DnsResponse{1.7}, nil},
		{"Mixed values", dns.DnsRequest{0.5, -0.5, 3, -2.5}, &dns.DnsResponse{0.5}, nil},
		{"Example given values", dns.DnsRequest{123.12, 456.56, 789.89, 20.0}, &dns.DnsResponse{1389.57}, nil},
	}

	for _, tc := range cases {
		actual, err := svc.CalculateLocation(context.Background(), tc.req)
		if tc.err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.description)
		}
		assert.Equal(t, actual, tc.res, tc.description)
	}
}
