package types

type UserRole string

const (
	UserRoleAdmin               UserRole = "admin"
	UserRoleLimitedUser         UserRole = "limited_user"
	UserRoleObserver            UserRole = "observer"
	UserRoleOwner               UserRole = "owner"
	UserRoleReadOnlyUser        UserRole = "read_only_user"
	UserRoleRestrictedAccess    UserRole = "restricted_access"
	UserRoleReadOnlyLimitedUser UserRole = "read_only_limited_user"
	UserRoleUser                UserRole = "user"
)

func (ur UserRole) String() string {
	return string(ur)
}
