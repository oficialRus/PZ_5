package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"cadet_project/store"
)

func main() {
	mem := store.NewMemoryStore()
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Cadet CLI. Команды:\n  list | add <name> <rank> | del <id> | upd <id> <name> <rank> | exit")

	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}
		cmd := strings.Fields(reader.Text())
		if len(cmd) == 0 {
			continue
		}
		switch cmd[0] {

		case "list":
			for _, c := range mem.GetAll() {
				fmt.Println(c)
			}

		case "add":
			if len(cmd) < 3 {
				fmt.Println("usage: add <name> <rank>")
				continue
			}
			c := mem.Add(cmd[1], cmd[2])
			fmt.Println("added:", c)

		case "del":
			if len(cmd) < 2 {
				fmt.Println("usage: del <id>")
				continue
			}
			id, _ := strconv.Atoi(cmd[1])
			if err := mem.Delete(id); err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Println("deleted", id)
			}

		case "upd":
			if len(cmd) < 4 {
				fmt.Println("usage: upd <id> <name> <rank>")
				continue
			}
			id, _ := strconv.Atoi(cmd[1])
			if err := mem.Update(id, cmd[2], cmd[3]); err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Println("updated", id)
			}

		case "exit":
			fmt.Println("bye!")
			return

		default:
			fmt.Println("unknown command")
		}
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
}
