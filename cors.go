package cors

import (
	"github.com/ivpusic/cmap"
	"github.com/ivpusic/neo"
)

var (
	allowOriginKey = "Access-Control-Allow-Origin"
	allowOriginDef = "*"

	allowCredentialsKey = "Access-Control-Allow-Credentials"
	allowCredentialsDef = "true"

	allowMethodsKey = "Access-Control-Allow-Methods"
	allowMethodsDef = "*"

	allowHeadersKey = "Access-Control-Allow-Headers"
	allowHeadersDef = "*"

	maxAgeKey = "Access-Control-Max-Age"
	maxAgeDef = "600"

	exposeHeadersKey = "Access-Control-Expose-Headers"
	exposeHeadersDef = ""
)

// Init neo-cors with default options
func Init() func(ctx *neo.Ctx, next neo.Next) {
	return func(ctx *neo.Ctx, next neo.Next) {
		ctx.Res.Header().Set(allowOriginKey, allowOriginDef)
		ctx.Res.Header().Set(allowCredentialsKey, allowCredentialsDef)
		ctx.Res.Header().Set(allowMethodsKey, allowMethodsDef)
		ctx.Res.Header().Set(allowHeadersKey, allowHeadersDef)
		ctx.Res.Header().Set(maxAgeKey, maxAgeDef)
		ctx.Res.Header().Set(exposeHeadersKey, exposeHeadersDef)

		next()
	}
}

// Init neo-cors with custom options
func InitWithOptions(conf cmap.C) func(ctx *neo.Ctx, next neo.Next) {
	allowOrigin := conf.StrOrDef(allowOriginKey, allowOriginDef)
	allowCredentials := conf.StrOrDef(allowCredentialsKey, allowCredentialsDef)
	allowMethods := conf.StrOrDef(allowMethodsKey, allowMethodsDef)
	allowHeaders := conf.StrOrDef(allowHeadersKey, allowHeadersDef)
	maxAge := conf.StrOrDef(maxAgeKey, maxAgeDef)
	exposeHeaders := conf.StrOrDef(exposeHeadersKey, exposeHeadersDef)

	return func(ctx *neo.Ctx, next neo.Next) {
		ctx.Res.Header().Set(allowOriginKey, allowOrigin)
		ctx.Res.Header().Set(allowCredentialsKey, allowCredentials)
		ctx.Res.Header().Set(allowMethodsKey, allowMethods)
		ctx.Res.Header().Set(allowHeadersKey, allowHeaders)
		ctx.Res.Header().Set(maxAgeKey, maxAge)
		ctx.Res.Header().Set(exposeHeadersKey, exposeHeaders)

		next()
	}
}
