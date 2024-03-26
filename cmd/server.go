package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ericolvr/goapi/api"
	"github.com/ericolvr/goapi/config"
	"github.com/ericolvr/goapi/internal/adapter/database/postgres"
	"github.com/ericolvr/goapi/internal/domain"
	"github.com/ericolvr/goapi/internal/services/users"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")

		// router := gin.Default()

		// db, err := database.NewMySQLConnection()
		// if err != nil {
		// 	log.Printf("Failed to connect to MySQL database: %v", err)
		// 	os.Exit(1)
		// }

		// load config
		cfg, err := config.NewConfig()
		if err != nil {
			log.Printf("Failed to load config: %v", err)
			os.Exit(1)
		}

		// load posgres
		opts := []postgres.Options{
			postgres.WithHost(cfg.GetPostgres().Host),
			postgres.WithPort(cfg.GetPostgres().Port),
			postgres.WithPassword(cfg.GetPostgres().Pass),
		}

		pg, err := postgres.Connection(opts...)
		if err != nil {
			log.Printf("Failed to connect to Postgres database: %v", err)
			os.Exit(1)
		}

		defer pg.Close()

		uopts := []users.Options{
			users.WithPostgres(pg),
		}

		users, err := users.NewUsers(uopts...)
		if err != nil {
			log.Printf("Failed to create users service: %v", err)
			os.Exit(1)
		}

		if err := users.Create(domain.User{}); err != nil {
			log.Printf("Failed to create user: %v", err)
			os.Exit(1)
		}

		apiopts := []api.Options{
			api.WithUsers(users),
			api.WithConfig(cfg),
		}

		api, err := api.New(apiopts...)
		if err != nil {
			log.Printf("Failed to create api: %v", err)
			os.Exit(1)
		}

		//Server name and version
		log.Printf("%s - v.: %s", cfg.GetApp().Name, cfg.GetApp().Version)

		// Start the server
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// Load the routes
		api.LoadRoutes()
		// Start the server
		api.Start()

		sig := <-sigChan
		log.Printf("*** Signal %s received, shutting down\n", sig.String())

		// Shutdown http server
		if err := api.GracefulShutdown(); err != nil {
			log.Printf("Failed to shutdown server: %s", err)
		}
		log.Printf("Server HTTP shutdown successfully")

		// userRepo := repository.NewMySQLUserRepository(db)
		// equipmentRepo := repository.NewMySQLEquipmentRepository(db)

		// userUsecase := usecase.NewUserUsecase(userRepo)
		// equipmentUsecase := usecase.NewEquipmentUsecase(equipmentRepo)

		// http.NewUserHandler(router, userUsecase)
		// http.NewEquipmentHandler(router, equipmentUsecase)

		// router.Run(":8080")

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
