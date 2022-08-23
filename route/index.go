package route

import (
	"KappaApi/controller"
	kappa "github.com/517962189/Kappa"
)

func init(){
	route := kappa.NewKappaServer().Gin()
	route.GET("/index",controller.Index)
}
