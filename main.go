package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	var (
		app = NewApp()
		//ctx      = app.ctx
		appMenu  = menu.NewMenu()
		fileMenu = appMenu.AddSubmenu("")
	)

	fileMenu.AddText("toggle", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.WindowMinimise(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "cover",
		Width:  1000,
		Height: 30,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 30, G: 30, B: 30, A: 10},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Menu:        appMenu,
		AlwaysOnTop: true,
		Frameless:   true,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
