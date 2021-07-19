package models

import (
	"database/sql"

	"fake.com/apirest/entities"
)

type StoreModel struct {
	Db *sql.DB
}

func (storeModel StoreModel) GetStores(tipoTienda int64) ([]entities.Store, error) {
	rows, err := storeModel.Db.Query("EXEC SELECT_TIENDAS_VIRTUALES_WOLK ?", tipoTienda)
	if err != nil {
		return nil, err
	} else {
		stores := []entities.Store{}
		for rows.Next() {
			var id int64
			var tienda string
			var nombrevirtual string
			var URL_tienda_virtual string
			var MonedaTienda string
			err2 := rows.Scan(&id, &tienda, &nombrevirtual, &URL_tienda_virtual, &MonedaTienda)
			if err2 != nil {
				return nil, err2
			} else {
				store := entities.Store{Id: id, Tienda: tienda, Nombrevirtual: nombrevirtual, Url: URL_tienda_virtual, Moneda: MonedaTienda}
				stores = append(stores, store)
			}
		}
		return stores, nil
	}
}
