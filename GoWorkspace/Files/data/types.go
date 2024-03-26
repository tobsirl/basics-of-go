package data

type integer = int
type json = map[string]string

var x integer

type distance float64
type distanceKm float64

// Method to convert distance from miles to km
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

// Method to convert distance from km to miles
func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(42.42)
	km := d.ToKm()
	km.ToMiles()
	println(km)
}
