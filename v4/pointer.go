package v4

func CheckStringJSONData(s string) *string {
	if len(s) > 0 {
		return &s
	}
	return nil
}

func CheckInt64JSONData(i int64) *int64 {
	if i > 0 {
		return &i
	}
	return nil
}

func CheckFloat64JSONData(f float64) *float64 {
	if f > 0 {
		return &f
	}
	return nil
}
