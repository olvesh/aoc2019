package main

import (
	"fmt"
	"log"
	"strings"
)

type MassMap struct {
	COM *Mass
	idx map[string]*Mass
}

func (m *MassMap) add(s string) {
	orbit := strings.Split(s, ")")

	if _, ok := m.idx[orbit[0]]; !ok {
		// this is the senter of gravity
		m.idx[orbit[0]] = &Mass{
			Id:      orbit[0],
			InOrbit: make([]*Mass, 0),
		}

		if orbit[0] == "COM" {
			m.COM = m.idx[orbit[0]]
		}
	}
	if _, ok := m.idx[orbit[1]]; !ok {
		m.idx[orbit[1]] = &Mass{Id: orbit[1], InOrbit: make([]*Mass, 0)}
	}

	m.idx[orbit[0]].AddInOrbit(m.idx[orbit[1]])

}

type Mass struct {
	Id string

	Orbits  *Mass
	InOrbit []*Mass
}

func (m Mass) String() string {
	return fmt.Sprintf("%s inorb: %v", len(m.InOrbit))
}

type OrbitCount int

func (m *Mass) NumDirectAndIndirect() int {
	num := m.walk(0)
	return num
}

func (m *Mass) AddInOrbit(orbitee *Mass) {
	m.InOrbit = append(m.InOrbit, orbitee)
	orbitee.Orbits = m
}

func (m Mass) walk(depth int) int {
	if m.Id != "COM" {
		depth++
	}
	pad := fmt.Sprintf("%%%ds", depth)
	log.Printf(pad, m.Id)

	currentDepth := depth
	for _, inOrbit := range m.InOrbit {
		depth += inOrbit.walk(currentDepth)

	}
	return depth

}
