package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	switch len(phrases) {
	case 0:
		return ""
	case 1:
		return phrases[0]
	case 2:
		return phrases[0] + " and " + phrases[1]
	default:
		head := strings.Join(phrases[:len(phrases)-1], ", ")
		return head + ", and " + phrases[len(phrases)-1]
	}
}

