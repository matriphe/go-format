package datetime

import "time"

const (
	// ISODate returns ISO 8601 `YYYY-MM-DD` date layout.
	ISODate = "2006-01-02"
	// ISOTime returns ISO 8601 `HH:mm:ss` time layout.
	ISOTime = "15:04:05"
	// ISOHourMin returns ISO 8601 `HH:mm` time layout.
	ISOHourMin = "15:04"
	// ISODateTime returns ISO 8601 `YYYY-MM-DD HH:mm:ss` date time layout.
	ISODateTime = "2006-01-02 15:04:05"
)

// ToISODate converts time to ISO 8601 `YYYY-MM-DD` date format.
func ToISODate(t time.Time) string {
	return t.Format(ISODate)
}

// ToISOTime converts time to ISO 8601 `HH:mm:ss` time format.
func ToISOTime(t time.Time) string {
	return t.Format(ISOTime)
}

// ToISOHourMin converts time to ISO 8601 `HH:mm` time format.
func ToISOHourMin(t time.Time) string {
	return t.Format(ISOHourMin)
}

// ToISODateTime converts time to ISO 8601 `YYYY-MM-DD HH:mm:ss` date time format.
func ToISODateTime(t time.Time) string {
	return t.Format(ISODateTime)
}

// FromISODate parses date from ISO 8601 `YYYY-MM-DD` date format.
func FromISODate(s string) (time.Time, error) {
	return time.Parse(ISODate, s)
}

// FromISOTime parses date from ISO 8601 `HH:mm:ss` time format.
func FromISOTime(s string) (time.Time, error) {
	return time.Parse(ISOTime, s)
}

// FromISOHourMin parses date from ISO 8601 `HH:mm` time format.
func FromISOHourMin(s string) (time.Time, error) {
	return time.Parse(ISOHourMin, s)
}

// FromISODateTime parses date from ISO 8601 `YYYY-MM-DD HH:mm:ss` date time format.
func FromISODateTime(s string) (time.Time, error) {
	return time.Parse(ISODateTime, s)
}
