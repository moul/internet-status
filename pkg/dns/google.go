package internetstatus_dns

import "github.com/miekg/dns"

func QueryGoogleDNSA() error {
	req := new(dns.Msg)
	req.Id = dns.Id()
	req.Question = make([]dns.Question, 1)
	req.Question[0] = dns.Question{".", dns.TypeNS, dns.ClassINET}
	_, err := dns.Exchange(req, "8.8.8.8:53")
	return err
}
