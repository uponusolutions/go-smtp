package tester

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMail_LookupKey(t *testing.T) {
	testCases := []struct {
		from       string
		recipients []string
		expected   string
	}{
		{"", []string{}, "+"},
		{"alice@i.com", []string{"bob@e.com"}, "alice@i.com+bob@e.com"},
		{"alice@i.com", []string{"bob@e.com", "mal@ev.com"}, "alice@i.com+bob@e.com+mal@ev.com"},
	}

	for i, tc := range testCases {
		run := fmt.Sprintf("run %d - %s", i, tc.from)
		t.Run(run, func(t *testing.T) {
			m := Mail{
				From:       tc.from,
				Recipients: tc.recipients,
			}
			assert.Equal(t, tc.expected, m.LookupKey())
		})
	}
}
