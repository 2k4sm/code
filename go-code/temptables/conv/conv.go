package temptable

type fahrenheit float32
type celsius float32

func CelConv(k float32) fahrenheit {
	farhen := fahrenheit((k * 9 / 5) + 32)
	return farhen
}

func FarhenConv(k float32) celsius {
	celc := celsius((k - 32) * 5 / 9)
	return celc
}
