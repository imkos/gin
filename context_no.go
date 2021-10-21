//go:build !noproto
// +build !noproto

package gin

import (
	"github.com/imkos/gin/render"
)

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func (c *Context) ProtoBuf(code int, obj interface{}) {
	c.Render(code, render.ProtoBuf{Data: obj})
}
