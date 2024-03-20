package data

type integer = int
type json = map[string]string

var x integer

type distance float64
type distanceKm	float64

func toKm(miles distance) distanceKm {
	return distanceKm(1.60934 * miles)
}

func test() {
	d := distance(42.42)
	println(d)
}