package entities

import "fmt"

type Store struct {
	Id            int64
	Tienda        string
	Nombrevirtual string
	Url           string
	Moneda        string
}

func (store Store) ToString() string {
	return fmt.Sprintf("id: %d\n tienda: %s\n nombrevirtual: %s\n URL_tienda_virtual: %s\n MonedaTienda: %s",
		store.Id, store.Tienda, store.Nombrevirtual, store.Url, store.Moneda)
}
