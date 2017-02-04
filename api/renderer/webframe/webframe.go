// Package webframe customizes the rendering of the current web page.
// Process: Renderer
// An example of zooming current page to 200%.
//
//  const {webFrame} = require('electron')
//
//  webFrame.setZoomFactor(2)
package webframe

import (
	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
)

// webFrame

// The webFrame module has the following methods:

// Get returns the webFrame object
func Get() *WebFrame {
	return &WebFrame{
		Object: electron.Get("webFrame"),
	}
}

// WebFrame customizes the rendering of the current web page.
type WebFrame struct {
	*js.Object
	// webFrame.setZoomFactor(factor)
	// factor Number - Zoom factor.
	// Changes the zoom factor to the specified factor. Zoom factor is zoom percent divided by 100, so 300% = 3.0.
	SetZoomFactor func(factor float64) `js:"setZoomFactor"`

	// webFrame.getZoomFactor()
	// Returns Number - The current zoom factor.
	GetZoomFactor func() (factor float64) `js:"getZoomFactor"`

	// webFrame.setZoomLevel(level)
	// level Number - Zoom level
	// Changes the zoom level to the specified level. The original size is 0 and each increment above or below represents zooming 20% larger or smaller to default limits of 300% and 50% of original size, respectively.
	SetZoomLevel func(level float64) `js:"setZoomlevel"`

	// webFrame.getZoomLevel()
	// Returns Number - The current zoom level.
	GetZoomLevel func() (level float64) `js:"getZoomlevel"`

	// webFrame.setZoomLevelLimits(minimumLevel, maximumLevel)
	// minimumLevel Number
	// maximumLevel Number
	// Deprecated: Call setVisualZoomLevelLimits instead to set the visual zoom level limits. This method will be removed in Electron 2.0.
	SetZoomLevelLimits func(minimumLevel, maximumLevel float64) `js:"setZoomLevelLimits"`

	// webFrame.setVisualZoomLevelLimits(minimumLevel, maximumLevel)
	// minimumLevel Number
	// maximumLevel Number
	// Sets the maximum and minimum pinch-to-zoom level.
	SetVisualZoomLevelLimits func(minimumLevel, maximumLevel float64) `js:"setVisualZoomLevelLimits"`

	// webFrame.setLayoutZoomLevelLimits(minimumLevel, maximumLevel)
	// minimumLevel Number
	// maximumLevel Number
	// Sets the maximum and minimum layout-based (i.e. non-visual) zoom level.
	SetLayoutZoomLevelLimits func(minimumLevel, maximumLevel float64) `js:"setLayoutZoomLevelLimits"`

	// webFrame.setSpellCheckProvider(language, autoCorrectWord, provider)
	// language String
	// autoCorrectWord Boolean
	// provider Object
	// spellCheck Function - Returns Boolean
	// text String
	// Sets a provider for spell checking in input fields and text areas.
	// The provider must be an object that has a spellCheck method that returns whether
	// the word passed is correctly spelled.
	// An example of using node-spellchecker as provider:
	//
	// const {webFrame} = require('electron')
	// webFrame.setSpellCheckProvider('en-US', true, {
	//   spellCheck (text) {
	//     return !(require('spellchecker').isMisspelled(text))
	//   }
	// })
	SetSpellCheckProvider func(language string, autoCorrectWord bool, provider *js.Object) `js:"setSpellCheckProvider"`

	// webFrame.registerURLSchemeAsSecure(scheme)
	// scheme String
	// Registers the scheme as secure scheme.

	// Secure schemes do not trigger mixed content warnings. For example, https and data are secure schemes because they cannot be corrupted by active network attackers.
	RegisterURLSchemeAsSecure func(scheme string) `js:"registerURLSchemeAsSecure"`

	// webFrame.registerURLSchemeAsBypassingCSP(scheme)
	// scheme String
	// Resources will be loaded from this scheme regardless of the current page’s Content Security Policy.
	registerURLSchemeAsBypassingCSP func(scheme string) `js:"registerURLSchemeAsBypassingCSP"`

	// webFrame.registerURLSchemeAsPrivileged(scheme[, options])
	// scheme String
	// options Object (optional)
	// secure Boolean - (optional) Default true.
	// bypassCSP Boolean - (optional) Default true.
	// allowServiceWorkers Boolean - (optional) Default true.
	// supportFetchAPI Boolean - (optional) Default true.
	// corsEnabled Boolean - (optional) Default true.
	// Registers the scheme as secure, bypasses content security policy for resources,
	// allows registering ServiceWorker and supports fetch API.
	//
	// Specify an option with the value of false to omit it from the registration.
	// An example of registering a privileged scheme, without bypassing Content Security Policy:
	//  const {webFrame} = require('electron')
	//  webFrame.registerURLSchemeAsPrivileged('foo', { bypassCSP: false })
	//  webFrame.insertText(text)
	// text String
	// Inserts text to the focused element.
	RegisterURLSchemeAsPrivileged func(scheme string, opts ...*js.Object) `js:"registerURLSchemeAsPrivileged"`

	// webFrame.executeJavaScript(code[, userGesture, callback])
	// code String
	// userGesture Boolean (optional) - Default is false.
	// callback Function (optional) - Called after script has been executed.
	// result Any
	// Evaluates code in page.
	//
	// In the browser window some HTML APIs like requestFullScreen can only be invoked by a gesture from the user. Setting userGesture to true will remove this limitation.
	ExecuteJavaScript func(code string) `js:"executeJavaScript"`

	// webFrame.getResourceUsage()
	// Returns Object:
	//
	// images MemoryUsageDetails
	// cssStyleSheets MemoryUsageDetails
	// xslStyleSheets MemoryUsageDetails
	// fonts MemoryUsageDetails
	// other MemoryUsageDetails
	// Returns an object describing usage information of Blink’s internal memory caches.

	// const {webFrame} = require('electron')
	// console.log(webFrame.getResourceUsage())
	// This will generate:
	//
	// {
	//   images: {
	//     count: 22,
	//     size: 2549,
	//     liveSize: 2542,
	//     decodedSize: 478,
	//     purgedSize: 0,
	//     purgeableSize: 0
	//   },
	//   cssStyleSheets: { /* same with "images" */ },
	//   xslStyleSheets: { /* same with "images" */ },
	//   fonts: { /* same with "images" */ },
	//   other: { /* same with "images" */ }
	// }
	getResourceUsage func() (obj *js.Object) `js:"getResourceUsage"`

	// webFrame.clearCache()
	// Attempts to free memory that is no longer being used (like images from a previous navigation).
	// Note that blindly calling this method probably makes Electron slower since it will have to
	// refill these emptied caches, you should only call it if an event in your app has occurred that
	// makes you think your page is actually using less memory (i.e.
	// you have navigated from a super heavy page to a mostly empty one, and intend to stay there).
	ClearCache func() `js:"clearCache"`
}
