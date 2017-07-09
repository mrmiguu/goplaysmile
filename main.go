package main

import (
	"github.com/mrmiguu/goop"
)

func main() {
	goop.New(
		450, 800,
		goop.Assets{},
		goop.States{
			"start": goop.State(
				start{},
				goop.Assets{
					"button": goop.Button("go_*.png"),
				},
			),
		},
	)
}
