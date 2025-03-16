package main

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/configs"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/entity"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces"
	webserverInterface "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/repository"
	auth "github.com/gabrielmellooliveira/competitions-manager-api/internal/infra/auth"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/infra/database"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/infra/http"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/infra/middleware"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/infra/queue"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/infra/security"
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/infra/webserver"
	user "github.com/gabrielmellooliveira/competitions-manager-api/internal/usecase/auth"
	broadcast "github.com/gabrielmellooliveira/competitions-manager-api/internal/usecase/broadcast"
	competition "github.com/gabrielmellooliveira/competitions-manager-api/internal/usecase/competition"
	match "github.com/gabrielmellooliveira/competitions-manager-api/internal/usecase/match"
	supporter "github.com/gabrielmellooliveira/competitions-manager-api/internal/usecase/supporter"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Http
	httpAdapter := http.NewHttpAdapter(config.FootballApiUrl)
	httpAdapter.AddHeader("X-Auth-Token", config.FootballApiAuthToken)

	// Password Hasher
	passwordHasher := security.NewBcryptHasher()

	// JWT Autenticator
	jwtAuthenticator := auth.NewJwtAuthenticator(config.JwtSecret)

	// RabbitMq
	rabbitmq := queue.NewRabbitMqAdapter(config.RabbitMQUrl)

	// Web Server
	webserverHandler := webserver.NewWebServerHandler()

	// Database
	postgresUrl := "host=" + config.PostgresHost + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname=" + config.PostgresDatabase + " port=" + config.PostgresPort + " sslmode=disable"
	database := database.NewGormAdapter(postgresUrl)

	database.Connect()

	database.MigrateEntity(&entity.User{})
	database.MigrateEntity(&entity.Supporter{})
	database.MigrateEntity(&entity.Notification{})

	// Middlewares
	middleware := middleware.NewJwtMiddleware(jwtAuthenticator)
	middlewares := []webserverInterface.Middleware{middleware}

	// Repositories
	userRepository := repository.NewUserRepository(database)
	supporterRepository := repository.NewSupporterRepository(database)
	notificationRepository := repository.NewNotificationRepository(database)

	routes := []interfaces.Route{
		{
			Path:    "/auth/login",
			Method:  "POST",
			UseCase: user.NewLoginUseCase(userRepository, passwordHasher, jwtAuthenticator),
		},
		{
			Path:    "/auth/registar",
			Method:  "POST",
			UseCase: user.NewSignUpUseCase(userRepository, passwordHasher),
		},
		{
			Path:        "/campeonatos",
			Method:      "GET",
			UseCase:     competition.NewListCompetitionsUseCase(httpAdapter),
			Middlewares: middlewares,
		},
		{
			Path:        "/campeonatos/:competitionId/partidas",
			Method:      "GET",
			UseCase:     match.NewListMatchesUseCase(httpAdapter),
			Middlewares: middlewares,
		},
		{
			Path:        "/torcedores",
			Method:      "POST",
			UseCase:     supporter.NewRegisterSupporterUseCase(supporterRepository),
			Middlewares: middlewares,
		},
		{
			Path:        "/broadcast",
			Method:      "POST",
			UseCase:     broadcast.NewNotifySupportersUseCase(supporterRepository, notificationRepository, rabbitmq),
			Middlewares: middlewares,
		},
	}

	internal.Router(webserverHandler, routes)

	// Web Server
	go webserverHandler.Run()

	// Consumer
	sendEmailSupportersUseCase := broadcast.NewSendEmailSupportersUseCase(notificationRepository, rabbitmq)
	sendEmailSupportersUseCase.Execute()
}
