/* Archivo: maodelos.go
   Autor: Edwin Yoner
   Fecha: 27/10/2024
   Versión: 1.0

   Este archivo muestra cómo consumir una API externa en Go.
*/

package api

// Usuario representa la estructura de datos de un usuario
type Usuario struct {
	ID        int    `json:"id"`
	Nombre    string `json:"name"`
	Usuario   string `json:"username"`
	Email     string `json:"email"`
	Telefono  string `json:"phone"`
	Sitio     string `json:"website"`
	Direccion struct {
		Calle      string `json:"street"`
		Suite      string `json:"suite"`
		Ciudad     string `json:"city"`
		CodigoPost string `json:"zipcode"`
		Geo        struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Empresa struct {
		Nombre        string `json:"name"`
		Eslogan       string `json:"catchPhrase"`
		ActividadBase string `json:"bs"`
	} `json:"company"`
}

// Post representa la estructura de un post
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Titulo string `json:"title"`
	Cuerpo string `json:"body"`
}
