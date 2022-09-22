# Palette
Tortuga provides a default palette.
However, each cart can define it's own palette.
<!-- 
```go

type cart struct{}

func (c cart) Update() {

}

func (c cart) Render() {

}

func main() {

}
``` -->

A custom palette can be defined in a `.palette` file. `.palette` files should be in csv format.
Each record contains 3 numbers(r, g, b) 0-255.

