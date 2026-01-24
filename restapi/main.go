package main

import (
	"context"
	"fmt"
	"restapi/http"
	"restapi/sql/connection"
	"restapi/sql/database"
	"restapi/src"
	"time"
	// "github.com/k0kubun/pp"
	// "time"
)

func main() {
	ctx := context.Background()
	conn, err := connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := database.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// if err := database.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"Погулять с собакой",
	// 	"Бобиком",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	// if err := database.UpdateRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := database.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// tasks, err := database.SelectRows(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }

	// pp.Println(tasks)

	tasks, err := database.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Покормить кота"
			task.Description = "Отсыпать коту 30 грамм корма"
			task.Completed = true
			now := time.Now()
			task.Completed_at = &now

			if err := database.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}

			break
		}
	}

	fmt.Println("Succeed")

	todoList := src.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start HTTP Server: ", err)
	}
}
