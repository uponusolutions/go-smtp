package smtp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBATV(t *testing.T) {
	testCases := []struct {
		from  string
		regex string
		want  string
		err   bool
	}{
		{
			"tester@internal.com",
			BATVRegEx,
			"tester@internal.com",
			false,
		},
		{
			"tester@internal.com",
			"",
			"tester@internal.com",
			true,
		},

		{
			"prvs=1940025abc=tester@internal.com",
			BATVRegEx,
			"tester@internal.com",
			false,
		},
		{
			"prvs=008310faaa=prtg-m365.microsoft@cloud-mgmt.net",
			BATVRegEx,
			"prtg-m365.microsoft@cloud-mgmt.net",
			false,
		},
		{
			"msprvs1=1202020UU5ulBw=bounces-tester@bounces.internal.com",
			"(?i)prvs=[a-zA-Z0-9]*=(.*@.*)",
			"msprvs1=1202020UU5ulBw=bounces-tester@bounces.internal.com",
			false,
		},
		{
			"msprvs1=1202020UU5ulBw=bounces-tester@bounces.internal.com",
			BATVRegEx,
			"bounces-tester@bounces.internal.com",
			false,
		},
		{
			"msprvs1=1202020UU5ulBw=bounces-tester@bounces.internal.com",
			"",
			"msprvs1=1202020UU5ulBw=bounces-tester@bounces.internal.com",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.from, func(t *testing.T) {
			got, err := ParseBATV(tc.regex, tc.from)
			t.Logf("Got: %s, err %v", got, err)
			assert.Equal(t, tc.err, err != nil)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseBATVEx(t *testing.T) {
	s, err := ParseBATVEx(nil, "test")
	assert.Error(t, err)
	assert.Equal(t, "test", s)
}
