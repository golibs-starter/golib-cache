package golibcache

import (
	"gitlab.com/golibs-starter/golib"
	"go.uber.org/fx"
)

func EnableCache() fx.Option {
	return fx.Options(
		golib.ProvideProps(NewCacheProperties),
		fx.Provide(NewCache),
	)
}
