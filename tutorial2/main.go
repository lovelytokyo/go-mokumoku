package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

//go:generate echo-generate-router -routerfile=main.go
func main() {
	e := echo.New()
	e = autoGenerateRouter(e)

	// create `/sample`
	// @router GET sample

	// create router group `/music`
	// @routerGroupStart music

	// create router `/music/artist`
	// @router GET artist:id
	// @router POST artist
	// @router DELETE artist

	// dist router group `/music`
	// @routerGroupEnd

	// @routerGroupStart mobile
	// @router GET list:id
	// @routerGroupStart smartphone
	// @router GET list:id
	// @routerGroupStart ios
	// @routerGroupStart version
	// @router GET changelog
	// @routerGroupEnd
	// @routerGroupStart app
	// @router GET list
	// @routerGroupEnd
	// @router POST app
	// @routerGroupEnd
	// @routerGroupEnd
	// @routerGroupEnd

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Run(standard.New(":7070"))
}

func getSample(c echo.Context) error {
	return c.String(http.StatusOK, "get sample")
}

func getMusicArtist(c echo.Context) error {
	return c.String(http.StatusOK, "get music artist")
}

func postMusicArtist(c echo.Context) error {
	return c.String(http.StatusOK, "post music artist")
}

func deleteMusicArtist(c echo.Context) error {
	return c.String(http.StatusOK, "delete music artist")
}

func getMobileList(c echo.Context) error {
	return c.String(http.StatusOK, "get mobile list")
}

func getMobileSmartphoneList(c echo.Context) error {
	return c.String(http.StatusOK, "get mobile smartphone list")
}

func getMobileSmartphoneIosVersionChangelog(c echo.Context) error {
	return c.String(http.StatusOK, "get mobile smartphone ios version changelog")
}

func getMobileSmartphoneIosAppList(c echo.Context) error {
	return c.String(http.StatusOK, "get mobile smartphone ios app list")
}

func postMobileSmartphoneIosApp(c echo.Context) error {
	return c.String(http.StatusOK, "post mobile smartphone ios app")
}