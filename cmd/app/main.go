package main

import (
	"fmt"
	"go-demo/pkg/cycles"
	"go-demo/pkg/ifelse"
	"go-demo/pkg/print"
	"go-demo/pkg/structures"
)

func main() {

	resWeather := ifelse.WeatherDef(true, false) // задание на if-else
	fmt.Println(resWeather)
	resStudent := ifelse.StatusStudent(3)
	fmt.Println(resStudent)

	resPrintFIO := print.TraingPrint() // задание на print
	fmt.Println(resPrintFIO)

	cycles.TrainingCycles() // задание на циклы

	banan := structures.Car{Brand: "MB CLS", Year: 2017} // задание на структуры
	fmt.Println(banan)
	banan.Year = 2020
	fmt.Println(banan)
	bmw5 := structures.CarCreation("BMW", 2008)
	fmt.Println(bmw5)
	structures.UpdateCarYear(&bmw5)
	fmt.Println(bmw5)
}
