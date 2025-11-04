package utils

import "unicode/utf8"

func HidePhoneNumber(phone string) string {
	if utf8.RuneCountInString(phone) != 11 {
		return phone
	}
	return phone[:3] + "****" + phone[7:]
}
