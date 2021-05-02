package main

import (
	"strconv"
	"strings"

	// "io/ioutil"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var Database *sql.DB

type Evento struct {
	Id      string `json:"id"`
	Home    string `json:"local"`
	Visit   string `json:"visita"`
	I_Date  string `json:"fecha_inicio"`
	F_Date  string `json:"fecha_final"`
	S_Home  string `json:"m_local"`
	S_Visit string `json:"m_visita"`
	Journey string `json:"jornada"`
	Sport   string `json:"deporte"`
}

type arrayEvent []Evento

type Usuario struct {
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Name  string `json:"name"`
	Last  string `json:"last"`
	Birth string `json:"birth"`
	Email string `json:"email"`
	Photo string `json:"photo"`
}
type UsuarioUpdate struct {
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Name  string `json:"name"`
	Last  string `json:"last"`
	Birth string `json:"birth"`
	Email string `json:"email"`
	Photo string `json:"photo"`
	Id    int    `json:"id"`
}

type UsuarioLogin struct {
	Id       int    `json:"id"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Name     string `json:"name"`
	Last     string `json:"last"`
	Birth    string `json:"birth"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
	Register string `json:"register"`
	Admin    int    `json:"admin"`
}

type Login struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

////////////////////////////////
type prediction struct {
	P_visitant int `json:"visitante"`
	P_local    int `json:"local"`
}

type result struct {
	R_visitant int `json:"visitante"`
	R_local    int `json:"local"`
}

type predictions struct {
	Sport      string     `json:"deporte"`
	Date       string     `json:"fecha"`
	Visit      string     `json:"visitante"`
	Local      string     `json:"local"`
	Prediction prediction `json:"prediccion"`
	Result     result     `json:"resultado"`
}

type journeys struct {
	Journey     string        `json:"jornada"`
	Predictions []predictions `json:"predicciones"`
}

type results struct {
	Season   string     `json:"temporada"`
	Tier     string     `json:"tier"`
	Journeys []journeys `json:"jornadas"`
}
type user struct {
	User    string    `json:"username"`
	Pass    string    `json:"password"`
	Name    string    `json:"nombre"`
	Last    string    `json:"apellido"`
	Results []results `json:"resultados"`
}

type Info map[string]user

/////////////////////////

func loadTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Carga Info
	json.NewDecoder(r.Body).Decode(&Carga)

	for key, element := range Carga {
		fmt.Println("ID:", key)
		fmt.Print("\t")
		fmt.Println("Usuario:", element.User)
		fmt.Print("\t")
		fmt.Println("Clave:", element.Pass)
		fmt.Print("\t")
		fmt.Println("Nombre:", element.Name)
		fmt.Print("\t")
		fmt.Println("Apellido:", element.Last)
		for _, element := range element.Results {
			fmt.Print("\t")
			fmt.Println("Resultados:")
			fmt.Print("\t\t")
			fmt.Println("Temporada:", element.Season)
			fmt.Print("\t\t")
			fmt.Println("Membresia:", element.Tier)
			for _, element := range element.Journeys {
				fmt.Print("\t\t")
				fmt.Println("Jornadas:")
				fmt.Print("\t\t\t")
				fmt.Println("Jornada:", element.Journey)
				for _, element := range element.Predictions {
					fmt.Print("\t\t\t")
					fmt.Println("Predicciones:")
					fmt.Print("\t\t\t\t")
					fmt.Println("Deporte:", element.Sport)
					fmt.Print("\t\t\t\t")
					fmt.Println("Local:", element.Local)
					fmt.Print("\t\t\t\t")
					fmt.Println("Visitante:", element.Visit)
					fmt.Print("\t\t\t\t")
					fmt.Println("Fecha:", element.Date)
					fmt.Print("\t\t\t\t")
					fmt.Println("Prediccion:")
					fmt.Print("\t\t\t\t\t")
					fmt.Println("P Local:", element.Result.R_local)
					fmt.Print("\t\t\t\t\t")
					fmt.Println("P Visita:", element.Result.R_visitant)
					fmt.Print("\t\t\t\t")
					fmt.Println("Resultado:")
					fmt.Print("\t\t\t\t\t")
					fmt.Println("R Local:", element.Prediction.P_local)
					fmt.Print("\t\t\t\t\t")
					fmt.Println("R Visita:", element.Prediction.P_visitant)

				}
			}
		}
	}
}
func commitDB(err error) {
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err.Error())
	} else {

		Database.Exec("COMMIT;")
	}
}
func load(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user string
	var season string
	var tier string
	var journey string

	var Carga Info
	json.NewDecoder(r.Body).Decode(&Carga)

	fmt.Println("Usuarios")
	for _, element := range Carga {
		_, err := Database.Exec(`CALL INSERT_USER(:1,:2, :3, :4,:5, :6, :7)`,
			element.Name, element.Last, element.Pass, element.User, nil, nil, nil)
		commitDB(err)
		user = element.User

		fmt.Println("Resultados")
		for _, element := range element.Results {
			arrayResultados := strings.Split(element.Season, "-Q")
			anyo, _ := strconv.Atoi(arrayResultados[0])
			mes, _ := strconv.Atoi(arrayResultados[1])

			_, err := Database.Exec(`CALL INSERT_SEASON(:1,:2, :3, :4)`,
				element.Season, anyo, mes, arrayResultados[0]+"-"+arrayResultados[1])
			commitDB(err)
			season = element.Season
			tier = element.Tier

			_, er := Database.Exec(`CALL INSERT_SEASON_DETAIL(:1,:2, :3)`,
				user, season, tier)
			commitDB(er)

			fmt.Println("Jornadas")
			for _, element := range element.Journeys {
				journey = element.Journey
				arrayJornada := strings.Split(element.Journey, "J")
				_, er := Database.Exec(`CALL INSERT_JOURNEY(:1,:2,:3,:4)`,
					element.Journey, season, arrayResultados[0]+"-"+arrayResultados[1], arrayJornada[1])
				commitDB(er)

				fmt.Println("Predicciones")
				for _, element := range element.Predictions {
					_, err := Database.Exec(`CALL INSERT_SPORT(:1,:2,:3)`,
						element.Sport, nil, nil)
					commitDB(err)

					_, er := Database.Exec(`CALL INSERT_EVENT(:1,:2,:3,:4,:5,:6,:7,:8)`,
						element.Local, element.Visit, element.Date, element.Result.R_local,
						element.Result.R_visitant, journey, season, element.Sport)
					commitDB(er)

					_, e := Database.Exec(`CALL CARGA_PREDICCION(:1,:2,:3,:4,:5,:6,:7)`,
						element.Prediction.P_local, element.Prediction.P_visitant, user,
						journey, season, element.Local, element.Visit)
					commitDB(e)
				}
			}
		}
	}
}
func registrarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usuario Usuario
	_ = json.NewDecoder(r.Body).Decode(&usuario)
	_, err := Database.Exec(`CALL INSERT_USER(:1,:2,:3,:4,:5,:6,:7)`,
		usuario.Name, usuario.Last, usuario.Pass, usuario.User, usuario.Birth, usuario.Email, usuario.Photo)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	} else {
		Database.Exec("COMMIT;")
	}
	json.NewEncoder(w).Encode(usuario)
}

