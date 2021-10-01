package mocks

import (
	"context"
	"drone_navigation_service/pkg/dns"

	"github.com/stretchr/testify/mock"
)

type MockDroneNavigationService struct {
	mock.Mock
}

func (m *MockDroneNavigationService) CalculateLocation(ctx context.Context, req dns.DnsRequest) (*dns.DnsResponse, error) {
	ret := m.Called(ctx, req)

	var res *dns.DnsResponse

	if ret.Get(0) != nil {
		res = ret.Get(0).(*dns.DnsResponse)
	}

	return res, ret.Error(1)
}
