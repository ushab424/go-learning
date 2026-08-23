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

func (t Team) AddPlayer(p Player) Team {
	t.Players = append(t.Players, p)
	return t
}

func (t Team) TotalTeamGoals() int {
	TotalGoals := 0
	for _, value := range t.Players {
		TotalGoals += value.Goals
	}
	return TotalGoals
}

func (t Team) BestPlayer() string {
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

func (t Team) Info() {
	fmt.Println(t.Name)
	fmt.Println(len(t.Players))
	fmt.Println(t.TotalTeamGoals())
	fmt.Println(t.BestPlayer())
}

func main() {
	players := Team{Name: "Makhachkala"}
	players = players.AddPlayer(Player{"Ivan", 5})
	players = players.AddPlayer(Player{"Oleg", 16})
	players = players.AddPlayer(Player{"Petr", 4})
	players = players.AddPlayer(Player{"Sebastian", 55})
	players = players.AddPlayer(Player{"Evgeniy", 7})
	players.Info()
}
