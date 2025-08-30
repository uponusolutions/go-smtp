package smtp

import (
	"regexp"
)

/*
   https://github.com/moisseev/rspamd/blob/master/rules/misc.lua

    Detect PRVS/BATV addresses to avoid FORGED_SENDER
    https://en.wikipedia.org/wiki/Bounce_Address_Tag_Validation

    Signature syntax:
        prvs=TAG=USER@example.com       BATV draft (https://tools.ietf.org/html/draft-levine-smtp-batv-01)
        prvs=USER=TAG@example.com
        btv1==TAG==USER@example.com     Barracuda appliance
        msprvs1=TAG=USER@example.com    Sparkpost email delivery service
*/

const (
	regexpBATV = "^(?:(?:prvs|msprvs1)=[^=]+=|btv1==[^=]+==)([^@]+@(?:[^@]+))$"
	regexpSRS  = "^([^+]+)\\+SRS=[^=]+=[^=]+=[^=]+=[^@]+@([^@]+)$"
)

var compiledRegexpBATV = regexp.MustCompile(regexpBATV)
var compiledRegexpSRS = regexp.MustCompile(regexpSRS)

// ParseBATV parses src to extract a BATV address.
// When BATV extration is not possible/needed src is returned.
func ParseBATV(src string) string {
	res := compiledRegexpBATV.FindStringSubmatch(src)
	if len(res) == 2 {
		return res[1]
	}
	return src
}

// ParseSRS parses src to extract the forwarding sender from SRS (Exchange Online).
// When SRS extration is not possible/needed src is returned.
func ParseSRS(src string) string {
	res := compiledRegexpSRS.FindStringSubmatch(src)
	if len(res) == 3 {
		return res[1] + "@" + res[2]
	}
	return src
}

// ParseSender combines ParseSRS and ParseBATV.
func ParseSender(src string) string {
	return ParseBATV(ParseSRS(src))
}
