package cad

import (
	"fmt"
)

type Cube struct {
	x float64
	y float64
	z float64
	CommonMutations
}

func (c *Cube) render() string {
	return fmt.Sprintf("cube([%f, %f, %f]);\n", c.x, c.y, c.z)
}

func (c *Cube) Center() Object {
	return NewTranslate(-c.x/2, -c.y/2, -c.z/2, c)
}

func NewCube(x float64, y float64, z float64) *Cube {
	c := &Cube{x, y, z, CommonMutations{}}
	c.parent = c
	return c
}

type Cylinder struct {
	h float64
	r float64
	CommonMutations
}

func (c *Cylinder) render() string {
	return fmt.Sprintf("cylinder(%f, %f, %f);\n", c.h, c.r, c.r)
}

func (c *Cylinder) Center() Object {
	return NewTranslate(0, 0, -c.h/2, c)
}

func NewCylinder(h float64, r float64) *Cylinder {
	c := &Cylinder{h, r, CommonMutations{}}
	c.parent = c
	return c
}

type Sphere struct {
	r float64
	CommonMutations
}

func (s *Sphere) render() string {
	return fmt.Sprintf("sphere(%f);\n", s.r)
}

func NewSphere(r float64) *Sphere {
	s := &Sphere{r, CommonMutations{}}
	s.parent = s
	return s
}

type RoundedCube struct {
	x      float64
	y      float64
	z      float64
	radius float64
	CommonMutations
}

func (r *RoundedCube) render() string {
	radius := 0.01
	if r.radius != 0 {
		radius = r.radius
	}

	s1 := NewTranslate(radius, radius, radius, NewSphere(radius))
	s2 := NewTranslate(r.x-2*radius, 0, 0, s1)
	s3 := NewTranslate(r.x-2*radius, r.y-2*radius, 0, s1)
	s4 := NewTranslate(0, r.y-2*radius, 0, s1)
	hull := NewHull(s1, s2, s3, s4, NewTranslate(0, 0, r.z-2*radius, s1, s2, s3, s4))
	return hull.render()
}

func (r *RoundedCube) Center() Object {
	return NewTranslate(-r.x/2, -r.y/2, -r.z/2, r)
}

func NewRoundedCube(
	x float64,
	y float64,
	z float64,
	radius float64,
) *RoundedCube {
	c := &RoundedCube{x, y, z, radius, CommonMutations{}}
	c.parent = c
	return c
}

type Box struct {
	x      float64
	y      float64
	z      float64
	radius float64
	wall   float64
	CommonMutations
}

func (b *Box) render() string {
	base := NewCube(b.x, b.y, b.z)
	radius := 0.01
	if b.radius != 0 {
		radius = b.radius
	}

	diff := NewRoundedCube(b.x-2*b.wall, b.y-2*b.wall, b.z+radius, radius).Translate(b.wall, b.wall, b.wall)
	res := base.Difference(diff)
	return res.render()
}

func (b *Box) Center() Object {
	return NewTranslate(-b.x/2, -b.y/2, -b.z/2, b)
}

func NewBox(
	x float64,
	y float64,
	z float64,
	radius float64,
	wall float64) *Box {
	b := &Box{x, y, z, radius, wall, CommonMutations{}}
	b.parent = b
	return b
}
