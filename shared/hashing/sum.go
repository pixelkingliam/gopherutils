package hashing

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"gopherutils/shared/convert"
	"os"
	"strconv"
	"strings"
)

type Sum struct {
	Hash     []byte
	File     string
	Mode     uint8
	HashType string
}

func (receiver Sum) String() string {
	indicator := " "

	if receiver.Mode&Binary != 0 {
		indicator = "*"
	} else if receiver.Mode&BitMode != 0 {
		indicator = "^"
	} else if receiver.Mode&Universal != 0 {
		indicator = "U"
	}

	if receiver.Mode&Tag == 1 {
		return fmt.Sprintf("%s (%s) = %x", receiver.HashType, receiver.File, receiver.Hash)
	} else {
		return fmt.Sprintf("%x %s%s", receiver.Hash, indicator, receiver.File)
	}
}

type SumTemplate struct {
	File     string
	Mode     uint8
	HashType string
}

func ReadSums(str string, algo string) ([]Sum, int) {
	sums := make([]Sum, 0)
	malformed := 0
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		var sum Sum
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		sum.Hash = []byte(strings.Split(line, " ")[0])
		if len(line) < len(sum.Hash)+3 || !strings.Contains(line, " ") || len(strings.Split(line, " ")[0]) != len(sum.Hash) {
			fmt.Printf("Line %v is improperly formatted.\n", i+1)
			malformed++
		}
		sum.File = line[len(sum.Hash)+2:]

		switch algo {
		case "sha":
			sum.HashType = shaAlgoString(shaAlgoFromLength(len(sum.Hash)))
		case "b2":
			sum.HashType = b2AlgoString(b2AlgoFromLength(len(sum.Hash)))
		}
		switch line[len(sum.Hash)+1] {
		case '^':
			sum.Mode |= BitMode
			break
		default:
			break
		}
		sums = append(sums, sum)
	}
	return sums, malformed
}
func GetSum(template SumTemplate) Sum {
	data, err := os.ReadFile(template.File)
	if err != nil {
		fmt.Printf("Error reading file %v: %v\n", template.File, err)
	}
	if template.Mode&BitMode == 1 && template.Mode&Universal == 1 {
		fmt.Printf("Ambiguous file mode, ignoring --UNIVERSAL")
	}
	if template.Mode&BitMode == 1 {
		data = convert.ReadAsciiBits(data)
	} else if template.Mode&Universal == 1 {
		data = bytes.Replace(data, []byte{'\r', '\n'}, []byte{'\n'}, -1)
		data = bytes.Replace(data, []byte{'\r'}, []byte{'\n'}, -1)
	}
	var sum Sum
	sum.File = template.File
	sum.Mode = template.Mode
	sum.HashType = template.HashType
	sum.Hash = Hash(data, sum.HashType)
	return sum
}
func Hash(data []byte, algorithm string) []byte {
	if strings.Contains(algorithm, "BLAKE2b-") {
		length, _ := strconv.ParseInt(algorithm[8:], 10, 32)
		hash, _ := blake2b.New(int(length)/8, nil)

		hash.Write(data)
		return hash.Sum(nil)
	}
	switch algorithm {
	case "SHA1":
		hash := sha1.Sum(data)
		return hash[:]
	case "SHA224":
		hash := sha256.Sum224(data)
		return hash[:]
	case "SHA256":
		hash := sha256.Sum256(data)
		return hash[:]
	case "SHA384":
		hash := sha512.Sum384(data)
		return hash[:]
	case "SHA512":
		hash := sha512.Sum512(data)
		return hash[:]
	case "SHA512/224":
		hash := sha512.Sum512_224(data)
		return hash[:]
	case "SHA512/256":
		hash := sha512.Sum512_256(data)
		return hash[:]
	default:
		return []byte("err")
	}
}
func VerifySum(sum Sum) (bool, error) {
	_, err := os.Stat(sum.File)
	if err != nil {
		return false, errors.New("file does not exist")
	}
	data, err := os.ReadFile(sum.File)
	if err != nil {
		fmt.Printf("Failed to read")
		os.Exit(1)
	}
	if sum.Mode&Universal == 1 {
		data = convert.ReadAsciiBits(data)
	}
	if sum.Mode&Universal == 1 {
		data = bytes.Replace(data, []byte{'\r', '\n'}, []byte{'\n'}, -1)
		data = bytes.Replace(data, []byte{'\r'}, []byte{'\n'}, -1)
	}
	if strings.Contains(sum.HashType, "BLAKE2b-") {

		processed := Hash(data, sum.HashType)
		return fmt.Sprintf("%x", processed) == fmt.Sprintf("%s", sum.Hash), nil
	}
	switch sum.HashType {
	case "SHA1":
		return fmt.Sprintf("%x", sha1.Sum(data)) == fmt.Sprintf("%s", sum.Hash), nil
	case "SHA224":
		return fmt.Sprintf("%x", sha256.Sum224(data)) == fmt.Sprintf("%s", sum.Hash), nil
	case "SHA256":
		return fmt.Sprintf("%x", sha256.Sum256(data)) == fmt.Sprintf("%s", sum.Hash), nil
	case "SHA384":
		return fmt.Sprintf("%x", sha512.Sum384(data)) == fmt.Sprintf("%s", sum.Hash), nil
	case "SHA512":
		return fmt.Sprintf("%x", sha512.Sum512(data)) == fmt.Sprintf("%s", sum.Hash), nil
	case "SHA512/224":
		return fmt.Sprintf("%x", sha512.Sum512_224(data)) == fmt.Sprintf("%s", sum.Hash), nil
	case "SHA512/256":
		return fmt.Sprintf("%x", sha512.Sum512_256(data)) == fmt.Sprintf("%s", sum.Hash), nil
	default:
		return false, errors.New("invalid sum algorithm")

	}
}
func shaAlgoFromLength(length int) int {
	switch length {
	case 40:
		return 1
	case 56:
		return 224 // or 512224
	case 64:
		return 256 // or 512256
	case 96:
		return 384
	case 128:
		return 512
	default:
		return -1
	}
}
func shaAlgoString(algorithm int) string {
	switch algorithm {
	case 1:
		return "SHA1"
	case 224:
		return "SHA224"
	case 256:
		return "SHA256"
	case 384:
		return "SHA384"
	case 512:
		return "SHA512"
	case 512224:
		return "SHA512/224"
	case 512256:
		return "SHA512/256"
	default:
		return "UNKNOWN"

	}
}
func b2AlgoFromLength(length int) int {
	return length * 4
}
func b2AlgoString(algorithm int) string {
	return fmt.Sprintf("BLAKE2b-%d", algorithm)
}
