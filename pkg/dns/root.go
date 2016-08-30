package internetstatus_dns

import (
	"fmt"

	"github.com/miekg/dns"
	"github.com/moul/internet-status/pkg/root-hints"
)

func QueryRootDNSA() error {
	req := new(dns.Msg)
	req.Id = dns.Id()
	req.Question = make([]dns.Question, 1)
	req.Question[0] = dns.Question{".", dns.TypeNS, dns.ClassINET}
	server := roothints.RandomA()
	_, err := dns.Exchange(req, fmt.Sprintf("%s:53", server.A))
	return err
}

func QueryRootDNSAAAA() error {
	req := new(dns.Msg)
	req.Id = dns.Id()
	req.Question = make([]dns.Question, 1)
	req.Question[0] = dns.Question{".", dns.TypeNS, dns.ClassINET}
	server := roothints.RandomAAAA()
	_, err := dns.Exchange(req, fmt.Sprintf("[%s]:53", server.AAAA))
	return err
}
