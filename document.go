package cad

import (
	"fmt"
	"strings"
)

type Object interface {
	render() string
	Translate(x float64, y float64, z float64) Object
	Mirror(x float64, y float64, z float64) Object
	Rotate(x float64, y float64, z float64) Object
	Scale(x float64, y float64, z float64) Object
	Union(o ...Object) Object
	Difference(o ...Object) Object
	Hull(o ...Object) Object
	Intersection(o ...Object) Object
	Minkowski(o ...Object) Object
}

type Document struct {
	objects []Object
	Fn      int
}

func (d *Document) Render() string {
	res := fmt.Sprintf("$fn = %d;\n", d.Fn)
	indent := 0
	for _, obj := range d.objects {
		res += obj.render()
	}
	res_formatted := ""
	for i, c := range res {
		if c == '{' {
			indent++
		} else if c == '}' {
			indent--
		}
		if c == '\n' && indent > 0 {
			res_formatted += string(c)
			if res[i+1] == '}' {
				res_formatted += strings.Repeat("\t", indent-1)
			} else {
				res_formatted += strings.Repeat("\t", indent)
			}
		} else {
			res_formatted += string(c)
		}
	}
	return res_formatted
}

func NewDocument(o ...Object) *Document {
	return &Document{o, 50}
}
