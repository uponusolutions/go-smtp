// Package smtp contains shared code between client and server implementations e.g. Status
package smtp

// BodyType describes the type of the body.
type BodyType string

const (
	// Body7Bit means the body type is 7BIT
	Body7Bit BodyType = "7BIT"
	// Body8BitMIME means the body type is 8BITMIME
	Body8BitMIME BodyType = "8BITMIME"
	// BodyBinaryMIME means the body type is BINARYMIME
	BodyBinaryMIME BodyType = "BINARYMIME"
)

// DSNReturn describes the DSN return.
type DSNReturn string

const (
	// DSNReturnFull means DNS return is full.
	DSNReturnFull DSNReturn = "FULL"
	// DSNReturnHeaders means DNS return is hdrs.
	DSNReturnHeaders DSNReturn = "HDRS"
)

// MailOptions contains parameters for the MAIL command.
type MailOptions struct {
	// Value of BODY= argument, 7BIT, 8BITMIME or BINARYMIME.
	Body BodyType

	// Size of the body. Can be 0 if not specified by client.
	Size int64

	// TLS is required for the message transmission.
	//
	// The message should be rejected if it can't be transmitted
	// with TLS.
	RequireTLS bool

	// The message envelope or message header contains UTF-8-encoded strings.
	// This flag is set by SMTPUTF8-aware (RFC 6531) client.
	UTF8 bool

	// Value of RET= argument, FULL or HDRS.
	Return DSNReturn

	// Envelope identifier set by the client.
	EnvelopeID string

	// Accepted Domain from Exchange Online, e.g. from OutgoingConnector
	XOORG *string

	// The authorization identity asserted by the message sender in decoded
	// form with angle brackets stripped.
	//
	// nil value indicates missing AUTH, non-nil empty string indicates
	// AUTH=<>.
	//
	// Defined in RFC 4954.
	Auth *string
}

// VrfyOptions contains parameters for the VRFY command.
type VrfyOptions struct {
	// The message envelope or message header contains UTF-8-encoded strings.
	// This flag is set by SMTPUTF8-aware (RFC 6531) client.
	UTF8 bool
}

// DSNNotify describes the DSN notify.
type DSNNotify string

const (
	// DSNNotifyNever sets the DSN notify to never.
	DSNNotifyNever DSNNotify = "NEVER"
	// DSNNotifyDelayed sets the DSN notify to delay.
	DSNNotifyDelayed DSNNotify = "DELAY"
	// DSNNotifyFailure sets the DSN notify to failure.
	DSNNotifyFailure DSNNotify = "FAILURE"
	// DSNNotifySuccess sets the DSN notify to succes.
	DSNNotifySuccess DSNNotify = "SUCCESS"
)

// DSNAddressType describes the DSN address type.
type DSNAddressType string

const (
	// DSNAddressTypeRFC822 means that the DSN address type is RFC822.
	DSNAddressTypeRFC822 DSNAddressType = "RFC822"
	// DSNAddressTypeUTF8 means that the DSN address type is UTF-8.
	DSNAddressTypeUTF8 DSNAddressType = "UTF-8"
)

// RcptOptions contains parameters for the RCPT command.
type RcptOptions struct {
	// Value of NOTIFY= argument, NEVER or a combination of either of
	// DELAY, FAILURE, SUCCESS.
	Notify []DSNNotify

	// Original recipient set by client.
	OriginalRecipientType DSNAddressType
	OriginalRecipient     string
}
