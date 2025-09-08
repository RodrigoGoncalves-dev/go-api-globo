package mock_util

import "time"

var (
	Date             = time.Date(2025, time.September, 8, 14, 30, 0, 0, time.UTC)
	ExpectedISO      = "2025-09-08T14:30:00.000Z"
	ExpectedDDMMYYYY = "08/09/2025"
	InvalidExpected  = "2025-09-08T14:30:00"
)
