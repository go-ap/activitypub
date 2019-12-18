package activitypub

import (
	"bytes"
	"fmt"
	"time"
)

const day = time.Hour * 24
const week = day * 7
const month = week * 4
const year = month * 12

func  Days(d time.Duration) float64 {
	dd := d / day
	h := d % day
	return float64(dd) + float64(h)/(24*60*60*1e9)
}
func  Weeks(d time.Duration) float64 {
	w := d / week
	dd := d % week
	return float64(w) + float64(dd)/(7*24*60*60*1e9)
}
func  Months(d time.Duration) float64 {
	m := d / month
	w := d % month
	return float64(m) + float64(w)/(4*7*24*60*60*1e9)
}
func  Years(d time.Duration) float64 {
	y := d / year
	m := d % year
	return float64(y) + float64(m)/(12*4*7*24*60*60*1e9)
}

func marshalXSD(d time.Duration)  ([]byte, error) {
	if d == 0 {
		return []byte{'P','T','0','S'}, nil
	}

	neg := d < 0
	if neg {
		d = -d
	}
	y := Years(d)
	d -= time.Duration(y) * year
	m := Months(d)
	d -= time.Duration(m) * month
	dd := Days(d)
	d -= time.Duration(dd) * day
	H := d.Hours()
	d -= time.Duration(H) * time.Hour
	M := d.Minutes()
	d -= time.Duration(M) * time.Minute
	s := d.Seconds()
	d -= time.Duration(s) * time.Second
	b := bytes.Buffer{}
	if neg {
		b.Write([]byte{'-'})
	}
	b.Write([]byte{'P'})
	if y > 0 {
		b.WriteString(fmt.Sprintf("%dY", int64(y)))
	}
	if m > 0 {
		b.WriteString(fmt.Sprintf("%dM", int64(m)))
	}
	if dd > 0 {
		b.WriteString(fmt.Sprintf("%dD", int64(dd)))
	}

	if H + M + s > 0 {
		b.Write([]byte{'T'})
		if H > 0 {
			b.WriteString(fmt.Sprintf("%dH", int64(H)))
		}
		if M > 0 {
			b.WriteString(fmt.Sprintf("%dM", int64(M)))
		}
		if s > 0 {
			b.WriteString(fmt.Sprintf("%dS", int64(s)))
		}
	}
	return b.Bytes(), nil
}
