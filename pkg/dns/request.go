package dns

type DnsRequest struct {
	X        float64
	Y        float64
	Z        float64
	Velocity float64
}

// Validates request values
func (r DnsRequest) Validate() error {
	// We may have bounded zones, so we should validate that coordinates remain in that zone.
	return nil
}
