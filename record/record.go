package record

import (
	"fmt"

	"github.com/zhengyi13/go-ddns/domain"
)

type Record struct {
	Subdomain string
	Domain domain.Domain
	Type   string
	Name   string
}

func New(t string, s string, d domain.Domain) Record {
	return Record{
		Subdomain: s,
		Domain: d,
		Type: t,
	}
}

func (r Record) String() string {
	return fmt.Sprintf("Type %s, FQDN: %s.%s", r.Type, r.Subdomain, r.Domain.Name)
}
