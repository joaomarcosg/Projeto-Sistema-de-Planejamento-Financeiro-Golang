package util

import "testing"

func TestStringToTime(t *testing.T) {

	var convertedTime = StringToTime("2025-03-08T12:55:10")

	if convertedTime.Year() != 2025 {
		t.Errorf("Espera que o ano seja 2025")
	}

	if convertedTime.Month() != 03 {
		t.Errorf("Esperar que o mês seja 03")
	}

	if convertedTime.Hour() != 12 {
		t.Errorf("Espera que a hora seja 10")
	}

}
