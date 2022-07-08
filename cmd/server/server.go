package server

import "log"

func Run() error {
	err := helpers.Load()
	if err != nil {
		log.Fatal(err)
		return err
	}

	db, err := repository.ConnectDb(&helpers.Config{
		DBUser:     helpers.Instance.DBUser,
		DBPass:     helpers.Instance.DBPass,
		DBHost:     helpers.Instance.DBHost,
		DBName:     helpers.Instance.DBName,
		DBPort:     helpers.Instance.DBPort,
		DBTimeZone: helpers.Instance.DBTimeZone,
		DBMode:     helpers.Instance.DBMode,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = repository.MigrateAll(db)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
