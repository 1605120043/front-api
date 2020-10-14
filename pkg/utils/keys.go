package utils

import (
	"fmt"
)

func MemberTokenKey(userId uint64) string {
	return fmt.Sprintf("goshop:member:token::%d", userId)
}

func SendValidateCode(mobile string) string {
	return fmt.Sprintf("goshop:member:sms::%s", mobile)
}
