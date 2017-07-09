package main

import "github.com/mrmiguu/goop"

type start struct {
	g *goop.Global
	a goop.Assets
}

func (s start) New(g *goop.Global, a goop.Assets) {
	s.g = g
	s.a = a
}
