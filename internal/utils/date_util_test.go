package utils_test

import (
	"testing"

	"example.com/go-auth-globo/internal/utils"
	mock_util "example.com/go-auth-globo/mock/util"
)

func TestFormatDateToISO(t *testing.T) {
	date := mock_util.Date
	formated := utils.FormatDateToISO(date)
	t.Run("it should return expected date", func(t *testing.T) {
		expected := mock_util.ExpectedISO

		if formated != expected {
			t.Errorf("expected '%s', received '%s'", expected, formated)
		}
	})

	t.Run("it should return a invalid string date", func(t *testing.T) {
		invalidExpected := mock_util.InvalidExpected

		if formated == invalidExpected {
			t.Errorf("expected '%s', received '%s'", invalidExpected, formated)
		}
	})
}

func TestFormatDateToDDMMYYY(t *testing.T) {
	date := mock_util.Date
	formated := utils.FormatDateToDDMMYYYY(date)
	t.Run("it should return expected date DD/MM/YYYY", func(t *testing.T) {
		expected := mock_util.ExpectedDDMMYYYY

		if formated != expected {
			t.Errorf("expected '%s', received '%s'", expected, formated)
		}
	})

	t.Run("it should return a invalid string date DD/MM/YYYY", func(t *testing.T) {
		invalidExpected := mock_util.InvalidExpected

		if formated == invalidExpected {
			t.Errorf("expected '%s', received '%s'", invalidExpected, formated)
		}
	})
}
