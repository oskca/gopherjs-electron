package electron

import "github.com/gopherjs/gopherjs/js"

// WebContentsModule version@1.4.15
//
// Render and control web pages.
type WebContentsModule struct {
	*js.Object
	GetAllWebContents     func() (Obj *js.Object)           `js:"getAllWebContents"`
	GetFocusedWebContents func() (Obj *WebContents)         `js:"getFocusedWebContents"`
	FromId                func(Id int64) (Obj *WebContents) `js:"fromId"`
}

func GetWebContentsModule() *WebContentsModule {
	o := Get("webContents")
	return &WebContentsModule{
		Object: o,
	}
}
