package cycles

import (
	"fmt"
	"go-demo/pkg/ifelse"
)

func TrainingCycles() {

	for i := 2; i < 6; i++ {
		res_statusStudent := ifelse.StatusStudent(i)
		// fmt.Println("Ну что сказать, получая оценки чисто", i, "он будет иметь статус", res_statusStudent)
		fmt.Printf("Ну что сказать, получая оценки чисто на %d, он будет иметь статус '%s'\n", i, res_statusStudent)
	}

	numbers := []int{1, 2, 3, 4, 5}

	for j, value := range numbers {
		if j == 3 || j == 4 {
			continue
		} else if j == 2 {
			break
		}
		fmt.Println(j, value)
	}

	k := 0
	for {
		fmt.Printf("%d kyyy\n", k)
		k++
		if k == 10 {
			break
		}
	}
}
