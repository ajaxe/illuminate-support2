package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

func AppNavBar() *NavBar {
	return &NavBar{}
}

func AppBanner() *Banner {
	return &Banner{}
}
func AppLogo() *Logo {
	return &Logo{}
}

func AppNavBarItems(options NavListOptions) *NavBarItems {
	return &NavBarItems{
		itemTextColor: options.TextColor,
		listCSS:       options.ListCSS,
		listParent:    app.Ul().Class("navbar-nav"),
	}
}

func AppNavLeftPanel() *NavLeftPanel {
	return &NavLeftPanel{}
}
