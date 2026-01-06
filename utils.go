package mangadex

import "time"

func String(str string) *string {
	if len(str) == 0 {
		return nil
	}
	return &str
}

func Int(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func Bool(b bool) *bool {
	if !b {
		return nil
	}
	return &b
}

func Int64(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}

func Float64(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

func Uint(u uint) *uint {
	if u == 0 {
		return nil
	}
	return &u
}

func Time(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func Ptr[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}
