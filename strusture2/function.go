package simple

import "sort"

type Figure struct {
	Dimension int
	Type      string
	Color     string
}

// Sort - makes a copy of a slice of figures and sorts them.
func Sort(in []Figure) (out []Figure) {
	if in == nil {
		return
	}
	out = make([]Figure, len(in))
	copy(out, in)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Dimension == out[j].Dimension {
			if out[i].Type == out[j].Type {
				return out[i].Color < out[j].Color
			}
			return out[i].Type < out[j].Type
		}
		return out[i].Dimension < out[j].Dimension
	})
	return
}
