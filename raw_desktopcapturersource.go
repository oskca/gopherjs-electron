package electron

import "github.com/gopherjs/gopherjs/js"

// DesktopCapturerSource a Structure
type DesktopCapturerSource struct {
	*js.Object
	// The identifier of a window or screen that can be used as a constraint when calling []. The format of the identifier will be or , where is a random generated number.
	Id string `js:"id"`
	// A screen source will be named either or , while the name of a window source will match the window title.
	Name string `js:"name"`
	// A thumbnail image. There is no guarantee that the size of the thumbnail is the same as the specified in the passed to . The actual size depends on the scale of the screen or window.
	Thumbnail *NativeImage `js:"thumbnail"`
}
