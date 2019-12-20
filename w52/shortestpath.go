package w52

import (
	"github.com/qdongxu/arts/w51"
	"math"
	"strconv"
)

type Vertex struct {
	Id    int
	Info  interface{}
	Edges []*Edge
}

type Edge struct {
	From   int
	To     int
	Weight int

	Path string
}

type Graph struct {
	Vertex []*Vertex
}

func NewVertex(id int) *Vertex {
	return &Vertex{Edges: make([]*Edge, 0), Id: id}
}

func NewGraph(vertexCount int) *Graph {
	g := &Graph{Vertex: make([]*Vertex, 0)}
	for i := 0; i < vertexCount; i++ {
		g.Vertex = append(g.Vertex, NewVertex(i))
	}
	return g
}

func (g *Graph) AddEdge(idA int, idB int, weight int) {
	e := &Edge{Weight: weight, From: idA, To: idB}
	g.Vertex[idA].Edges = append(g.Vertex[idA].Edges, e)
}

func (g *Graph) GetWeight(idA int, idB int) (int, bool) {
	for _, e := range g.GetVertex(idA).Edges {
		if e.To == idB {
			return e.Weight, true
		}
	}

	return 0, false
}

func (g *Graph) VertexCount() int {
	return len(g.Vertex)
}

func (g *Graph) GetVertex(id int) *Vertex {
	return g.Vertex[id]
}

func ShortestPath(g *Graph, idA int) []*Edge {
	s := make([]*Edge, 0)
	u := w51.NewPriorityQueue(false)

	s = append(s, &Edge{From: idA, To: idA, Weight: 0, Path: "0"})

	for i := 0; i < g.VertexCount(); i++ {
		if i == idA {
			continue
		}

		w, ok := g.GetWeight(idA, i)
		if !ok {
			u.Push(&w51.Data{Id: int32(math.MaxInt32), Value: &Edge{From: idA, To: i, Weight: math.MaxInt32}})
			continue
		}

		u.Push(&w51.Data{Id: int32(w), Value: &Edge{From: idA, To: i, Weight: w}})
	}

	updateU(g, s, u)
	for u.Size() > 0 {
		e := u.Take()
		s = addSmallest(g, s, e.Value.(*Edge))
		updateU(g, s, u)
	}

	return s
}

func addSmallest(g *Graph, s []*Edge, e *Edge) []*Edge {
	if e.Weight < math.MaxInt32 {
		s = append(s, e)
	}

	return s
}

func updateU(g *Graph, s []*Edge, u *w51.PriorityQueue) {

	us := make([]*Edge, 0)
	for u.Size() > 0 {
		us = append(us, u.Take().Value.(*Edge))
	}

	for _, e := range us {
		smallest, pos := e.Weight, e.From
		for i, sV := range s {
			w, ok := g.GetWeight(sV.To, e.To)
			if !ok {
				continue
			}

			if e.Weight > sV.Weight+w {
				smallest, pos = sV.Weight+w, i
			}
		}

		e.Weight = smallest
		e.From = pos
		e.Path = s[pos].Path + " -> " + strconv.Itoa(e.To)
		u.Push(&w51.Data{Id: int32(e.Weight), Value: e})
	}
}
