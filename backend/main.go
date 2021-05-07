package main

import (
	"math"
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

type EventoParametros struct {
	IdUser int `json:"user"`
}
type GananciasParametros struct {
	Year string `json:"year"`
}

type Grafica struct {
	X []string `json:"labels"`
	Y []int    `json:"data"`
}
type GraficaStack struct {
	Membership string `json:"label"`
	Count      [3]int `json:"data"`
	Stack      string `json:"stack"`
	Color      string `json:"backgroundColor"`
}
type arrayStack []GraficaStack

////////////////////////////

type Membresia struct {
	Id    string `json:"id"`
	Title string `json:"descripcion"`
	Price string `json:"precio"`
}
type arrayTier []Membresia

///////////////////////////
type EventoRetorno struct {
	Id      string `json:"id"`
	Home    string `json:"local"`
	Visit   string `json:"visita"`
	I_Date  string `json:"fecha_inicio"`
	F_Date  string `json:"fecha_final"`
	S_Home  string `json:"m_local"`
	S_Visit string `json:"m_visita"`
	Journey string `json:"jornada"`
	Sport   string `json:"deporte"`
	P_Home  string `json:"p_local"`
	P_Visit string `json:"p_visita"`
}

type arrayEventReturn []EventoRetorno

/////////////////////////
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

//////////////////////////////

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
			element.Name, element.Last, element.Pass, element.User, nil, element.User, nil)
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

					_, e := Database.Exec(`CALL CARGA_PREDICCION(:1,:2,:3,:4,:5,:6,:7,:8,:9,:10)`,
						element.Prediction.P_local, element.Prediction.P_visitant, user,
						journey, season, element.Local, element.Visit, element.Date, element.Sport,
						punteo(element.Result.R_local,
							element.Prediction.P_local,
							element.Result.R_visitant,
							element.Prediction.P_visitant))
					commitDB(e)
				}
			}
		}
	}
}
func punteo(resultado_local, prediccion_local, resultado_visita, prediccion_visita int) int {
	if resultado_local == prediccion_local && resultado_visita == prediccion_visita {
		return 10
	} else if (resultado_local >= resultado_visita && prediccion_local >= prediccion_visita) || (resultado_local <= resultado_visita && prediccion_local <= prediccion_visita) {
		results := float64(resultado_local + resultado_visita)
		predics := float64(prediccion_local + prediccion_visita)
		sum := math.Abs(results - predics)
		if sum <= 2 {
			return 5
		}
		return 3
	}
	return 0
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

func obtenerEventosAdmin(w http.ResponseWriter, r *http.Request) {
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
		Inicio := strings.Split(event.I_Date, "Z")
		event.I_Date = Inicio[0]
		Final := strings.Split(event.F_Date, "Z")
		event.F_Date = Final[0]
		lista = append(lista, event)
	}
	defer rows.Close()

	json.NewEncoder(w).Encode(lista)
}
func obtenerEventosUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var eventUser EventoParametros
	_ = json.NewDecoder(r.Body).Decode(&eventUser)
	fmt.Println("//////////////////////////////////")
	fmt.Println(strconv.Itoa(eventUser.IdUser))
	fmt.Println("//////////////////////////////////")
	var event EventoRetorno
	var lista = arrayEventReturn{}

	rows, err := Database.Query(" SELECT "+
		" EVENTO.ID_EVENTO, EVENTO.LOCAL, EVENTO.VISITANTE, "+
		" EVENTO.FECHA_INICIO, EVENTO.MARCADOR_LOCAL, EVENTO.MARCADOR_VISITA, "+
		" EVENTO.ID_JORNADA, EVENTO.ID_DEPORTE, EVENTO.FECHA_FIN, PREDICCION.LOCAL, PREDICCION.VISITANTE "+
		" FROM EVENTO "+
		" INNER JOIN PREDICCION "+
		" ON PREDICCION.ID_EVENTO = EVENTO.ID_EVENTO "+
		" INNER JOIN USUARIO "+
		" ON PREDICCION.ID_USUARIO = USUARIO.ID_USUARIO "+
		" WHERE USUARIO.ID_USUARIO = :1", strconv.Itoa(eventUser.IdUser))
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	for rows.Next() {
		fmt.Println(rows)

		rows.Scan(&event.Id, &event.Home, &event.Visit,
			&event.I_Date, &event.S_Home, &event.S_Visit,
			&event.Journey, &event.Sport, &event.F_Date, &event.P_Home, &event.P_Visit)
		Inicio := strings.Split(event.I_Date, "Z")
		event.I_Date = Inicio[0]
		Final := strings.Split(event.F_Date, "Z")
		event.F_Date = Final[0]
		lista = append(lista, event)
		fmt.Println("/////////////////////")
		fmt.Println(event.Id)
		fmt.Println(event.Home)
		fmt.Println(event.Visit)
		fmt.Println(event.I_Date)
		fmt.Println(event.S_Home)
		fmt.Println(event.S_Visit)
		fmt.Println(event.Journey)
		fmt.Println(event.Sport)
		fmt.Println(event.F_Date)
		fmt.Println(event.P_Home)
		fmt.Println(event.P_Visit)
	}
	defer rows.Close()

	json.NewEncoder(w).Encode(lista)
}
func obtenerOjiva(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datos Grafica
	var minimo int
	var maximo int
	var contador int

	Database.QueryRow(" SELECT SUM(MEMBRESIA.PRECIO) *0.20 " +
		" from DETALLE_MEMBRESIA " +
		" INNER JOIN TEMPORADA  " +
		" ON DETALLE_MEMBRESIA.ID_TEMPORADA =TEMPORADA.ID_TEMPORADA " +
		" INNER JOIN MEMBRESIA " +
		" ON DETALLE_MEMBRESIA.ID_MEMBRESIA=MEMBRESIA.ID_MEMBRESIA " +
		" GROUP BY TEMPORADA.NOMBRE " +
		" order by SUM(MEMBRESIA.PRECIO) ASC " +
		" FETCH FIRST 1 ROW ONLY ").Scan(&minimo)

	Database.QueryRow(" SELECT SUM(MEMBRESIA.PRECIO) *0.20 " +
		" from DETALLE_MEMBRESIA " +
		" INNER JOIN TEMPORADA  " +
		" ON DETALLE_MEMBRESIA.ID_TEMPORADA =TEMPORADA.ID_TEMPORADA " +
		" INNER JOIN MEMBRESIA " +
		" ON DETALLE_MEMBRESIA.ID_MEMBRESIA=MEMBRESIA.ID_MEMBRESIA " +
		" GROUP BY TEMPORADA.NOMBRE " +
		" order by SUM(MEMBRESIA.PRECIO) DESC " +
		" FETCH FIRST 1 ROW ONLY").Scan(&maximo)

	datos.X = append(datos.X, strconv.Itoa(0))
	datos.Y = append(datos.Y, 0)
	for minimo < maximo+250 {

		Database.QueryRow(` SELECT COUNT(MEMBERSHIP.GANANCIA) 
							from (
							SELECT SUM(MEMBRESIA.PRECIO)*0.20 GANANCIA, TEMPORADA.NOMBRE 
							FROM TEMPORADA 
							INNER JOIN DETALLE_MEMBRESIA 
							ON DETALLE_MEMBRESIA.ID_TEMPORADA = TEMPORADA.ID_TEMPORADA
							INNER JOIN MEMBRESIA 
							ON MEMBRESIA.ID_MEMBRESIA = DETALLE_MEMBRESIA.ID_MEMBRESIA
							GROUP BY TEMPORADA.NOMBRE
							) MEMBERSHIP
							WHERE GANANCIA <= :1`, minimo).Scan(&contador)
		datos.X = append(datos.X, strconv.Itoa(minimo))
		datos.Y = append(datos.Y, contador)
		minimo += 250
	}
	json.NewEncoder(w).Encode(datos)
}
func obtenerGanadores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datos [3]GraficaStack
	// var lista = arrayStack{}

	var ranking int
	var membresia string
	var conteo int

	rows, _ := Database.Query(`SELECT RANKING,MEMBERS,COUNT(*)
	FROM (
		SELECT MEMBERS,(ROW_NUMBER() OVER(PARTITION BY TEMPORADA ORDER BY PUNTOS DESC)) AS RANKING
			FROM(
					SELECT
						DISTINCT SUM(PREDICCION.PUNTAJE) PUNTOS,
						DETALLE_MEMBRESIA.ID_TEMPORADA TEMPORADA,
						MEMBRESIA.NOMBRE MEMBERS
	
					FROM
						PREDICCION
						INNER JOIN USUARIO ON PREDICCION.ID_USUARIO = USUARIO.ID_USUARIO
						INNER JOIN DETALLE_MEMBRESIA ON DETALLE_MEMBRESIA.ID_USUARIO = USUARIO.ID_USUARIO
						INNER JOIN JORNADA ON JORNADA.ID_TEMPORADA = DETALLE_MEMBRESIA.ID_TEMPORADA
						INNER JOIN EVENTO ON EVENTO.ID_EVENTO = PREDICCION.ID_PREDICCION
						AND EVENTO.ID_JORNADA = JORNADA.ID_JORNADA
						INNER JOIN MEMBRESIA ON MEMBRESIA.ID_MEMBRESIA = DETALLE_MEMBRESIA.ID_MEMBRESIA
					GROUP BY
						USUARIO.ID_USUARIO,
						DETALLE_MEMBRESIA.ID_TEMPORADA,
						MEMBRESIA.NOMBRE                
				)
		)
	GROUP BY RANKING,MEMBERS
	HAVING
		RANKING <= 3`)
	for rows.Next() {
		rows.Scan(&ranking, &membresia, &conteo)
		if ranking == 1 {
			if membresia == "bronze" {
				datos[0].Color="#d2691e"
				datos[0].Stack="Stack"
				datos[0].Count[0] = conteo
				datos[0].Membership = membresia
			} else if membresia == "silver" {
				datos[1].Color="#C0C0C0"
				datos[1].Stack="Stack"
				datos[1].Count[0] = conteo
				datos[1].Membership = membresia
			} else if membresia == "gold" {
				datos[2].Color="#ffd700"
				datos[2].Stack="Stack"
				datos[2].Count[0] = conteo
				datos[2].Membership = membresia
			}
		} else if ranking == 2 {
			if membresia == "bronze" {
				datos[0].Color="#d2691e"
				datos[0].Stack="Stack"
				datos[0].Count[1] = conteo
				datos[0].Membership = membresia
			} else if membresia == "silver" {
				datos[1].Color="#C0C0C0"
				datos[1].Stack="Stack"
				datos[1].Count[1] = conteo
				datos[1].Membership = membresia
			} else if membresia == "gold" {
				datos[2].Color="#ffd700"
				datos[2].Stack="Stack"
				datos[2].Count[1] = conteo
				datos[2].Membership = membresia
			}
		} else if ranking == 3 {
			if membresia == "bronze" {
				datos[0].Color="#d2691e"
				datos[0].Stack="Stack"
				datos[0].Count[2] = conteo
				datos[0].Membership = membresia
			} else if membresia == "silver" {
				datos[2].Color="#C0C0C0"
				datos[1].Stack="Stack"
				datos[1].Count[2] = conteo
				datos[1].Membership = membresia
			} else if membresia == "gold" {
				datos[2].Color="#ffd700"
				datos[2].Stack="Stack"
				datos[2].Count[2] = conteo
				datos[2].Membership = membresia
			}
		}

		// lista = append(lista, datos)
	}
	fmt.Println(datos)
	json.NewEncoder(w).Encode(datos)
}
func obtenerGanancias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datos Grafica

	var x string
	var y int

	rows, _ := Database.Query(`SELECT SUM(MEMBRESIA.PRECIO)*0.20 GANANCIA, TEMPORADA.NOMBRE 
	FROM TEMPORADA 
	INNER JOIN DETALLE_MEMBRESIA 
	ON DETALLE_MEMBRESIA.ID_TEMPORADA = TEMPORADA.ID_TEMPORADA
	INNER JOIN MEMBRESIA 
	ON MEMBRESIA.ID_MEMBRESIA = DETALLE_MEMBRESIA.ID_MEMBRESIA
	GROUP BY TEMPORADA.ANYO,TEMPORADA.MES,TEMPORADA.NOMBRE 
	ORDER BY TEMPORADA.ANYO,TEMPORADA.MES ASC`)
	for rows.Next() {
		rows.Scan(&y, &x)
		datos.X = append(datos.X, x)
		datos.Y = append(datos.Y, y)

	}
	json.NewEncoder(w).Encode(datos)
}

func obtenerGananciasYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var year GananciasParametros
	_ = json.NewDecoder(r.Body).Decode(&year)
	fmt.Println(year.Year)
	var datos Grafica
	var x string
	var y int

	rows, _ := Database.Query(`SELECT SUM(MEMBRESIA.PRECIO)*0.20 GANANCIA, TEMPORADA.NOMBRE 
	FROM TEMPORADA 
	INNER JOIN DETALLE_MEMBRESIA 
	ON DETALLE_MEMBRESIA.ID_TEMPORADA = TEMPORADA.ID_TEMPORADA
	INNER JOIN MEMBRESIA 
	ON MEMBRESIA.ID_MEMBRESIA = DETALLE_MEMBRESIA.ID_MEMBRESIA
	WHERE TEMPORADA.ANYO=:1
	GROUP BY TEMPORADA.ANYO,TEMPORADA.MES,TEMPORADA.NOMBRE 
	ORDER BY TEMPORADA.ANYO,TEMPORADA.MES ASC`, year.Year)
	for rows.Next() {
		rows.Scan(&y, &x)
		datos.X = append(datos.X, x)
		datos.Y = append(datos.Y, y)

	}
	json.NewEncoder(w).Encode(datos)
}
func obtenerMembresias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tier Membresia
	var lista = arrayTier{}
	rows, err := Database.Query("SELECT * FROM MEMBRESIA")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	for rows.Next() {

		rows.Scan(&tier.Id, &tier.Title, &tier.Price)
		lista = append(lista, tier)
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
	router.HandleFunc("/eventos/", obtenerEventosAdmin).Methods("GET")
	router.HandleFunc("/eventosUsuario/", obtenerEventosUser).Methods("POST")
	router.HandleFunc("/membresias/", obtenerMembresias).Methods("GET")
	router.HandleFunc("/ojiva/", obtenerOjiva).Methods("GET")
	router.HandleFunc("/ganancias/", obtenerGanancias).Methods("GET")
	router.HandleFunc("/gananciasY/", obtenerGananciasYear).Methods("POST")
	router.HandleFunc("/ganadores/", obtenerGanadores).Methods("GET")

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
