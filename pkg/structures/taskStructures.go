package structures

type Car struct {
	Brand string
	Year  int
}

func CarCreation(br string, ye int) Car {
	return Car{
		Brand: br,
		Year:  ye,
	}
}

func UpdateCarYear(c *Car) {
	c.Year++
}
