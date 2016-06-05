package main
import (
	"github.com/labstack/echo"
	_ "github.com/labstack/echo/engine/standard"
)
func autoGenerateRouter(e *echo.Echo) *echo.Echo{
    e.GET("/sample", getSample)
        e.GET("/music/artist:id", getMusicArtist)
        e.POST("/music/artist", postMusicArtist)
        e.DELETE("/music/artist", deleteMusicArtist)
        e.GET("/mobile/list:id", getMobileList)
            e.GET("/mobile/smartphone/list:id", getMobileSmartphoneList)
                    e.GET("/mobile/smartphone/ios/version/changelog", getMobileSmartphoneIosVersionChangelog)
                    e.GET("/mobile/smartphone/ios/app/list", getMobileSmartphoneIosAppList)
                e.POST("/mobile/smartphone/ios/app", postMobileSmartphoneIosApp)

    return e
}