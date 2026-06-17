package main

import "github.com/Dokito555/mizuki/internal/configs"


func main() {
	viper := configs.NewViper()
	log := configs.NewLogger(viper)
	app := configs.NewGin()

	// inject configs to app
	configs.Bootstrap(&configs.BootstrapConfig{
		App:         app,
		Log:         log,
		Config:      viper,
	})

	port := viper.GetString("APP_PORT")
	err := app.Run(":" + port)
	log.Info("Listening to port: " + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}