package cad

import "fmt"

type Transform struct {
	x            float64
	y            float64
	z            float64
	functionName string
	objects      []Object
	CommonMutations
}

func newTransform(x float64, y float64, z float64, functionName string, objects ...Object) *Transform {
	t := &Transform{x, y, z, functionName, objects, CommonMutations{}}
	t.parent = t
	return t
}

func (t *Transform) render() string {
	o := ""
	for _, obj := range t.objects {
		o += obj.render()
	}
	return fmt.Sprintf("%s([%f,%f,%f]) {\n%s}\n", t.functionName, t.x, t.y, t.z, o)
}

type Translate struct {
	Transform
}

func NewTranslate(x float64, y float64, z float64, o ...Object) *Translate {
	return &Translate{*newTransform(x, y, z, "translate", o...)}
}

type Mirror struct {
	Transform
}

func NewMirror(x float64, y float64, z float64, o ...Object) *Mirror {
	return &Mirror{*newTransform(x, y, z, "mirror", o...)}
}

type Rotate struct {
	Transform
}

func NewRotate(x float64, y float64, z float64, o ...Object) *Rotate {
	return &Rotate{*newTransform(x, y, z, "rotate", o...)}
}

type Scale struct {
	Transform
}

func NewScale(x float64, y float64, z float64, o ...Object) *Scale {
	return &Scale{*newTransform(x, y, z, "scale", o...)}
}
