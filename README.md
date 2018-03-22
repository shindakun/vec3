# vec3

<p align="center">
  <img style="float: right;" src="assets/vec3.png" alt="vec3 gopher"/>
</p>

Simple Golang vector3 library, basically the same as the one presented on [Games With Go](https://gameswithgo.org) [episode 13](https://www.youtube.com/watch?v=SjqXBrGpRM8).

## Installation

If you are a Go user you can run `go install github.com/shindakun/vec3`.

## Functions

func Distance(a, b Vector3) float32
func DistanceSquared(a, b Vector3) float32
type Vector3
    func Add(a, b Vector3) Vector3
    func Mult(a Vector3, b float32) Vector3
    func Normalize(a Vector3) Vector3
    func (a Vector3) Length() float32

## Usage example

The following code

```go
v1 := vec3.Vector3{X: 1, Y: 1, Z: 0}
v2 := vec3.Vector3{X: 1, Y: 1, Z: 0}
newV := vec3.Add(v1, v2)
fmt.Println(newV)
```

results in the new vector of `{2 2 0}`.