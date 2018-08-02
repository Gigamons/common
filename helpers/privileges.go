package helpers

import "github.com/Gigamons/common/consts"

// HasPrivileges if user has those permissions else not!
func HasPrivileges(p uint64, u *consts.User) bool {
	return u.Privileges&p > 0
}
