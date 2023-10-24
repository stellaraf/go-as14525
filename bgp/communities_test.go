package bgp_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stellaraf/go-as14525/bgp"
	"github.com/stretchr/testify/require"
)

func Test_InverseNumberMatch(t *testing.T) {
	communities := make([]string, 999, 999)
	t.Run("fixtures", func(t *testing.T) {
		for i := 1; i <= 999; i++ {
			community := fmt.Sprintf("14525:51%03d", i)
			communities[i-1] = community
		}
	})
	for i := 1; i < 999; i++ {
		i := i
		this := fmt.Sprintf("14525:51%03d", i)
		t.Run(this, func(t *testing.T) {
			generated := bgp.InverseNumberMatch(i)
			pattern, err := regexp.Compile(fmt.Sprintf("14525:51%s", generated))
			require.NoError(t, err)
			for _, com := range communities {
				match := pattern.MatchString(com)
				if this != com {
					require.True(t, match, "pattern for %s should match %s", this, com)
				} else {
					require.False(t, match, "pattern for %s should NOT match %s", this, com)
				}
			}
		})
	}
}
