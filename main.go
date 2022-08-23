package main

import (
	_ "KappaApi/hooks"
	_ "KappaApi/route"
	"github.com/517962189/Kappa"
)

func main()  {
	kappaServer := kappa.NewKappaServer()
	addr := kappaServer.RegisterConfigStorage("kappa").GetString("main.HTTPAddr")
	kappaServer.Run(addr)
}
