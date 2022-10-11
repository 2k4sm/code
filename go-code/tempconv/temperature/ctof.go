package temperature

func Ctof(temp int) int {
	f := (temp * 9.0 / 5.0) + 32.0
	return f
}
func Ktof(temp int) int {
	c := temp - 273

	f := (c * 9.0 / 5.0) + 32.0

	return f
}
