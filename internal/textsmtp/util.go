package textsmtp

import (
	"errors"

	"github.com/uponusolutions/go-smtp"
)

// IsPrintableASCII checks if string contains only printable ascii.
func IsPrintableASCII(val string) bool {
	for _, ch := range val {
		if ch < ' ' || '~' < ch {
			return false
		}
	}
	return true
}

// CheckNotifySet checks if a DSNNotify array isn't malformed.
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
