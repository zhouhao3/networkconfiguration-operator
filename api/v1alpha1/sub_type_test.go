package v1alpha1

import "testing"

func TestAssess(t *testing.T) {
	n := NICHint{
		Speed:    1000,
		SmartNIC: false,
	}
	cases := []struct {
		n        NICHint
		expected float64
	}{
		{
			n: NICHint{
				Speed:    1100,
				SmartNIC: false,
			},
			expected: 10 + 1000.0/1100.0*100,
		},
		{
			n: NICHint{
				Speed:    1000,
				SmartNIC: true,
			},
			expected: 100,
		},
		{
			n: NICHint{
				Speed:    1000,
				SmartNIC: false,
			},
			expected: 110,
		},
		{
			n: NICHint{
				Speed:    900,
				SmartNIC: false,
			},
			expected: 0,
		},
	}

	for _, c := range cases {
		got := n.Assess(c.n)
		if c.expected != got {
			t.Errorf("expected: %v, got: %v", c.expected, got)
		}
	}
}
