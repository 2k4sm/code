package temptable

type Temperature float32
type Celsius float32
type Fahrenheit float32

var K Temperature

func (K Celsius) CelConv() Temperature {
	farhen := Temperature((K * 9 / 5) + 32)
	return farhen
}

func (K Fahrenheit) FarhenConv() Temperature {
	celc := Temperature((K - 32) * 5 / 9)
	return celc
}
