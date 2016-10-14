package text

import "strings"

// TimeOfDay converts the minutes into a day, to text in English speech.
// Multiply hours by 60.
func TimeOfDay(minutes int64) string {
	if minutes==0{return Midnight}
	m:=minutes%1440
	if m==720{return Noon}
	accumulator := make([]string, 0, 5)
	if m/60<12{
		number(m/60, &accumulator)
		if m%60<10 {accumulator = append(accumulator, TimeZero)}
		number(m%60, &accumulator)	
		if TimeBeforenoon!="" {accumulator = append(accumulator, TimeBeforenoon)}
	}else{
		number(m/60-12, &accumulator)
		if minutes%60==0{
			if TimeHour!="" {accumulator = append(accumulator, TimeHour)}
		}else{
			if m%60<10 {accumulator = append(accumulator, TimeZero)}
			number(m%60, &accumulator)
			if TimeAfternoon!="" {accumulator = append(accumulator, TimeAfternoon)}
		}
	}
	return strings.Join(accumulator, Joiner)
}
