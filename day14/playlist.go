package main

import "fmt"

type Playlist struct {
	Title    string
	Duration int
}

type PlayList struct {
	Songs []Playlist
}

func (p PlayList) Add(s Playlist) PlayList {
	p.Songs = append(p.Songs, s)
	return p
}

func (p PlayList) TotalDuration() int {
	total := 0
	for _, totalduration := range p.Songs {
		total += totalduration.Duration
	}
	return total
}

func (p PlayList) Longest() string {
	var MaxSong string
	MaxDuration := 0
	for _, song := range p.Songs {
		if song.Duration > MaxDuration {
			MaxDuration = song.Duration
			MaxSong = song.Title
		}
	}
	return MaxSong
}

func main() {
	pl := PlayList{}
	pl = pl.Add(Playlist{Title: "Arbuz", Duration: 155})
	pl = pl.Add(Playlist{Title: "banana", Duration: 394})
	pl = pl.Add(Playlist{Title: "cherry", Duration: 333})

	fmt.Println(pl.TotalDuration())
	fmt.Println(pl.Longest())
}
