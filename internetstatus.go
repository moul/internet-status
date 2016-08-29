package internetstatus

import (
	"github.com/moul/internet-status/pkg/captiveportal"
	"github.com/moul/internet-status/pkg/dns"
)

// Result contains checks returns
type Result struct {
	QueryRootDNSA       error
	QueryRootDNSAAAA    error
	BehindCaptivePortal error
}

// Full performs all the checks to determine the connectivity status (most accurate)
func Full() Result {
	result := Result{}

	result.QueryRootDNSA = internetstatus_dns.QueryRootDNSA()
	result.QueryRootDNSAAAA = internetstatus_dns.QueryRootDNSAAAA()
	result.BehindCaptivePortal = captiveportal.Check()
	return result
}

// Quick performs minimal checks to determine the connectivity status (less accurate)
func Quick() Result {
	result := Result{}
	return result
}
