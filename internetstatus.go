package internetstatus

import (
	"fmt"

	"github.com/moul/internet-status/pkg/captiveportal"
	"github.com/moul/internet-status/pkg/checkhttp"
	"github.com/moul/internet-status/pkg/dns"
)

var errorCheckNotPerformed = fmt.Errorf("check not performed")

// Result contains checks returns
type Result struct {
	Type string

	BehindCaptivePortal error
	QueryRootDNSA       error
	QueryRootDNSAAAA    error
	CheckFacebookHTTP   error
}

// IPv4Access returns true if IPv4 connectivity is available
func (r *Result) IPv4Access() bool {
	return r.QueryRootDNSA == nil && r.BehindCaptivePortal == nil
}

// IPv6Access returns true if IPv6 connectivity is available
func (r *Result) IPv6Access() bool {
	return r.QueryRootDNSAAAA == nil
}

// Access returns true if IPv4 or IPv6 connectivity is available
func (r *Result) Access() bool {
	return r.IPv4Access() || r.IPv6Access()
}

// Full performs all the checks to determine the connectivity status (most accurate)
func Full() Result {
	return Result{
		Type:                "full",
		BehindCaptivePortal: captiveportal.Check(),
		CheckFacebookHTTP:   checkhttp.CheckFacebookHTTP(),
		QueryRootDNSA:       internetstatus_dns.QueryRootDNSA(),
		QueryRootDNSAAAA:    internetstatus_dns.QueryRootDNSAAAA(),
	}
}

// Quick performs minimal checks to determine the connectivity status (less accurate)
func Quick() Result {
	return Result{
		Type:                "full",
		BehindCaptivePortal: captiveportal.Check(),
		CheckFacebookHTTP:   checkhttp.CheckFacebookHTTP(),
		QueryRootDNSA:       internetstatus_dns.QueryRootDNSA(),
		QueryRootDNSAAAA:    errorCheckNotPerformed,
	}
}
