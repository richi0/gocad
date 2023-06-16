package cad

type CommonMutations struct {
	parent Object
}

func (c *CommonMutations) Translate(x float64, y float64, z float64) Object {
	return NewTranslate(x, y, z, c.parent)
}

func (c *CommonMutations) Mirror(x float64, y float64, z float64) Object {
	return NewMirror(x, y, z, c.parent)
}

func (c *CommonMutations) Rotate(x float64, y float64, z float64) Object {
	return NewRotate(x, y, z, c.parent)
}

func (c *CommonMutations) Scale(x float64, y float64, z float64) Object {
	return NewScale(x, y, z, c.parent)
}

func (c *CommonMutations) Union(o ...Object) Object {
	p := []Object{c.parent}
	return NewUnion(append(p, o...)...)
}

func (c *CommonMutations) Difference(o ...Object) Object {
	p := []Object{c.parent}
	return NewDifference(append(p, o...)...)
}

func (c *CommonMutations) Intersection(o ...Object) Object {
	p := []Object{c.parent}
	return NewIntersection(append(p, o...)...)
}

func (c *CommonMutations) Hull(o ...Object) Object {
	p := []Object{c.parent}
	return NewHull(append(p, o...)...)
}

func (c *CommonMutations) Minkowski(o ...Object) Object {
	p := []Object{c.parent}
	return NewMinkowski(append(p, o...)...)
}
