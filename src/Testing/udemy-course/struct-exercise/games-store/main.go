package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const data = `

[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 30
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}
type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

func main() {

	// games := []game{

	// 	{item: item{1, "god of war", 50},
	// 		genre: "action adventure"},

	// 	{item: item{2, "x-com 2", 30},
	// 		genre: "strategy"},
	// 	{item: item{3, "minecraft", 20},
	// 		genre: "sandbox"},
	// }

	var decoded []jsonGame

	if err := json.Unmarshal([]byte(data), &decoded); err != nil {

		fmt.Println("Error")
		return
	}

	var games []game

	for _, dg := range decoded {

		games = append(games, game{item{dg.ID, dg.Name, dg.Price}, dg.Genre})
	}

	//fmt.Println(games)

	// for _, g := range games {

	// 	fmt.Printf("%d %s  %d Genre: %s\n", g.id, g.name, g.price, g.genre)
	// }

	m := make(map[int]game)

	for _, g := range games {

		m[g.id] = g
	}

	fmt.Println("Map", m)
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
		> list : list all the games
		> id : search games by id
		> save : exports the data to json and quit
		> quit: quit from the game

		`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())

		if len(cmd) == 0 {

			continue
		}

		switch cmd[0] {

		case "quit":

			fmt.Println("Bye!")
			return

		case "list":

			for _, g := range games {

				fmt.Printf("%d %s  %d Genre: %s\n", g.id, g.name, g.price, g.genre)

			}

		case "id":

			if len(cmd) != 2 {

				fmt.Println("Worng ID")
				continue
			}

			id, err := strconv.Atoi(cmd[1])

			if err != nil {

				fmt.Println("wrong ID")
				continue
			}

			g, ok := m[id]

			if !ok {

				fmt.Println("Sorry, I dont have the game")
				continue
			}

			fmt.Printf("%d %s  %d Genre: %s\n", g.id, g.name, g.price, g.genre)

		case "save":

			var enocodable []jsonGame

			for _, g := range games {

				enocodable = append(enocodable, jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(enocodable, "", "\t")

			if err != nil {

				fmt.Println("Error")
				continue
			}

			fmt.Println(string(out))
			break

		}

	}

}
