package main

import "fmt"

type Player struct {
	Name  string
	Goals int
}

type Team struct {
	Name    string
	Players []Player
}

func (t *Team) AddPlayer(p Player) {
	t.Players = append(t.Players, p)
}

func (t *Team) TotalTeamGoals() int {
	TotalGoals := 0
	for _, value := range t.Players {
		TotalGoals += value.Goals
	}
	return TotalGoals
}

func (t *Team) BestPlayer() string {
	var BestName string
	MaxGoals := 0
	for _, value := range t.Players {
		if value.Goals > MaxGoals {
			MaxGoals = value.Goals
			BestName = value.Name
		}
	}
	return BestName
}

func main() {
	players := Team{Name: "Makhachkala"}
	players.AddPlayer(Player{"Ivan", 5})
	players.AddPlayer(Player{"Oleg", 16})
	players.AddPlayer(Player{"Petr", 4})
	players.AddPlayer(Player{"Sebastian", 55})
	players.AddPlayer(Player{"Evgeniy", 7})
	fmt.Println(players.TotalTeamGoals())
	fmt.Println(players.BestPlayer())
}
