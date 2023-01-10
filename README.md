# GOCAD

A simple library for writing [OpenSCAD](https://openscad.org/about.html) code in go.

## Supported features

Shapes

- sphere
- cube
- cylinder

Transformations

- translate
- rotate
- mirror
- scale

Boolean operations

- union
- difference
- intersection

## Usage

Same as SCAD but with type safety, the ability to chain operations and an easy way to create your own complex shapes.

The following code creates the 3D model below.

```go
func main() {
    s := cad.NewSphere(11).Difference(
		cad.NewSphere(3).Translate(0, 0, 8).Scale(1.5, 1.5, 1.5),
        cad.NewSphere(3).Translate(0, 0, -11),
        cad.NewSphere(3).Translate(11, 0, 0),
        cad.NewSphere(3).Translate(-11, 0, 0),
        cad.NewSphere(3).Translate(0, 11, 0),
        cad.NewSphere(3).Translate(0, -11, 0),
    )
    c := cad.NewCube(20, 20, 20).Center().Difference(cad.NewSphere(12), cad.NewCube(20, 20, 20))
    d := cad.NewDocument(s, c)
    fmt.Println(d.Render())
}
```

![example image](example1.png "Example")

## Build your own shapes

Creating custom shapes is simple. For example, the `box` shape that is included in this library. `Render` is the only method you need to define. The `NewBox` function is just for convenience.

```go
type Box struct {
	x      float64
	y      float64
	z      float64
	radius float64
	wall   float64
	Shape
}

func (b *Box) Render() string {
	base := NewCube(b.x, b.y, b.z)
	radius := 0.01
	if b.radius != 0 {
		radius = b.radius
	}

	diff := NewRoundedCube(b.x-2*b.wall, b.y-2*b.wall, b.z+radius, radius).Translate(b.wall, b.wall, b.wall)
	res := base.Difference(diff)
	return res.Render()
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
```

Then use it like this.

```go
func main() {
	b := cad.NewBox(120, 60, 20, 5, 5)
	d := cad.NewDocument(b)
	fmt.Println(d.Render())
}
```

![example image](example2.png "Example")