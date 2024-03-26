package data

var Countries [10]string

func init() {
	Countries[0] = "USA"
	Countries[1] = "Canada"
	Countries[2] = "Mexico"
	Countries[3] = "Brazil"
	Countries[4] = "Argentina"
	Countries[5] = "Chile"
	Countries[6] = "Colombia"
	Countries[7] = "Peru"
	Countries[8] = "Ecuador"
	Countries[9] = "Venezuela"

	// get the length of an array
	qty := len(Countries)
	println(qty)

	Prices := [5]float64{10.5, 20.5, 30.5, 40.5, 50.5}
	println(Prices[0])
}
