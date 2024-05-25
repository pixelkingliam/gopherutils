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
func ToBinary[T constraints.Integer](value T) string {
	finalValue := uint64(value)
	switch {
	case finalValue >= (1024 * 1024 * 1024 * 1024 * 1024):
		return fmt.Sprintf("%v PiB", ToPeta(finalValue))
	case finalValue >= (1024 * 1024 * 1024 * 1024):
		return fmt.Sprintf("%v TiB", ToTera(finalValue))
	case finalValue >= (1024 * 1024 * 1024):
		return fmt.Sprintf("%v GiB", ToGiga(finalValue))
	case finalValue >= (1024 * 1024):
		return fmt.Sprintf("%v MiB", ToMega(finalValue))
	case finalValue >= 1024:
		return fmt.Sprintf("%v KiB", ToKilo(finalValue))
	default:
		return fmt.Sprintf("%v B", finalValue)
	}
}

func ToSI[T constraints.Integer](value T) string {
	finalValue := uint64(value)
	switch {
	case finalValue >= 1000000000000000:
		return fmt.Sprintf("%v PB", ToPeta(finalValue, true))
	case finalValue >= 1000000000000:
		return fmt.Sprintf("%v TB", ToTera(finalValue, true))
	case finalValue >= 1000000000:
		return fmt.Sprintf("%v GB", ToGiga(finalValue, true))
	case finalValue >= 1000000:
		return fmt.Sprintf("%v MB", ToMega(finalValue, true))
	case finalValue >= 1000:
		return fmt.Sprintf("%v KB", ToKilo(finalValue, true))
	default:
		return fmt.Sprintf("%v B", value)
	}
}
