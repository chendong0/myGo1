package test

import "fmt"


func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		return 0, fmt.Errorf("身高数值不能是0或者负数")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("体重不能为负数")
	}

	return weight / (tall * tall), nil
}
