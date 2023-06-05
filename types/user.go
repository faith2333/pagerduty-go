package types

type User struct {
	BaseObject
	Name string `json:"name"`
	// The user's email address.
	Email string `json:"email"`
	// The preferred time zone name. If null, the account's time zone will be used.
	TimeZone string `json:"time_zone"`
	// The schedule color.
	Color string `json:"color"`
	// The user role. Account must have the read_only_users ability to set a user as a read_only_user or a read_only_limited_user,
	// and must have advanced permissions abilities to set a user as observer or restricted_access.
	//    Allowed values:  admin  limited_user observer owner read_only_user restricted_access read_only_limited_user user
	Role UserRole `json:"role"`
	// The URL of the user's avatar.
	AvatarURL string `json:"avatar_url"`
	// The user's bio
	Description string `json:"description"`
	// If true, the user has an outstanding invitation.
	InvitationSent bool `json:"invitation_sent"`
	// The user's title.
	JobTitle string `json:"job_title"`
	// The list of teams to which the user belongs. Account must have the teams ability to set this.
	Teams []Team `json:"teams"`
	// The list of contact methods for the user.
	ContactMethods []ContactMethod `json:"contact_methods"`
	// The list of notification rules for the user.
	NotificationRules []BaseObjectReference `json:"notification_rules"`
	// The License assigned to the User
	License BaseObjectReference `json:"license"`
}

type CreateAndUpdateUserPayload struct {
	// The name of the user.
	Name string `json:"name"`
	// The user's email address.
	Email string `json:"email"`
	// The preferred zone names. if null, the account's time zone will be used.
	TimeZone string `json:"time_zone"`
	// The schedule color.
	Color string `json:"color"`
	// The user role. Account must have the read_only_users ability to set a user as a read_only_user or a read_only_limited_user,
	// and must have advanced permissions abilities to set a user as observer or restricted_access.
	//    Allowed values:  admin  limited_user observer owner read_only_user restricted_access read_only_limited_user user
	Role UserRole `json:"role"`
	// The user's bio.
	Description string `json:"description"`
	// The user's title.
	JobTitle string `json:"job_title"`
	// The License assigned to the User
	License BaseObjectReference `json:"license"`
}
type GetUserResp struct {
	User *User `json:"user"`
}
