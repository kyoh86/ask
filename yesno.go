package ask

import "strconv"

// parseYesNo returns the boolean value represented by the string.
// It accepts y, Y, YES, yes, Yes, n, N, NO, no, No.
// Any other value returns an error.
func parseYesNo(str string) (bool, error) {
	switch str {
	case "y", "Y", "yes", "YES", "Yes":
		return true, nil
	case "n", "N", "no", "NO", "No":
		return false, nil
	}

	return false, &strconv.NumError{Func: "parseYesNo", Num: str, Err: strconv.ErrSyntax}
}
