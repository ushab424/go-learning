package converter

import "errors"

func RubToDol(sum float64, rate float64) (float64, error) {
	if sum <= 0 {
		return 0, errors.New("negative amount")
	} else {
		return sum / rate, nil
	}
}

func DolToRub(sum float64, rate float64) (float64, error) {
	if sum <= 0 {
		return 0, errors.New("negative amount")
	} else {
		return sum * rate, nil
	}
}
