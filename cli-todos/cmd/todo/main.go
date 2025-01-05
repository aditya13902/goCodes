package main

import (
	internal "adi/cli-todo/internal/repository"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "todos.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = internal.Create(db)
	if err != nil {
		panic(err)
	}
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: task <command> [arguments]")
		fmt.Println("\nAvailable commands:")
		fmt.Println("  add    Add a new task")
		fmt.Println("  do     Mark a task as done")
		fmt.Println(" list    List all the tasks to be done")
		return
	}

	switch args[1] {
	case "add":
		if len(args) == 2 {
			fmt.Println("Give an appropriate task")
		}
		task := strings.Join(args[2:], " ")
		err = internal.Add(db, task)
		if err != nil {
			panic(err)
		}
		fmt.Println("Task added:", task)
	case "do":
		if len(args) == 2 {
			fmt.Println("Give an appropriate task")
		}
		taskID, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
		res, err := internal.Delete(db, taskID)
		if err != nil {
			panic(err)
		}
		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			fmt.Println("No task found with ID:", taskID)
		} else {
			fmt.Println("Task completed and removed:", taskID)
		}
	case "list":
		rows, err := internal.List(db)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		fmt.Println("Current tasks:")
		for rows.Next() {
			var id int
			var task string
			err = rows.Scan(&id, &task)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d: %s\n", id, task)
		}
		if err = rows.Err(); err != nil {
			panic(err)
		}
	}

}
