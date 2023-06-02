package types

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

type ContactMethod interface {
}

type BaseContactMethod struct {
	ID      string            `json:"id"`
	Type    ContactMethodType `json:"type"`
	Summary string            `json:"summary"`
	// A short-form, server-generated string that provides succinct,
	// important information about an object suitable for primary labeling of an entity in a client.
	// In many cases, this will be identical to name, though it is not intended to be an identifier.
	Self string `json:"self"`
	// a URL at which the entity is uniquely displayed in the Web app
	HtmlURL string `json:"html_url"`
	// The "address" to deliver to: email, phone number, etc., depending on the type.
	Address string `json:"address"`
}

type PhoneContactMethod struct {
	BaseContactMethod
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
	BaseContactMethod

	DeviceType DeviceType `json:"device_type"`
	CreatedAt  string     `json:"created_at"`
	Sounds     []Sound    `json:"sounds"`
	Blacklist  string     `json:"blacklist"`
}

type EmailContactMethod struct {
	BaseContactMethod

	// Send an abbreviated email message instead of the standard email output. Useful for email-to-SMS gateways and email based pagers.
	// default is false
	SendShortEmail bool `json:"send_short_email"`
}
