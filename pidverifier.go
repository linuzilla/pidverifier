package pidverifier

import (
	"fmt"
	"regexp"
)

type PidVerifier struct {
}

func (svc *PidVerifier) getCode(val byte) int {
	switch val {
	case 'I':
		return 34
	case 'O':
		return 35
	case 'W':
		return 32
	case 'Z':
		return 33
	case 'X':
		return 30
	case 'Y':
		return 31
	default:
		code := int(val)
		rc := code - 'A' + 10
		if rc > 17 {
			rc--
		}
		if rc > 22 {
			rc--
		}
		return rc
	}
}

func (svc *PidVerifier) get_d_number(code byte) int {
	if code >= '0' && code <= '9' {
		return int(code - '0')
	} else {
		return svc.getCode(code)
	}
}

func (svc *PidVerifier) Verify(str string) bool {
	r, err := regexp.Compile(`^[A-Z][A-Z0-9][0-9]{8}$`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return false
	}

	if r.MatchString(str) == true {
		var list [11]int

		byteArray := []byte(str)

		code := svc.getCode(byteArray[0])
		x2 := code % 10
		list[0] = (code - x2) / 10
		list[1] = x2
		for i := 1; i <= 9; i++ {
			list[i+1] = svc.get_d_number(byteArray[i]) % 10
		}

		multiply := []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1}
		sum := 0

		for i := 0; i < 11; i++ {
			sum += list[i] * multiply[i]
		}

		return sum%10 == 0
	}

	return false
}
