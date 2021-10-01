package http

import "drone_navigation_service/pkg/dns"

type DnsResponse struct {
	Location float64 `json:"loc"`
}

func ToHttpResponse(in dns.DnsResponse) DnsResponse {
	return DnsResponse{in.Location}
}
