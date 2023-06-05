package types

import "time"

const (
	ContactMethodTypePhone            ContactMethodType = "phone_contact_method"
	ContactMethodTypeEmail            ContactMethodType = "email_contact_method"
	ContactMethodTypePushNotification ContactMethodType = "push_notification_contact_method"
	ContactMethodTypeSMS              ContactMethodType = "sms_contact_method"
)

type ContactMethodType string

func (cm ContactMethodType) String() string {
	return string(cm)
}

type ContactMethod struct {
	BaseContactMethod
	PhoneContactMethod
	PushContactMethod
	EmailContactMethod
}

type BaseContactMethod struct {
	BaseObject
	Type    ContactMethodType `json:"type"`
	Address string            `json:"address"`
}

type PhoneContactMethod struct {
	// The 1-to-3 digit country calling code. >=1 and <= 1999
	CountryCode int `json:"country_code"`
	// If true, this phone is capable of receiving SMS messages.
	Enabled bool `json:"enabled"`
	// If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it.
	Blacklist bool `json:"blacklist"`
	// The label (e.g., "Work", "Mobile", etc.).
	Label string `json:"label"`
}

type DeviceType string

const (
	DeviceTypeIOS     DeviceType = "ios"
	DeviceTypeAndroid DeviceType = "android"
)

func (dt DeviceType) String() string {
	return string(dt)
}

type PushContactMethod struct {
	DeviceType DeviceType `json:"device_type"`
	CreatedAt  string     `json:"created_at"`
	Sounds     []Sound    `json:"sounds"`
	Blacklist  string     `json:"blacklist"`
}

type EmailContactMethod struct {
	// Send an abbreviated email message instead of the standard email output. Useful for email-to-SMS gateways and email based pagers.
	// default is false
	SendShortEmail bool `json:"send_short_email"`
}

type CreateAndUpdateContactMethodPayload struct {
	Type ContactMethodType `json:"type"`

	// The 1-to-3 digit country calling code. used only for phone contact method
	CountryCode int `json:"country_code"`
	// for push contact method only
	Sounds []Sound `json:"sounds"`
	// Time at which the contact method was created. push contact method only
	CreatedAt time.Time `json:"created_at"`
	// Send an abbreviated email message instead of the standard email output.
	// Useful for email-to-SMS gateways and email based pagers.
	// default: false
	SendShortEmail bool `json:"send_short_email"`

	// The label (e.g., "Work", "Mobile", etc.).
	Label string `json:"label"`
	// The "address" to deliver to: email, phone number, etc., depending on the type.
	Address string `json:"address"`
}

type GetContactMethodResp struct {
	ContactMethod *ContactMethod `json:"contact_method"`
}

type ListContactMethodsResp struct {
	ContactMethods []*ContactMethod `json:"contact_methods"`
}
