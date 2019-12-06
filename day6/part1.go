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

func (m MassMap) distance(id1, id2 string) int {
	mass1, mass2 := m.idx[id1], m.idx[id2]
	return mass1.distance(mass2)
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

func (m *Mass) pathToCom() (path []string) {

	for mass := m.Orbits; mass != nil; mass = mass.Orbits {
		path = append(path, mass.Id)
	}

	return
}

func (m Mass) distance(mass2 *Mass) int {
	path1, path2 := m.pathToCom(), mass2.pathToCom()

	for i, s := range path1 {
		for i2, s2 := range path2 {
			if s == s2 {
				return i + i2
			}
		}
	}
	return -1
}
