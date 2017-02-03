package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/main/browserwindow"
	"github.com/oskca/gopherjs-electron/api/main/menu"
	nodejs "github.com/oskca/gopherjs-nodejs"
)

func main() {
	app := electron.GetApp()
	app.On(electron.EvtReady, func(args ...*js.Object) {
		opt := browserwindow.NewOption()
		bw := browserwindow.New(opt)
		subm := menu.BuildFromTemplate([]menu.Option{
			{
				Label: "sub1",
				ClickEx: func(mi *menu.Item) {
					println("sub1 clicked")
				},
			},
		})
		m := menu.BuildFromTemplate([]menu.Option{
			{
				Label: "Show &Help",
				ClickEx: func(mi *menu.Item) {
					println(time.Now().String(), "clicked:", mi.Label)
				},
			},
			{
				Label:   "submenu",
				SubMenu: subm,
			},
			{
				Type: menu.TypeSeparator,
			},
			{
				Role: menu.RoleAbout,
			},
			{
				Type: menu.TypeSeparator,
			},
			{
				Label: "sub options",
				SubMenuOptions: []menu.Option{
					{
						Label: "sub option",
						ClickEx: func(mi *menu.Item) {
							println(mi.Label, "==> called")
						},
					},
				},
			},
		})
		bw.SetMenu(m)
		bw.LoadURL("file://" + nodejs.DirName() + "/index.html")
	})
}
