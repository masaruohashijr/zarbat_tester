package notifications

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	SendNotification(to, from, message string) error
	ViewNotification(smsSid string) (domains.Notification, error)
	ListNotifications(from, to string) ([]domains.Notification, error)
}

type SecondaryPort interface {
	SendNotification(to, from, message string) error
	ViewNotification(mmsSid string) (domains.Notification, error)
	ListNotifications(from, to string) ([]domains.Notification, error)
}