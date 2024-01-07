package itrations

func Repeat(times int, char string) string {
	str := ""
	for i := 0; i<times; i++ {
		str += char 
	}
	return str
}