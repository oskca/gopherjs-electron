package electron

import "github.com/gopherjs/gopherjs/js"

// WebFrameModule version@1.4.15
//
// Customize the rendering of the current web page.
type WebFrameModule struct {
	*js.Object
	// Changes the zoom factor to the specified factor. Zoom factor is zoom percent divided by 100, so 300% = 3.0.
	SetZoomFactor func(Factor float64) `js:"setZoomFactor"`
	GetZoomFactor func() (Obj float64) `js:"getZoomFactor"`
	// Changes the zoom level to the specified level. The original size is 0 and each increment above or below represents zooming 20% larger or smaller to default limits of 300% and 50% of original size, respectively.
	SetZoomLevel func(Level float64)  `js:"setZoomLevel"`
	GetZoomLevel func() (Obj float64) `js:"getZoomLevel"`
	// Deprecated: Call setVisualZoomLevelLimits instead to set the visual zoom level limits. This method will be removed in Electron 2.0.
	SetZoomLevelLimits func(MinimumLevel float64, MaximumLevel float64) `js:"setZoomLevelLimits"`
	// Sets the maximum and minimum pinch-to-zoom level.
	SetVisualZoomLevelLimits func(MinimumLevel float64, MaximumLevel float64) `js:"setVisualZoomLevelLimits"`
	// Sets the maximum and minimum layout-based (i.e. non-visual) zoom level.
	SetLayoutZoomLevelLimits func(MinimumLevel float64, MaximumLevel float64) `js:"setLayoutZoomLevelLimits"`
	// Sets a provider for spell checking in input fields and text areas. The provider must be an object that has a spellCheck method that returns whether the word passed is correctly spelled. An example of using node-spellchecker as provider:
	SetSpellCheckProvider func(Language string, AutoCorrectWord bool, Provider *WebFrameModuleSetSpellCheckProviderProvider) `js:"setSpellCheckProvider"`
	// Registers the scheme as secure scheme. Secure schemes do not trigger mixed content warnings. For example, https and data are secure schemes because they cannot be corrupted by active network attackers.
	RegisterURLSchemeAsSecure func(Scheme string) `js:"registerURLSchemeAsSecure"`
	// Resources will be loaded from this scheme regardless of the current page's Content Security Policy.
	RegisterURLSchemeAsBypassingCSP func(Scheme string) `js:"registerURLSchemeAsBypassingCSP"`
	// Registers the scheme as secure, bypasses content security policy for resources, allows registering ServiceWorker and supports fetch API. Specify an option with the value of false to omit it from the registration. An example of registering a privileged scheme, without bypassing Content Security Policy:
	RegisterURLSchemeAsPrivileged func(Scheme string, Options *WebFrameModuleRegisterURLSchemeAsPrivilegedOptions) `js:"registerURLSchemeAsPrivileged"`
	// Inserts text to the focused element.
	InsertText func(Text string) `js:"insertText"`
	// Evaluates code in page. In the browser window some HTML APIs like requestFullScreen can only be invoked by a gesture from the user. Setting userGesture to true will remove this limitation.
	ExecuteJavaScript func(Code string, UserGesture bool, Callback WebFrameModuleExecuteJavaScriptCallback) `js:"executeJavaScript"`
	// Returns an object describing usage information of Blink's internal memory caches. This will generate:
	GetResourceUsage func() (Obj *WebFrameModuleGetResourceUsageObj) `js:"getResourceUsage"`
	// Attempts to free memory that is no longer being used (like images from a previous navigation). Note that blindly calling this method probably makes Electron slower since it will have to refill these emptied caches, you should only call it if an event in your app has occurred that makes you think your page is actually using less memory (i.e. you have navigated from a super heavy page to a mostly empty one, and intend to stay there).
	ClearCache func() `js:"clearCache"`
}

func GetWebFrameModule() *WebFrameModule {
	o := Get("webFrame")
	return &WebFrameModule{
		Object: o,
	}
}

type WebFrameModuleSetSpellCheckProviderProvider struct {
	*js.Object
	// Returns
	SpellCheck WebFrameModuleProviderSpellCheck `js:"spellCheck"`
}

type WebFrameModuleProviderSpellCheck func(Text string)
type WebFrameModuleRegisterURLSchemeAsPrivilegedOptions struct {
	*js.Object
	// (optional) Default true.
	Secure bool `js:"secure"`
	// (optional) Default true.
	BypassCSP bool `js:"bypassCSP"`
	// (optional) Default true.
	AllowServiceWorkers bool `js:"allowServiceWorkers"`
	// (optional) Default true.
	SupportFetchAPI bool `js:"supportFetchAPI"`
	// (optional) Default true.
	CorsEnabled bool `js:"corsEnabled"`
}

type WebFrameModuleExecuteJavaScriptCallback func(Result *js.Object)
type WebFrameModuleGetResourceUsageObj struct {
	*js.Object
	Images         *js.Object `js:"images"`
	CssStyleSheets *js.Object `js:"cssStyleSheets"`
	XslStyleSheets *js.Object `js:"xslStyleSheets"`
	Fonts          *js.Object `js:"fonts"`
	Other          *js.Object `js:"other"`
}
