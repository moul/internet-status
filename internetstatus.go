package internetstatus

import (
	"github.com/moul/internet-status/pkg/dns"
)

// Result contains checks returns
type Result struct {
	QueryRootDNSA    error
	QueryRootDNSAAAA error
}

// Full performs all the checks to determine the connectivity status (most accurate)
func Full() (Result, error) {
	result := Result{
		QueryRootDNSA:    internetstatus_dns.QueryRootDNSA(),
		QueryRootDNSAAAA: internetstatus_dns.QueryRootDNSAAAA(),
	}
	return result, nil
}

// Quick performs minimal checks to determine the connectivity status (less accurate)
func Quick() (Result, error) {
	result := Result{}
	return result, nil
}
