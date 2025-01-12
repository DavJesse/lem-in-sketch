package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"lem-in/models"
)

func ReadFile(filename string) []models.Graph {
	var room models.Graph
	var graph []models.Graph
	var params []string
	var totalAnts int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if totalAnts == 0 {
			totalAnts, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if strings.HasPrefix(line, "#") {
				if line == "##start" {
					scanner.Scan()
					params = strings.Split(scanner.Text(), " ")
					room.Start = true
					room.Room = params[0]

					for i := range params[1:] {
						room.Connection = append(room.Connection, params[i])
					}
					graph = append(graph, room)
					room = models.Graph{}
				} else if line == "##end" {
					scanner.Scan()
					params = strings.Split(scanner.Text(), " ")
					room.End = true
					room.Room = params[0]

					for i := range params[1:] {
						room.Connection = append(room.Connection, params[i])
					}
					graph = append(graph, room)
					room = models.Graph{}
				} else {
					continue
				}
			}
			params = strings.Split(scanner.Text(), " ")
			room.Room = params[0]
			for i := range params[1:] {
				room.Connection = append(room.Connection, params[i])
			}
			graph = append(graph, room)
			room = models.Graph{}
		}

	}

	return graph
}
