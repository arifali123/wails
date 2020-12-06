package mac

import "github.com/wailsapp/wails/v2/pkg/menu"

// Options are options specific to Mac
type Options struct {
	TitleBar                      *TitleBar
	Appearance                    AppearanceType
	WebviewIsTransparent          bool
	WindowBackgroundIsTranslucent bool
	Menu                          *menu.Menu
	Tray                          *menu.Menu
}
