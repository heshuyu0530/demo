package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todolistinmemory/model"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := map[int]string{}
	fmt.Println("Welcome todolist!")
	for{
		fmt.Println("Please select an option")
		fmt.Println("1. add task.2. delete task.3. view task.4. exit")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		switch input {
			case "1":
				fmt.Println("Please input the task you want to add")
				scanner.Scan()
				input = scanner.Text()
				task := model.NewTask(input)
				tasks[task.Id] = task.Content
				break
			case "2":
				fmt.Println("Please select the task you want to delete")
				scanner.Scan()
				input,_ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
				delete(tasks, input)
				break
			case "3":
				break
			case "4":
				fmt.Println("Thank you for using todolist!")
				return
			default:
				fmt.Println("Invalid option, please try again")
		}
		fmt.Printf("Current tasks:")
		for id, content := range tasks{
			fmt.Printf("%d. %s\n", id, content)
		}
	}
	
}