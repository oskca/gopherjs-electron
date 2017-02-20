package electron

import "github.com/gopherjs/gopherjs/js"

// DesktopCapturerModule version@1.4.15
//
// Access information about media sources that can be used to capture audio and
//video from the desktop using the navigator.webkitGetUserMedia API.
type DesktopCapturerModule struct {
	*js.Object
	// Starts gathering information about all available desktop media sources, and calls callback(error, sources) when finished. sources is an array of DesktopCapturerSource objects, each DesktopCapturerSource represents a screen or an individual window that can be captured.
	GetSources func(Options *DesktopCapturerModuleGetSourcesOptions, Callback DesktopCapturerModuleGetSourcesCallback) `js:"getSources"`
}

func GetDesktopCapturerModule() *DesktopCapturerModule {
	o := Get("desktopCapturer")
	return &DesktopCapturerModule{
		Object: o,
	}
}

type DesktopCapturerModuleGetSourcesOptions struct {
	*js.Object
	// An array of Strings that lists the types of desktop sources to be captured, available types are and .
	Types *js.Object `js:"types"`
	// The suggested size that the media source thumbnail should be scaled to, defaults to .
	ThumbnailSize *DesktopCapturerModuleOptionsThumbnailSize `js:"thumbnailSize"`
}

type DesktopCapturerModuleOptionsThumbnailSize struct {
	*js.Object
}

type DesktopCapturerModuleGetSourcesCallback func(Error *js.Object, Sources *js.Object)
