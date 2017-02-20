package electron

import "github.com/gopherjs/gopherjs/js"

// NativeImageModule version@1.4.15
//
// Create tray, dock, and application icons using PNG or JPG files.
type NativeImageModule struct {
	*js.Object
	// Creates an empty NativeImage instance.
	CreateEmpty func() (Obj *NativeImage) `js:"createEmpty"`
	// Creates a new NativeImage instance from a file located at path. This method returns an empty image if the path does not exist, cannot be read, or is not a valid image.
	CreateFromPath func(Path string) (Obj *NativeImage) `js:"createFromPath"`
	// Creates a new NativeImage instance from buffer.
	CreateFromBuffer func(Buffer *js.Object, Options *NativeImageModuleCreateFromBufferOptions) (Obj *NativeImage) `js:"createFromBuffer"`
	// Creates a new NativeImage instance from dataURL.
	CreateFromDataURL func(DataURL string) `js:"createFromDataURL"`
}

func GetNativeImageModule() *NativeImageModule {
	o := Get("nativeImage")
	return &NativeImageModule{
		Object: o,
	}
}

type NativeImageModuleCreateFromBufferOptions struct {
	*js.Object
	// Required for bitmap buffers.
	Width int64 `js:"width"`
	// Required for bitmap buffers.
	Height int64 `js:"height"`
	// Defaults to 1.0.
	ScaleFactor float64 `js:"scaleFactor"`
}
