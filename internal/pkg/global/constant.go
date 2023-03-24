package global

const (
	TimeFormat = "2006-01-02 15:04:05"
	MaxLimit   = 1000
)

const Version = "2.8.0"
const PageSizeDefault = 10 //默认每页显示数量

const LogMsg = "hss" //日志收集关键信息标识

const (
	SecondsPerMinute = 60
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerDay    = 24 * SecondsPerHour
	SecondsPerWeek   = 7 * SecondsPerDay
	DaysPer400Years  = 365*400 + 97
	DaysPer100Years  = 365*100 + 24
	DaysPer4Years    = 365*4 + 1
)
