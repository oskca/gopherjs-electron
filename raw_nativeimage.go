package electron

import "github.com/gopherjs/gopherjs/js"

// NativeImage version@1.4.15
//
// Natively wrap images such as tray, dock, and application icons.
type NativeImage struct {
	*js.Object
	ToPNG     func() (Obj *js.Object)              `js:"toPNG"`
	ToJPEG    func(Quality int64) (Obj *js.Object) `js:"toJPEG"`
	ToBitmap  func() (Obj *js.Object)              `js:"toBitmap"`
	ToDataURL func() (Obj string)                  `js:"toDataURL"`
	// The difference between getBitmap() and toBitmap() is, getBitmap() does not copy the bitmap data, so you have to use the returned Buffer immediately in current event loop tick, otherwise the data might be changed or destroyed.
	GetBitmap func() (Obj *js.Object) `js:"getBitmap"`
	// Notice that the returned pointer is a weak pointer to the underlying native image instead of a copy, so you must ensure that the associated nativeImage instance is kept around.
	GetNativeHandle func() (Obj *js.Object)             `js:"getNativeHandle"`
	IsEmpty         func() (Obj bool)                   `js:"isEmpty"`
	GetSize         func() (Obj *NativeImageGetSizeObj) `js:"getSize"`
	// Marks the image as a template image.
	SetTemplateImage func(Option bool)                                  `js:"setTemplateImage"`
	IsTemplateImage  func() (Obj bool)                                  `js:"isTemplateImage"`
	Crop             func(Rect *NativeImageCropRect) (Obj *NativeImage) `js:"crop"`
	// If only the height or the width are specified then the current aspect ratio will be preserved in the resized image.
	Resize         func(Options *NativeImageResizeOptions) (Obj *NativeImage) `js:"resize"`
	GetAspectRatio func() (Obj float64)                                       `js:"getAspectRatio"`
}

func WrapNativeImage(o *js.Object) *NativeImage {
	return &NativeImage{
		Object: o,
	}
}

type NativeImageGetSizeObj struct {
	*js.Object
	Width  int64 `js:"width"`
	Height int64 `js:"height"`
}

type NativeImageCropRect struct {
	*js.Object
	X      int64 `js:"x"`
	Y      int64 `js:"y"`
	Width  int64 `js:"width"`
	Height int64 `js:"height"`
}

type NativeImageResizeOptions struct {
	*js.Object
	Width  int64 `js:"width"`
	Height int64 `js:"height"`
	// The desired quality of the resize image. Possible values are , or . The default is . These values express a desired quality/speed tradeoff. They are translated into an algorithm-specific method that depends on the capabilities (CPU, GPU) of the underlying platform. It is possible for all three methods to be mapped to the same algorithm on a given platform.
	Quality string `js:"quality"`
}
