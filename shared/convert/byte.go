package convert

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

func ToKilo[T constraints.Integer](value T, options ...bool) float64 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float64(value) / 1000
	} else {
		return float64(value) / 1024
	}
}

func ToMega[T constraints.Integer](value T, options ...bool) float64 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float64(value) / 1000000
	} else {
		return float64(value) / (1024 * 1024)
	}
}

func ToGiga[T constraints.Integer](value T, options ...bool) float64 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float64(value) / 1000000000
	} else {
		return float64(value) / (1024 * 1024 * 1024)
	}
}

func ToTera[T constraints.Integer](value T, options ...bool) float64 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float64(value) / 1000000000000
	} else {
		return float64(value) / (1024 * 1024 * 1024 * 1024)
	}
}

func ToPeta[T constraints.Integer](value T, options ...bool) float64 {
	si := false // Default value
	if len(options) > 0 {
		si = options[0]
	}
	if si {
		return float64(value) / 1000000000000000
	} else {
		return float64(value) / (1024 * 1024 * 1024 * 1024 * 1024)
	}
}
func ToBinary[T constraints.Integer](value T, round bool) string {
	finalValue := uint64(value)
	switch {
	case finalValue >= (1024 * 1024 * 1024 * 1024 * 1024):
		result := ToPeta(finalValue)
		if round {
			return fmt.Sprintf("%v PiB", math.Round(result))
		}
		return fmt.Sprintf("%v PiB", result)
	case finalValue >= (1024 * 1024 * 1024 * 1024):
		result := ToTera(finalValue)
		if round {
			return fmt.Sprintf("%v TiB", math.Round(result))
		}
		return fmt.Sprintf("%v TiB", result)
	case finalValue >= (1024 * 1024 * 1024):
		result := ToGiga(finalValue)
		if round {
			return fmt.Sprintf("%v GiB", math.Round(result))
		}
		return fmt.Sprintf("%v GiB", result)
	case finalValue >= (1024 * 1024):
		result := ToMega(finalValue)
		if round {
			return fmt.Sprintf("%v MiB", math.Round(result))
		}
		return fmt.Sprintf("%v MiB", result)
	case finalValue >= 1024:
		result := ToKilo(finalValue)
		if round {
			return fmt.Sprintf("%v KiB", math.Round(result))
		}
		return fmt.Sprintf("%v KiB", result)
	default:
		return fmt.Sprintf("%v B", finalValue)
	}
}

func ToSI[T constraints.Integer](value T, round bool) string {
	finalValue := uint64(value)

	switch {
	case finalValue >= 1000000000000000:
		result := ToPeta(finalValue, true)
		if round {
			return fmt.Sprintf("%v PB", math.Round(result))
		}
		return fmt.Sprintf("%v PB", result)
	case finalValue >= 1000000000000:
		result := ToTera(finalValue, true)
		if round {
			return fmt.Sprintf("%v TB", math.Round(result))
		}
		return fmt.Sprintf("%v TB", result)
	case finalValue >= 1000000000:
		result := ToGiga(finalValue, true)
		if round {
			return fmt.Sprintf("%v GB", math.Round(result))
		}
		return fmt.Sprintf("%v GB", result)
	case finalValue >= 1000000:
		result := ToMega(finalValue, true)
		if round {
			return fmt.Sprintf("%v MB", math.Round(result))
		}
		return fmt.Sprintf("%v MB", result)
	case finalValue >= 1000:
		result := ToKilo(finalValue, true)
		if round {
			return fmt.Sprintf("%v KB", math.Round(result))
		}
		return fmt.Sprintf("%v KB", result)
	default:
		return fmt.Sprintf("%v B", value)
	}
}
func ReadAsciiBits(data []byte) []byte {
	var length = 0
	for _, b := range data {
		if b != '1' && b != '0' {
			continue
		}
		length++
	}
	var outBytes = make([]byte, (length+7)/8)

	var tByte = uint8(0)
	var iByte = 0
	var iBit = 0
	for _, b := range data {
		if b != '1' && b != '0' {
			continue
		}
		if b == '1' {
			tByte |= 1 << (7 - iBit)
		}
		iBit++

		if iBit == 8 {
			iBit = 0
			outBytes[iByte] = tByte
			tByte = uint8(0)
			iByte++
		}
	}
	if iBit != 0 {
		outBytes[iByte] = tByte
	}
	for _, b := range outBytes {
		fmt.Printf("%08b\n", b)
	}
	return outBytes
}
