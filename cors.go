package cors

import (
	"github.com/ivpusic/cmap"
	"github.com/ivpusic/neo"
)

func Init(conf cmap.C) func(ctx *neo.Ctx, next neo.Next) {
	allowOrigin := conf.StrOrDef("Access-Control-Allow-Origin", "*")
	allowCredentials := conf.StrOrDef("Access-Control-Allow-Credentials", "true")
	allowMethods := conf.StrOrDef("Access-Control-Allow-Methods", "*")
	allowHeaders := conf.StrOrDef("Access-Control-Allow-Headers", "*")

	return func(ctx *neo.Ctx, next neo.Next) {
		ctx.Res.Header.Set("Access-Control-Allow-Origin", allowOrigin)
		ctx.Res.Header.Set("Access-Control-Allow-Credentials", allowCredentials)
		ctx.Res.Header.Set("Access-Control-Allow-Methods", allowMethods)
		ctx.Res.Header.Set("Access-Control-Allow-Headers", allowHeaders)

		next()
	}
}
