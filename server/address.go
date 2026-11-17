package server

import (
	"fmt"

	"github.com/uponusolutions/go-smtp/internal/parse"
)

// UTF8Clean is an address validator for WithAddressValidator which rejects
// addresses containing invisible or direction changing characters.
//
// RFC 6531 allows every UTF-8 character in an address, but some of them make
// an address look different from what it is, or end a line in logs:
//
//   - bidirectional formatting characters (U+202A to U+202E, U+2066 to U+2069)
//   - line and paragraph separators (U+0085, U+2028, U+2029)
//   - zero width space, word joiner and byte order mark (U+200B, U+2060, U+FEFF)
//
// Zero width joiner and non-joiner (U+200D, U+200C) and the directional marks
// (U+200E, U+200F, U+061C) are accepted, since they are needed to spell names
// in some scripts. Other spaces (e.g. U+00A0) and the Unicode normalization
// form are not checked.
func UTF8Clean(addr string) error {
	if parse.IsASCII(addr) {
		return nil
	}

	for _, r := range addr {
		switch {
		case r >= 0x202A && r <= 0x202E, // bidirectional embedding and override
			r >= 0x2066 && r <= 0x2069,            // bidirectional isolates
			r == 0x0085, r == 0x2028, r == 0x2029, // line and paragraph separators
			r == 0x200B, r == 0x2060, r == 0xFEFF: // zero width space, word joiner, BOM
			return fmt.Errorf("address contains the disallowed character %U", r)
		}
	}

	return nil
}
