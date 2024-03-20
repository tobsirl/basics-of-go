package data

type integer = int
type json = map[string]string

var x integer

type distance float64

func test() {
	d := distance(42.42)
	println(d)
}