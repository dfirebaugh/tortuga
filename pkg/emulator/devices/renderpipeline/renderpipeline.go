package renderpipeline

import "github.com/dfirebaugh/tortuga/pkg/texture"

type RenderPipeline struct {
	items []*texture.Texture
}

// AddToRenderPipeline allows you to push images into a queue that will be rendered
//
//	This is more efficient than rendering directly to the frame buffer
func (r *RenderPipeline) Append(t *texture.Texture) {
	t.Render()
	r.items = append(r.items, t)
}

func (r RenderPipeline) Get() []*texture.Texture {
	return r.items
}

func (r *RenderPipeline) Clear() {
	r.items = make([]*texture.Texture, 0)
}
