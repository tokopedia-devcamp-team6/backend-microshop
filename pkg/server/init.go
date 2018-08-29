package server

import (
	"team6minihack/pkg/domain"
	"team6minihack/pkg/lib/config"
	"team6minihack/pkg/usecase"

	"github.com/tokopedia/tdk/go/app"
	"github.com/tokopedia/tdk/go/app/resource"
)

var cfg config.Config

var orderDomain domain.OrderDomain
var orderUsecase *usecase.OrderUsecase
var userDomain domain.UserDomain

// We do all the wire up in this Init() function
// please return any error if you fail to initialize something
func Init(app *app.App) error {
	cfg = config.GetConfig()

	// you can init below variables using app.Resource()
	var db resource.SQLDB
	// using the name you specified in config, you can do this
	// db, _ = app.Resource().GetSQLDB("mydatabase")

	// or we can init redis
	var redis resource.Redis
	// redis, _ = app.Resource().GetRedis("myredis")

	orderDomain = domain.InitOrderDomain(
		// OrderResource here needed to separate
		// domain logic with data layer (resource)
		domain.OrderResource{db, redis},
	)

	userDomain = domain.InitUserDomain(
		domain.UserResource{db},
	)

	// in this case orderUsecase needs multiple domain
	// orderDomain and userDomain
	orderUsecase = usecase.InitOrderUsecase(
		cfg, orderDomain, userDomain,
	)

	return nil
}
