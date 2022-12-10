package temptable

type Temperature float32

var K float32

func CelConv() Temperature {
	farhen := Temperature((K * 9 / 5) + 32)
	return farhen
}

func FarhenConv() Temperature {
	celc := Temperature((K - 32) * 5 / 9)
	return celc
}
