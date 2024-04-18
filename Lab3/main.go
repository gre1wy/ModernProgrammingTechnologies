package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func userInteraction() {
	var queue Queue

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Виберіть структуру для реалізації черги:")
	fmt.Println("1 - Масив (кільцева черга)")
	fmt.Println("2 - Звязний список")
	fmt.Println("0 - Для виходу")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		queue, _ = getQueue("CircularArrayQueue", os.Stdin)
	case "2":
		queue, _ = getQueue("LinkedListQueue", nil)
	case "0":
		os.Exit(0)
	default:
		fmt.Println("Неправильний вибір операції. Введіть 1-2, або 0 для того щоб вийти.")
		userInteraction()
	}

	for {
		fmt.Println("\nВиберіть операцію (напишіть тільки число):")
		fmt.Println("1 - Enqueue (Додати елемент в чергу);")
		fmt.Println("2 - Dequeue (Видалити елемент з черги);")
		fmt.Println("3 - IsEmpty (Перевірити чи черга пуста);")
		fmt.Println("4 - Display (показати чергу);")
		fmt.Println("5 - Front (показати елемент в початку черги);")
		fmt.Println("6 - Rear (показати елемент в кінці черги);")
		fmt.Println("7 - Size (показати поточний розмір черги);")
		fmt.Println("0 - Вийти.")

		operation, _ := reader.ReadString('\n')
		operation = strings.TrimSpace(operation)

		switch operation {
		case "1":
			fmt.Println("Введіть елемент для вставки в чергу:")
			fmt.Println("Для відміни вставки напишіть 'cancel'")
			itemInput, _ := reader.ReadString('\n')
			itemInput = strings.TrimSpace(itemInput)

			if itemInput == "cancel" {
				return
			}

			queue.Enqueue(itemInput)

		case "2":
			item := queue.Dequeue()
			if item == nil {
				fmt.Println("Черга і так порожня.")
			} else {
				fmt.Println("Видалений елемент:", item)
			}

		case "3":
			if queue.IsEmpty() {
				fmt.Println("Черга порожня.")
			} else {
				fmt.Println("Черга не порожня.")
			}

		case "4":
			queue.Display()

		case "5":
			fmt.Println("Елемент на початку черги:", queue.Front())

		case "6":
			fmt.Println("Елемент вкінці черги:", queue.Rear())

		case "7":
			fmt.Println("Поточний розмір черги:", queue.Size())

		case "0":
			return

		default:
			fmt.Println("Неправильний вибір операції. Введіть 1-7, або 0 для того щоб вийти")
		}
	}
}

func main() {
	userInteraction()
}
