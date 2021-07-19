package main

import (
	"fake.com/apirest/config"
	"fake.com/apirest/entities"
	"fake.com/apirest/models"
)

func getStores(tipoTienda int64) ([]entities.Store, error) {
	//Declare an array because if there's error, we return it empty
	stores := []entities.Store{}
	db, err := config.GetDB()
	if err != nil {
		return stores, err
	}

	storeModel := models.StoreModel{
		Db: db,
	}

	//obtenenemos las tienddas segun los tipo tiendas
	stores, err2 := storeModel.GetStores(tipoTienda)
	if err2 != nil {
		return stores, err
	}
	return stores, nil
}

// func createVideoGame(videoGame VideoGame) error {
// 	bd, err := getDB()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", videoGame.Name, videoGame.Genre, videoGame.Year)
// 	return err
// }

// func deleteVideoGame(id int64) error {

// 	bd, err := getDB()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = bd.Exec("DELETE FROM video_games WHERE id = ?", id)
// 	return err
// }

// // It takes the ID to make the update
// func updateVideoGame(videoGame VideoGame) error {
// 	bd, err := getDB()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = bd.Exec("UPDATE video_games SET name = ?, genre = ?, year = ? WHERE id = ?", videoGame.Name, videoGame.Genre, videoGame.Year, videoGame.Id)
// 	return err
// }
