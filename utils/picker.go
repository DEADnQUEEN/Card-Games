package utils

import "fmt"

func ChooseItem[T any](items ...T) (int, T) {
	for {
		fmt.Println("Choose item:")
		for i, item := range items {
			fmt.Printf("%d\t%s\n", i, item)
		}

		var answer int
		_, err := fmt.Scanf("%d", &answer)
		if err != nil {
			continue
		}

		if answer < len(items) && answer >= 0 {
			return answer, items[answer]
		}
	}
}
