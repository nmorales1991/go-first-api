package routers

import (
	"encoding/json"
	"fmt"
	"github.com/nmorales1991/go-first-api/bd"
	"github.com/nmorales1991/go-first-api/models"
	"net/http"
)

func RegistroUsuario(writer http.ResponseWriter, request *http.Request) {
	var usuario models.Usuario
	// el body es de tipo stream, se lee 1 vez y se destruye. Lo decodifico y lo guardo con un puntero a usuario
	err := json.NewDecoder(request.Body).Decode(&usuario)
	fmt.Println(usuario)
	if err != nil {
		http.Error(writer, "Error en los datos recibidos"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(usuario.Nombre) == 0 || len(usuario.Correo) == 0 || len(usuario.Clave) == 0 {
		http.Error(writer, "Faltan datos requeridos", http.StatusBadRequest)
		return
	}
	// llamo al método que insertará el registro
	id, status, err := bd.InsertoRegistro(usuario)
	if err != nil {
		http.Error(writer, "Error al insertar "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(writer, "Hubo otro error al insertar el usuario ", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(id))
}
