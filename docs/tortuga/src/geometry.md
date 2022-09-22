# Geometry

The `geom` library provides several geometry primatives that can be drawn to the screen.

With the `geom` library, we can build primitives such as:
* circles
* rectangles
* triangles
* lines

Most primitives implement the `Render` method that accept two arguments.
i.e. the display to draw to and the color that the primitive should be drawn as.


e.g.

```go
rect := geom.MakeRect(20, 20, 20, 20)

// note the render calls can only be called inside a render loop
rect.Render(display, 2)
```
