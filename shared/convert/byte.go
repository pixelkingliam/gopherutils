package convert

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func ToKilo[T constraints.Integer](value T, options ...bool) float32 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float32(value) / 1000
	} else {
		return float32(value) / 1024
	}
}

func ToMega[T constraints.Integer](value T, options ...bool) float32 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float32(value) / 1000000
	} else {
		return float32(value) / (1024 * 1024)
	}
}

func ToGiga[T constraints.Integer](value T, options ...bool) float32 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float32(value) / 1000000000
	} else {
		return float32(value) / (1024 * 1024 * 1024)
	}
}

func ToTera[T constraints.Integer](value T, options ...bool) float32 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float32(value) / 1000000000000
	} else {
		return float32(value) / (1024 * 1024 * 1024 * 1024)
	}
}

func ToPeta[T constraints.Integer](value T, options ...bool) float32 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float32(value) / 1000000000000000
	} else {
		return float32(value) / (1024 * 1024 * 1024 * 1024 * 1024)
	}
}
func ToBinary(value int64) string {
	switch {
	case value >= (1024 * 1024 * 1024 * 1024 * 1024):
		return fmt.Sprintf("%v PiB", ToPeta(value))
	case value >= (1024 * 1024 * 1024 * 1024):
		return fmt.Sprintf("%v TiB", ToTera(value))
	case value >= (1024 * 1024 * 1024):
		return fmt.Sprintf("%v GiB", ToGiga(value))
	case value >= (1024 * 1024):
		return fmt.Sprintf("%v MiB", ToMega(value))
	case value >= 1024:
		return fmt.Sprintf("%v KiB", ToKilo(value))
	default:
		return fmt.Sprintf("%v B", value)
	}
}

func ToSI(value int64) string {
	switch {
	case value >= 1000000000000000:
		return fmt.Sprintf("%v PB", ToPeta(value, true))
	case value >= 1000000000000:
		return fmt.Sprintf("%v TB", ToTera(value, true))
	case value >= 1000000000:
		return fmt.Sprintf("%v GB", ToGiga(value, true))
	case value >= 1000000:
		return fmt.Sprintf("%v MB", ToMega(value, true))
	case value >= 1000:
		return fmt.Sprintf("%v KB", ToKilo(value, true))
	default:
		return fmt.Sprintf("%v B", value)
	}
}
