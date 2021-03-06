package mms

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	SendMMS(to, from, message string) (domains.Mms, error)
	ViewMMS(smsSid string) (domains.Mms, error)
	ListMMS(from, to string) ([]domains.Mms, error)
}

type SecondaryPort interface {
	SendMMS(to, from, message string) (domains.Mms, error)
	ViewMMS(mmsSid string) (domains.Mms, error)
	ListMMS(from, to string) ([]domains.Mms, error)
}
