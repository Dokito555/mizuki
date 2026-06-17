package main

import "github.com/Dokito555/mizuki/internal/configs"


func main() {
	viper := configs.NewViper()
	log := configs.NewLogger(viper)
	db := configs.NewDB(viper, log)
	app := configs.NewGin()

	// inject configs to app
	configs.Bootstrap(&configs.BootstrapConfig{
		App:         app,
		DB:          db,
		Log:         log,
		Config:      viper,
	})

	port := viper.GetString("APP_PORT")
	log.Info("Listening on port: " + port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}