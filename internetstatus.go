package internetstatus

import (
	"github.com/moul/internet-status/pkg/dns"
)

type Result struct {
	QueryRootDNSA    error
	QueryRootDNSAAAA error
}

func Full() (Result, error) {
	result := Result{
		QueryRootDNSA:    internetstatus_dns.QueryRootDNSA(),
		QueryRootDNSAAAA: internetstatus_dns.QueryRootDNSAAAA(),
	}
	return result, nil
}
