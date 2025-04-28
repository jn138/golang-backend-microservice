package time

const (
	YEAR        = "2006"
	MONTH       = "01"
	DATE        = "02"
	HOUR        = "15"
	MINUTE      = "04"
	SECOND      = "05"
	MILLISECOND = "000000"
)

const (
	DATE_LAYOUT       = YEAR + "-" + MONTH + "-" + DATE
	TIME_LAYOUT       = HOUR + ":" + MINUTE + ":" + SECOND
	TIME_LAYOUT_IN_MS = TIME_LAYOUT + "." + MILLISECOND
	DATE_TIME_LAYOUT  = DATE_LAYOUT + " " + TIME_LAYOUT
)
