package prose

import "strings"

// JoinWithCommas join all slice elements by comma (,), at the join with the word (and) just before the last word
// version with ifs else
//func JoinWithCommas(phrases []string) string {
//	if len(phrases) == 0 {
//		return ""
//	} else if len(phrases) == 1 {
//		return phrases[0]
//	} else if len(phrases) == 2 {
//		return phrases[0] + " and " + phrases[1]
//	} else {
//		result := strings.Join(phrases[:len(phrases)-1], ", ")
//		result += ", and "
//		result += phrases[len(phrases)-1]
//		return result
//	}
//}

// JoinWithCommas join all slice elements by comma (,), at the join with the word (and) just before the last word
// version with switch statement
func JoinWithCommas(phrases []string) string {
	switch len(phrases) {
	case 0:
		return ""
	case 1:
		return phrases[0]
	case 2:
		return phrases[0] + " and " + phrases[1]
	default:
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}
