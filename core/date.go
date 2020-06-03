package core

import (
	"errors"
	"time"
)

func FormatDate(s string) (t time.Time, err error) {
	layout := "2006-01-02 15:04:05"
	t, err = time.Parse(layout, s)
	if err != nil {
		err = errors.New("invalid format date. Example format date 2020-06-03 21:41:46")
		return
	}

	return
}
