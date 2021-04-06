package helper

import "time"

func SetPointerString(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}

func SetPointerInt64(val int64) *int64 {
	if val == 0 {
		return nil
	}
	return &val
}

func SetPointerTime(val time.Time) *time.Time {
	return &val
}

func GetStringFromPointer(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

func GetInt64FromPointer(val *int64) int64 {
	if val == nil {
		return 0
	}
	return *val
}

func GetTimeFromPointer(val *time.Time) time.Time {
	return *val
}

func GetCurrentTimeUTC() (standartTime time.Time, unixTime int64) {
	current := time.Now().UTC()
	return current, current.Unix()
}
