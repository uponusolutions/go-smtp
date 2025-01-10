package textsmtp

import (
	"errors"

	"github.com/uponusolutions/go-smtp"
)

func IsPrintableASCII(val string) bool {
	for _, ch := range val {
		if ch < ' ' || '~' < ch {
			return false
		}
	}
	return true
}

func CheckNotifySet(values []smtp.DSNNotify) error {
	if len(values) == 0 {
		return errors.New("malformed NOTIFY parameter value")
	}

	seen := map[smtp.DSNNotify]struct{}{}
	for _, val := range values {
		switch val {
		case smtp.DSNNotifyNever, smtp.DSNNotifyDelayed, smtp.DSNNotifyFailure, smtp.DSNNotifySuccess:
			if _, ok := seen[val]; ok {
				return errors.New("malformed NOTIFY parameter value")
			}
		default:
			return errors.New("malformed NOTIFY parameter value")
		}
		seen[val] = struct{}{}
	}
	if _, ok := seen[smtp.DSNNotifyNever]; ok && len(seen) > 1 {
		return errors.New("malformed NOTIFY parameter value")
	}

	return nil
}
