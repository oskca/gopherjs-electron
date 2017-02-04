package nativeimage

import (
	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
)

// Class: NativeImage
// ===============================================================

// NativeImage Natively wrap images such as tray, dock, and application icons.
// Process: Main, Renderer
type NativeImage struct {
	*js.Object

	// Instance Methods
	// The following methods are available on instances of the NativeImage class:

	// image.toPNG()
	// Returns Buffer - A Buffer that contains the image’s PNG encoded data.
	ToPNG func() *js.Object `js:"toPNG"`

	// image.toJPEG(quality)
	// quality Integer (required) - Between 0 - 100.
	// Returns Buffer - A Buffer that contains the image’s JPEG encoded data.
	ToJPEG func(quality int) *js.Object `js:"toJPEG"`

	// image.toBitmap()
	// Returns Buffer - A Buffer that contains a copy of the image’s raw bitmap pixel data.
	ToBitmap func() *js.Object `js:"toBitmap"`

	// image.toDataURL()
	// Returns String - The data URL of the image.
	ToDataURL func() string `js:"toDataURL"`

	// image.getBitmap()
	// Returns Buffer - A Buffer that contains the image’s raw bitmap pixel data.
	GetBitmap func() *js.Object `js:"getBitmap"`

	// The difference between getBitmap() and toBitmap() is, getBitmap() does not copy the bitmap data, so you have to use the returned Buffer immediately in current event loop tick, otherwise the data might be changed or destroyed.

	// image.getNativeHandle() macOS
	// Returns Buffer - A Buffer that stores C pointer to underlying native handle of the image. On macOS, a pointer to NSImage instance would be returned.

	// Notice that the returned pointer is a weak pointer to the underlying native image instead of a copy, so you must ensure that the associated nativeImage instance is kept around.

	// image.isEmpty()
	// Returns Boolean - Whether the image is empty.
	IsEmpty func() bool `js:"isEmpty"`

	// image.getSize()
	// Returns Object:
	//
	// width Integer
	// height Integer
	GetSize func() *js.Object `js:"getSize"`

	// image.setTemplateImage(option)
	// option Boolean
	// Marks the image as a template image.
	SetTemplateImage func(opt bool) `js:"setTemplateImage"`

	// image.isTemplateImage()
	// Returns Boolean - Whether the image is a template image.
	IsTemplateImage func() bool `js:"isTemplateImage"`

	// image.crop(rect)
	// rect Object - The area of the image to crop
	// x Integer
	// y Integer
	// width Integer
	// height Integer
	// Returns NativeImage - The cropped image.
	Crop func(rect *js.Object) `js:"crop"`

	// image.resize(options)
	// options Object
	// width Integer (optional)
	// height Integer (optional)
	// quality String (optional) - The desired quality of the resize image. Possible values are good, better or best. The default is best. These values express a desired quality/speed tradeoff. They are translated into an algorithm-specific method that depends on the capabilities (CPU, GPU) of the underlying platform. It is possible for all three methods to be mapped to the same algorithm on a given platform.
	// Returns NativeImage - The resized image.
	//
	// If only the height or the width are specified then the current aspect ratio will be preserved in the resized image.
	Resize func(options *js.Object) `js:"resize"`

	// image.getAspectRatio()
	// Returns Float - The image’s aspect ratio.
	GetAspectRatio func() float64 `js:"getAspectRatio"`
}

// Methods
// =================================================================================

// The nativeImage module has the following methods, all of which return an instance of the NativeImage class:

// nativeImage.createEmpty()
// Returns NativeImage
//
// Creates an empty NativeImage instance.
func CreateEmpty() *NativeImage {
	return &NativeImage{
		Object: electron.Get("nativeImage").Call("createEmpty"),
	}
}

// nativeImage.createFromPath(path)
// path String
// Returns NativeImage
//
// Creates a new NativeImage instance from a file located at path. This method returns an empty image if the path does not exist, cannot be read, or is not a valid image.
func CreateFromPath(path string) *NativeImage {
	return &NativeImage{
		Object: electron.Get("nativeImage").Call("createFromPath", path),
	}
}

// const nativeImage = require('electron').nativeImage

// let image = nativeImage.createFromPath('/Users/somebody/images/icon.png')
// console.log(image)

// nativeImage.createFromBuffer(buffer[, options])
// buffer Buffer
// options Object (optional)
// width Integer (optional) - Required for bitmap buffers.
// height Integer (optional) - Required for bitmap buffers.
// scaleFactor Double (optional) - Defaults to 1.0.
// Returns NativeImage
func CreateFromBuffer(buffer *js.Object, opts ...*js.Object) *NativeImage {
	img := &NativeImage{}
	if len(opts) > 0 {
		img.Object = electron.Get("nativeImage").Call("createFromBuffer", buffer, opts[0])
	} else {
		img.Object = electron.Get("nativeImage").Call("createFromBuffer", buffer)
	}
	return img
}

// Creates a new NativeImage instance from buffer.

// nativeImage.createFromDataURL(dataURL)
// dataURL String
// Creates a new NativeImage instance from dataURL.
func CreateFromDataURL(dataURL string) *NativeImage {
	return &NativeImage{
		Object: electron.Get("nativeImage").Call("createFromDataURL", dataURL),
	}
}
