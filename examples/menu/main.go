package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	electron "github.com/oskca/gopherjs-electron"
	"github.com/oskca/gopherjs-electron/api/main/menu"
	nodejs "github.com/oskca/gopherjs-nodejs"
)

func main() {
	app := electron.GetApp()
	app.On(electron.EvtAppReady, func(args ...*js.Object) {
		opt := electron.NewBrowserWindowOption()
		bw := electron.NewBrowserWindow(opt)
		subm := electron.BuildFromTemplateEx([]electron.MenuItemOptionEx{
			{
				Label: "sub1",
				ClickEx: func(mi *electron.MenuItem) {
					println("sub1 clicked")
				},
			},
		})
		m := electron.BuildFromTemplateEx([]electron.MenuItemOptionEx{
			{
				Label: "Show &Help",
				ClickEx: func(mi *electron.MenuItem) {
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
				SubMenuOptions: []electron.MenuItemOptionEx{
					{
						Label: "sub option",
						ClickEx: func(mi *electron.MenuItem) {
							println(mi.Label, "==> called")
						},
					},
				},
			},
		})
		bw.SetMenu(m)
		bw.LoadURL("file://"+nodejs.DirName()+"/index.html", nil)
	})
}
