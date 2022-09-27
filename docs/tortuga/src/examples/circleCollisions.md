# Circle Collision

> controls: right click generates balls, left click removes them all

## Static Resolution
Static resolution ensures that the circles do not overlap.

```
go run github.com/dfirebaugh/tortuga/examples/staticball
```

<wasm-view height=400 width=530 src="staticball.wasm"></wasm-view>


## Dynamic Resolution
Dynamic resolution takes into account the velocity and mass of the circle and will deflect the circle into the opposite direction of the collision.

```
go run github.com/dfirebaugh/tortuga/examples/dynamicball
```

<wasm-view height=400 width=530 src="dynamicball.wasm"></wasm-view>

