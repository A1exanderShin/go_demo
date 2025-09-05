package print

func TraingPrint() string {
	res_fullName := fullName("Alexander", "Shin")
	return res_fullName
}

func fullName(firstName, lastName string) string {
	return firstName + " " + lastName
}
