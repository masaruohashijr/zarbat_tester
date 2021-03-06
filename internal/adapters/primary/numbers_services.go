package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/numbers"
)

type port_number struct {
	driven numbers.SecondaryPort
}

func NewNumbersService(driven numbers.SecondaryPort) numbers.PrimaryPort {
	return &port_number{
		driven,
	}
}

func (p *port_number) ViewNumber(n string) (*domains.IncomingPhoneNumber, error) {
	vn, err := p.driven.ViewNumber(n)
	return vn, err
}

func (p *port_number) AddNumber(n string) error {
	err := p.driven.AddNumber(n)
	return err
}

func (p *port_number) DeleteNumber(n string) error {
	err := p.driven.DeleteNumber(n)
	return err
}

func (p *port_number) UpdateNumber() error {
	err := p.driven.UpdateNumber()
	return err
}

func (p *port_number) ListAvailableNumbers() ([]string, error) {
	list, err := p.driven.ListAvailableNumbers()
	return list, err
}

func (p *port_number) ListNumbers() (*[]domains.IncomingPhoneNumber, error) {
	list, err := p.driven.ListNumbers()
	return list, err
}
