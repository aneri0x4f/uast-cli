package utils

import "slices"

func init() {
	slices.Sort(unAspiratedConsonants)
	slices.Sort(allowedSymbols)
	slices.Sort(iastAllowed)
}
