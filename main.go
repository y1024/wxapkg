package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

var version = "v2.0.0"
var github = "https://github.com/wux1an/wxapkg"

func main() {
	application.RegisterEvent[string](EventUnpackProgress)

	var service = NewAppService()
	app := application.New(application.Options{
		Name: "微信小程序解包工具",
		Services: []application.Service{
			application.NewService(service),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	service.SetContext(app)

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "微信小程序解包工具",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		Width:  1024,
		Height: 768,
		//BackgroundColour: application.NewRGB(27, 38, 54),
		URL: "/",
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
