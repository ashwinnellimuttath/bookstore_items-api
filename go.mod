module itemsModule

replace utils => ../bookstore_utils-go

replace common_go => ../bookstore_oauth-go

go 1.17

require (
	common_go v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mercadolibre/golang-restclient v0.0.0-20170701022150-51958130a0a0 // indirect
	github.com/olivere/elastic v6.2.37+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.20.0 // indirect
)
