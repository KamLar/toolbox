package toolbox

import "unicode"

//IsASCIIText return true if supplied string does not have binary data
func IsASCIIText(candidate string) bool {
	for _, r := range candidate {
		if r == '\n' || r == '\r' || r == '\t' {
			continue
		}
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
