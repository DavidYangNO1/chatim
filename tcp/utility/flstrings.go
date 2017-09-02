package utility

import "strconv"

func StrToInt(value string) (int, bool) {
	iValue, err := strconv.Atoi(value)
	if err == nil {
		return iValue, true
	}
	return 0, false
}

func StrToBool(value string) (bool, error) {
	iValue, err := strconv.Atoi(value)
	if err == nil {
		if iValue == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, err
}

func StrToUInt(value string) (uint64, bool) {
	UValue, err := strconv.ParseUint(value, 10, 0)
	if err == nil {
		return UValue, true
	}
	return 0, false
}

func StrToFloat(value string) (float64, bool) {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return floatValue, true
	}
	return 0.0, false
}
