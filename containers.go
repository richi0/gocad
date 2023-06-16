package cad

import "fmt"

type Container struct {
	functionName string
	objects      []Object
	CommonMutations
}

func newContainer(functionName string, objects ...Object) *Container {
	c := &Container{functionName, objects, CommonMutations{}}
	c.parent = c
	return c
}

func (c *Container) render() string {
	o := ""
	for _, obj := range c.objects {
		o += obj.render()
	}
	return fmt.Sprintf("%s() {\n%s}\n", c.functionName, o)
}

type Union struct {
	Container
}

func NewUnion(o ...Object) *Union {
	return &Union{*newContainer("union", o...)}
}

type Difference struct {
	Container
}

func NewDifference(o ...Object) *Difference {
	return &Difference{*newContainer("difference", o...)}
}

type Intersection struct {
	Container
}

func NewIntersection(o ...Object) *Intersection {
	return &Intersection{*newContainer("intersection", o...)}
}

type Hull struct {
	Container
}

func NewHull(o ...Object) *Hull {
	return &Hull{*newContainer("hull", o...)}
}

type Minkowski struct {
	Container
}

func NewMinkowski(o ...Object) *Minkowski {
	return &Minkowski{*newContainer("minkowski", o...)}
}
