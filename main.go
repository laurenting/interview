package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed appIcon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:             "interview",
		Width:             1024,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		MaxWidth:          1024,
		MaxHeight:         768,
		DisableResize:     true,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 145, G: 145, B: 233, A: 100},
		Assets:            assets,
		Menu:              nil,
		Logger:            logger.NewDefaultLogger(),
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Interview",
				Message: "Copyright © 2022 Renting Liu and DAOZHUODA. All rights reserved.",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
