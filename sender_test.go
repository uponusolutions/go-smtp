package smtp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBATV(t *testing.T) {
	testCases := []struct {
		from string
		want string
	}{
		{
			"tester@internal.com",
			"tester@internal.com",
		},
		{
			"prvs=1940025abc=tester@internal.com",
			"tester@internal.com",
		},
		{
			"prvs=008310faaa=prtg-m365.microsoft@cloud-mgmt.net",
			"prtg-m365.microsoft@cloud-mgmt.net",
		},
		{
			"msprvs1=1202020UU5ulBw=bounces-tester@bounces.internal.com",
			"bounces-tester@bounces.internal.com",
		},
		{
			"btv1==TAG==USER@example.com",
			"USER@example.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.from, func(t *testing.T) {
			got := ParseBATV(tc.from)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseSRS(t *testing.T) {
	testCases := []struct {
		from string
		want string
	}{
		{
			"tester@internal.com",
			"tester@internal.com",
		},
		{
			"max.mustermann+SRS=StfaP=3J=uponu.io=eliza.musterfrau@uponu.cloud",
			"max.mustermann@uponu.cloud",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.from, func(t *testing.T) {
			got := ParseSRS(tc.from)
			assert.Equal(t, tc.want, got)
		})
	}
}
