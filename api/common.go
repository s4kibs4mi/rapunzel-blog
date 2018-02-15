package api

import "github.com/s4kibs4mi/rapunzel-blog/models"

func isUserStatusValid(status string) bool {
	switch {
	case status == string(models.UserStatusRegistered) ||
		status == string(models.UserStatusVerified) ||
		status == string(models.UserStatusBlocked):
		return true
	}
	return false
}

func isUserTypeValid(uType string) bool {
	switch {
	case uType == string(models.UserTypeParent) ||
		uType == string(models.UserTypeFamily) ||
		uType == string(models.UserTypeGhost):
		return true
	}
	return false
}
