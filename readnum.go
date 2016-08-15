package readnum

import (
	"bufio"
	"math"
	"strings"
)

const digitlist = "0123456789"

func read_uint(r *bufio.Reader) (int, int, bool, error) {
	ok := false
	value := 0
	ten := 1
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			return value, ten, ok, err
		}
		n := strings.IndexRune(digitlist, ch)
		if n < 0 {
			r.UnreadRune()
			return value, ten, ok, nil
		}
		value = value*10 + n
		ten *= 10
		ok = true
	}
}

func read_int(r *bufio.Reader) (int, bool, bool, error) {
	value := 0
	ch, _, err := r.ReadRune()
	if err != nil {
		return 0, false, false, err
	}
	var minus bool = false
	if ch == '-' {
		minus = true
	} else if ch != '+' {
		r.UnreadRune()
	}
	var ok bool
	value, _, ok, err = read_uint(r)
	return value, minus, ok, err
}

func Float(r *bufio.Reader) (float64, bool, error) {
	intval, minus, ok, err := read_int(r)
	if err != nil {
		if minus {
			intval = -intval
		}
		return float64(intval), ok, err
	}
	var ch rune
	ch, _, err = r.ReadRune()
	if err != nil {
		if minus {
			intval = -intval
		}
		return float64(intval), ok, err
	}
	var base = float64(intval)
	if ch == '.' {
		var ten int
		var lower int
		lower, ten, ok, err = read_uint(r)
		if ok {
			base += float64(lower) / float64(ten)
		}
	}
	if minus {
		base = -base
	}
	ch, _, err = r.ReadRune()
	if err != nil {
		return base, true, err
	}
	if ch != 'e' && ch != 'E' {
		r.UnreadRune()
		return base, true, nil
	}
	var ex int
	ex, minus, ok, err = read_int(r)
	if minus {
		ex = -ex
	}
	return base * math.Pow10(+ex), true, err
}