func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usuario UsuarioUpdate
	_ = json.NewDecoder(r.Body).Decode(&usuario)
	_, err := Database.Exec(`CALL UPDATE_USER(:1,:2,:3,:4,:5,:6,:7,:8)`,
		usuario.Id, usuario.Name, usuario.Last, usuario.Pass, usuario.User, usuario.Birth, usuario.Email, usuario.Photo)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	} else {
		Database.Exec("COMMIT;")
	}
	json.NewEncoder(w).Encode(usuario)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var login Login
	_ = json.NewDecoder(r.Body).Decode(&login)
	var isUser int
	var isAd int
	var idUser int
	var User UsuarioLogin

	_, err := Database.Exec("CALL LOGIN(:1,:2,:3,:4,:5)", login.User, login.Pass, sql.Out{Dest: &isUser},
		sql.Out{Dest: &isAd}, sql.Out{Dest: &idUser})
	fmt.Println(login)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	} else {
		if isUser == 1 {

			Database.QueryRow("SELECT * FROM USUARIO WHERE ID_USUARIO = :1", idUser).Scan(&User.Id, &User.Name,
				&User.Last, &User.Pass, &User.User, &User.Birth,
				&User.Register, &User.Email, &User.Photo)
			fmt.Println(&User.Id, &User.Name,
				&User.Last, &User.Pass, &User.User, &User.Birth,
				&User.Register, &User.Email, &User.Photo)
			if isAd == 1 {
				User.Admin = 1
			} else {
				User.Admin = 0
			}
		}
	}
	fix := strings.Split(User.Birth, "T")
	User.Birth = fix[0]
	json.NewEncoder(w).Encode(User)
}

func obtenerEventos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event Evento
	var lista = arrayEvent{}
	rows, err := Database.Query("SELECT * FROM EVENTO")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	for rows.Next() {

		rows.Scan(&event.Id, &event.Home, &event.Visit,
			&event.I_Date, &event.S_Home, &event.S_Visit,
			&event.Journey, &event.Sport, &event.F_Date)

		fix := strings.Split(event.I_Date, "T")
		event.I_Date = fix[0]
		fix2 := strings.Split(event.F_Date, "T")
		event.F_Date = fix2[0]

		lista = append(lista, event)
	}
	defer rows.Close()

	json.NewEncoder(w).Encode(lista)
}

func main() {
	// routes
	router := mux.NewRouter()
	router.HandleFunc("/registrar/", registrarUsuario).Methods("POST")
	router.HandleFunc("/login/", login).Methods("POST")
	router.HandleFunc("/actualizar/", actualizarUsuario).Methods("POST")
	router.HandleFunc("/test/", load).Methods("POST")
	router.HandleFunc("/eventos/", obtenerEventos).Methods("GET")

	db, err := sql.Open("godror", "admin/admin@localhost:1521/ORCLCDB.localdomain")
	Database = db
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// server
	port, defport := os.LookupEnv("GOPORT")

	if !defport {
		port = "4000"
	}
	fmt.Println("Listen on port " + port)
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":"+port, handler)

}
