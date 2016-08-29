package internetstatus_dns

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
	"github.com/moul/internet-status/pkg/root-hints"
)

func QueryRootDNSA() error {
	req := new(dns.Msg)
	req.Id = dns.Id()
	req.Question = make([]dns.Question, 1)
	req.Question[0] = dns.Question{".", dns.TypeNS, dns.ClassINET}
	server := roothints.RandomA()
	in, err := dns.Exchange(req, fmt.Sprintf("%s:53", server.A))
	log.Printf("QueryRootDNSA: in=%v, err=%v", in, err)
	return err
}

func QueryRootDNSAAAA() error {
	req := new(dns.Msg)
	req.Id = dns.Id()
	req.Question = make([]dns.Question, 1)
	req.Question[0] = dns.Question{".", dns.TypeNS, dns.ClassINET}
	server := roothints.RandomAAAA()
	in, err := dns.Exchange(req, fmt.Sprintf("[%s]:53", server.AAAA))
	log.Printf("QueryRootDNSA: in=%v, err=%v", in, err)
	return err
}
