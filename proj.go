package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	////////////// проверка на наличие файла и его создание
	file, err := os.OpenFile("params.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("unable to create file:", err)
		os.Exit(1)
	}

	/////////////// достать записи из файла и добавление тестовых параметров
	human := make(map[string]string)

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		var perv []string = strings.Split(sc.Text(), " ")
		human[perv[0]] = perv[2]
	}
	if len(human) < 1 {
		human["age"] = "25"
		human["height"] = "183"
		human["weight"] = "70"
	}

	////////////// действия в зависимости от параметров запуска
	if len(os.Args) <= 1 {
		for key, value := range human {
			var str string = key + " = " + value
			fmt.Println(str)
		}
		if err != nil {
			panic(err)
		}
		fmt.Print("Программа имеет 3 функции (add, remove, rewrite)")
	} else {
		switch os.Args[1] {
		case "add":
			var new_key string
			var new_value string
			fmt.Print("Введите ключ - ")
			fmt.Scan(&new_key)
			fmt.Print("Введите значение - ")
			fmt.Scan(&new_value)
			human[new_key] = new_value

		case "remove":
			var to_remove string
			fmt.Print("Введите ключ, который необходимо удалить - ")
			fmt.Scan(&to_remove)
			delete(human, to_remove)

		case "rewrite":

			var key_to_rewrite string
			var to_rewrite string
			fmt.Print("Введите ключ, который необходимо перезаписать - ")
			fmt.Scan(&key_to_rewrite)
			if val, ok := human[key_to_rewrite]; ok {
				fmt.Print("Текущее значение - " + val + ", введите новое - ")
				fmt.Scan(&to_rewrite)
				human[key_to_rewrite] = to_rewrite
			} else {
				fmt.Print("Такого ключа не существует.")
			}

		default:
			fmt.Print("Unknown command")
		}
	}

	os.Truncate("params.txt", 0)
	for key, value := range human {
		var str string = key + " = " + value + "\n"
		file.WriteString(str)
	}
	defer file.Close()
}
