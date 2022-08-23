package hooks

import (
	"fmt"
	"github.com/517962189/Kappa"
	"github.com/gin-contrib/pprof"
)

func init(){
	kappa.NewKappaServer().RegisterHooksStorage(func() error {
		pprof.Register(kappa.NewKappaServer().Gin())
		fmt.Printf("[启用pprof] 注入 gin 成功 \n")
		return nil
	})
}
