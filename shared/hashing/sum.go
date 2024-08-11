package hashing

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"gopherutils/shared/convert"
	"os"
	"strings"
)

type Sum struct {
	Hash     []byte
	File     string
	Mode     uint8
	HashType string
}

func ReadSums(str string, algo string) ([]Sum, int) {
	fmt.Println("ok")
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
		}
		switch line[len(sum.Hash)+1] {
		case '^':
			sum.Mode = 1
			break
		default:
			sum.Mode = 0
			break
		}
		sums = append(sums, sum)
	}
	return sums, malformed
}
func VerifySum(sum Sum) (bool, error) {
	_, err := os.Stat(sum.File)
	if err != nil {
		return false, errors.New(fmt.Sprintf("File does not exist"))
	}
	data, err := os.ReadFile(sum.File)
	if err != nil {
		fmt.Printf("Failed to read")
		os.Exit(1)
	}
	if sum.Mode == 1 {
		data = convert.ReadAsciiBits(data)
	}
	switch sum.HashType {
	case "SHA1":
		return fmt.Sprintf("%x", sha1.Sum(data)) == fmt.Sprintf("%x", sum.Hash), nil
	case "SHA224":
		return fmt.Sprintf("%x", sha256.Sum224(data)) == fmt.Sprintf("%x", sum.Hash), nil
	case "SHA256":
		return fmt.Sprintf("%x", sha256.Sum256(data)) == fmt.Sprintf("%x", sum.Hash), nil
	case "SHA384":
		return fmt.Sprintf("%x", sha512.Sum384(data)) == fmt.Sprintf("%x", sum.Hash), nil
	case "SHA512":
		return fmt.Sprintf("%x", sha512.Sum512(data)) == fmt.Sprintf("%x", sum.Hash), nil
	case "SHA512/224":
		return fmt.Sprintf("%x", sha512.Sum512_224(data)) == fmt.Sprintf("%x", sum.Hash), nil
	case "SHA512/256":
		return fmt.Sprintf("%x", sha512.Sum512_256(data)) == fmt.Sprintf("%x", sum.Hash), nil
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
