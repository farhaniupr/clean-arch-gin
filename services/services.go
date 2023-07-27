package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewAuthorService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(NewComicService),
	fx.Provide(NewMemberService),
)
