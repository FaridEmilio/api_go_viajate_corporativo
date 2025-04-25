package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/faridEmilio/api_go_viajate_corporativo_corporativo/internal/config"
	"github.com/gofiber/fiber/v2"
)

/*
	 	MiddlewareManager tiene un campo HTTPClient,
		que es un puntero a un cliente HTTP. Este cliente HTTP se utiliza para realizar
		solicitudes a servidores externos.
*/
type MiddlewareManager struct {
	HTTPClient *http.Client
}

/*
Esta función toma un parámetro scope y devuelve otra función
que actúa como un middleware de autorización para las solicitudes HTTP.
El middleware verifica si se proporciona un token de autorización en la cabecera Authorization
de la solicitud.
*/
func (m *MiddlewareManager) ValidarPermiso(scope string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// El middleware extrae el token de autorización de la cabecera Authorization
		// de la solicitud de Fiber y lo almacena en la variable bearer.
		bearer := c.Get("Authorization")

		// Aca validamos la longitud del token
		// Si es menor o igual a cero significa que no se proporcionó ningún token de autorización
		// en la solicitud.
		if len(bearer) <= 0 {
			return errors.New("acceso no autorizado, debe enviar un token de autenticación")
		}

		var result struct {
			Acceso string `json:"acceso"`
			ID     int64  `json:"user_id"`
		}

		/*
			Se analiza la URL base que se encuentra en la variable config.AUTH.
			Si hay un error al analizar la URL, se devuelve un error.
		*/
		base, err := url.Parse(config.AUTH)
		if err != nil {
			return fmt.Errorf("error al crear base url" + err.Error())
		}

		/*
			Se agrega "/users/permiso" a la ruta de la URL base.
			Esto se hace para construir la URL completa a la que se enviará la solicitud POST.
		*/
		base.Path += "/users/permiso"

		/*
			Preparación de datos para la solicitud: Se crea una estructura llamada values que contiene
			dos campos: SistemaID y Scope. El SistemaID se establece en 1 (aunque este valor
			puede necesitar ser dinámico dependiendo del sistema).
			 El Scope se establece utilizando el parámetro recibido en la función.
		*/
		var values struct {
			SistemaID int64  `json:"sistema_id"`
			Scope     string `json:"scope"`
		}
		//FIXME definir como vamos a setear el sistemaId
		values.SistemaID = 1
		values.Scope = scope

		// Se codifica la estructura values en formato JSON utilizando json.Marshal.
		json_data, _ := json.Marshal(values)

		/*
			Se crea una nueva solicitud HTTP POST utilizando la URL completa creada anteriormente
			y los datos codificados en JSON.
			Se establecen los encabezados Authorization y Content-Type en la solicitud.
		*/
		req, _ := http.NewRequest("POST", base.String(), bytes.NewBuffer(json_data))

		req.Header.Add("Authorization", bearer)
		req.Header.Add("Content-Type", "application/json")

		/*
			Envío de la solicitud a la API externa:
			Se realiza la solicitud HTTP utilizando el cliente HTTP almacenado en m.HTTPClient.
		*/
		resp, err := m.HTTPClient.Do(req)

		// Si hay un error durante el envío de la solicitud, se devuelve un error.
		if err != nil {
			return fmt.Errorf("error al enviar solicitud a api externa")
		}

		/*
			Manejo de la respuesta: Si la solicitud se realiza correctamente,
			se verifica el código de estado de la respuesta.
			Si el código de estado no es 200 (OK), se lee el cuerpo de la respuesta y
			se devuelve un error de fiber.NewError con un código de estado 403 (Forbidden)
			y un mensaje de error adecuado.
		*/
		if resp.StatusCode != 200 {
			info, _ := ioutil.ReadAll(req.Body)
			erro := fmt.Errorf("acceso denegado o permisos insuficientes: %s", info)
			return fiber.NewError(403, erro.Error())
		}

		//Si la solicitud es exitosa, se decodifica el cuerpo de la respuesta JSON en la estructura result.
		json.NewDecoder(resp.Body).Decode(&result)

		/*
			Almacenamiento del ID del usuario en el contexto de Fiber:
			Se extrae el ID del usuario de la estructura result y se almacena
			en el contexto de Fiber con la clave "user_id".
		*/
		c.Set("user_id", fmt.Sprint(result.ID))

		/*
			Continuación del flujo de middleware: Finalmente, se llama a c.Next() para
			continuar con el flujo del middleware y permitir que otros middleware o controladores
			manejen la solicitud.
		*/
		return c.Next()
	}
}
