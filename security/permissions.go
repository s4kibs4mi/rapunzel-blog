package security

import "github.com/s4kibs4mi/rapunzel-blog/models"

func HasLoginPermissions(u *models.User) bool {
	if u.UserType == models.Ghost ||
		u.UserStatus == models.Registered ||
		u.UserStatus == models.Blocked {
		return false
	}
	return true
}
