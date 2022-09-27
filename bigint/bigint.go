package bigint

import (
	"errors"
	"strconv"
	"strings"
)

// Bigint ...
type Bigint struct {
	Value string
}

var (
	// ErrorNotNumber used when input consists of not only numbers
	ErrorNotNumber = errors.New("input not a number")
)

func cleanUp(num string) string {
	c := 0
	if strings.HasPrefix(num, "-") {
		c = 1
		num = num[1:]
	} else if strings.HasPrefix(num, "+") {
		num = num[1:]
	}
	for strings.HasPrefix(num, "0") &&
		len(num) > 1 {
		num = num[1:]
	}
	if c == 1 && num != "0" {
		num = "-" + num
	}
	return num
}

func validation(num string) (bool, string) {
	allowed := "1234567890"
	var err bool
	start := 0
	if strings.HasPrefix(num, "+") || strings.HasPrefix(num, "-") {
		start = 1
	}
	arr := strings.Split(num[start:], "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			err = true
		}
	}
	return err, num
}

// NewInt returns new Bigint value based on input
func NewInt(num string) (Bigint, error) {
	err, num := validation(num)
	if err {
		return Bigint{Value: num},

			ErrorNotNumber
	} else {
		num = cleanUp(num)
		return Bigint{Value: num}, nil
	}
}

// Set methods updates value
func (z *Bigint) Set(num string) error {
	err, num := validation(num)
	if err {
		return ErrorNotNumber
	}
	num = cleanUp(num)
	z.Value = num
	return nil
}

// Add ...
func Add(a, b Bigint) Bigint {

	num1 := strings.Split(a.Value, "")
	num2 := strings.Split(b.Value, "")
	rest := 0
	result := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1 := 0
		n2 := 0
		if i >= 0 {
			n1, _ = strconv.Atoi(num1[i])
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(num2[j])
		}
		summa := n1 + n2 + rest
		rest = summa / 10
		result = strconv.Itoa(summa%10) + result

	}
	if rest > 0 {
		result = strconv.Itoa(rest) + result

	}

	return Bigint{Value: result}
}

// Sub ...
func Sub(a, b Bigint) Bigint {
	resp := Bigint{}
	num1 := strings.Split(a.Value, "")
	num2 := strings.Split(b.Value, "")
	isNeed := false
	n1 := 0
	n2 := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1 = 0
		n2 = 0
		if i >= 0 {
			n1, _ = strconv.Atoi(num1[i])
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(num2[j])
		}
		if isNeed {
			n1 = n1 - 1
			isNeed = false
		}
		if n1-n2 >= 0 {
			isNeed = false
			resp.Value = strconv.Itoa(n1-n2) + resp.Value

		} else {
			resp.Value = strconv.Itoa(n1-n2+10) + resp.Value
			isNeed = true
		}
	}
	return resp
}

// Multiply ...
func Multiply(a, b Bigint) Bigint {
	string1 := a.Value
	string2 := b.Value
	num1, _ := strconv.Atoi(string1)
	num2, _ := strconv.Atoi(string2)
	result := num1 * num2
	return Bigint{Value: strconv.Itoa(result)}
}

// Mod ...

func Mod(a, b Bigint) Bigint {
	string1 := a.Value
	string2 := b.Value
	num1, _ := strconv.Atoi(string1)
	num2, _ := strconv.Atoi(string2)
	result := num1 / num2
	return Bigint{Value: strconv.Itoa(result)}
}

// Abs ...
func (x *Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	return Bigint{
		Value: x.Value,
	}
}
