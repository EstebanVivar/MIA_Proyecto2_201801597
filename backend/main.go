package main

import (
	"math"
	"strconv"
	"strings"
	"time"

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
type GraficaPie struct {
	Data  [3]int    `json:"data"`
	Color [3]string `json:"backgroundColor"`
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

////////////////////////
type SetTier struct {
	Suscribed string `json:"subscription"`
	User      int    `json:"user"`
	Tier      string `json:"tier"`
}

//////////////////////////
type Temporada struct {
	Season string `json:"nombre"`
}
type arraySeason []Temporada

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

type ResultadoEvento struct {
	Id    string `json:"id"`
	Home  string `json:"r_local"`
	Visit string `json:"r_visita"`
}

type Posiciones struct {
	Id     string `json:"id"`
	User   string `json:"user"`
	Name   string `json:"name"`
	Last   string `json:"last"`
	Tier   string `json:"tier"`
	Season string `json:"season"`
	P10    string `json:"p10"`
	P5     string `json:"p5"`
	P3     string `json:"p3"`
	P0     string `json:"p0"`
	Total  string `json:"total"`
}
type arrayPosiciones []Posiciones

type NuevoEvento struct {
	Home    string `json:"local"`
	Visit   string `json:"visita"`
	I_Date  string `json:"fecha_inicio"`
	Journey int    `json:"jornada"`
	Sport   int    `json:"deporte"`
}

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
type NuevaPrediccion struct {
	User     int    `json:"user"`
	Event    string `json:"event"`
	P_visita string `json:"p_visita"`
	P_local  string `json:"p_local"`
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

func generarNombreTemporada(last, year string) string {
	ant := strings.Split(last, "-Q")

	// Temporada sigue en mismo anio solo aumentamos ID -Q#
	if year == ant[0] {
		id, _ := strconv.Atoi(ant[1])
		id++
		return ant[0] + "-Q" + strconv.Itoa(id)
	}

	return year + "-Q1"

}

func generarTemporada(w http.ResponseWriter, r *http.Request) {
	var ultimo string
	err := Database.QueryRow(`SELECT NOMBRE 
		FROM TEMPORADA ORDER BY FECHA_INICIO DESC 
		FETCH FIRST 1 ROWS ONLY`).Scan(&ultimo)
	commitDB(err)

	now := time.Now()
	year := now.Year()
	fecha := strings.Split(now.String(), ".") // Separador '.' para obtener formato 'YYYY-MM-DD HH:MM:SS'

	nombre := generarNombreTemporada(ultimo, strconv.Itoa(year)) // Generando el Nuevo ID de la temporada ->  2021-Q1 incremetnal
	fmt.Println(fecha[0])
	_, err1 := Database.Exec("CALL INSERT_NEW_SEASON(:1,:2)", nombre, fecha[0]) // Insertando la nueva temporada y creando su jornada 1
	commitDB(err1)
fmt.Println("'Q11111")
	rows, err2 := Database.Query(`SELECT to_char(DM.ID_PENDIENTE),DM.ID_USUARIO 
								FROM DETALLE_MEMBRESIA DM
								INNER JOIN TEMPORADA T 
								ON DM.ID_TEMPORADA = T.ID_TEMPORADA 
								WHERE T.NOMBRE = :1`, ultimo) // Obteniendo las membresias de los usuarios de la ultima teporada '2020-Q12' *TIER_PROXIMO ES EL ID
	commitDB(err2)
	defer rows.Close()
	var tier, user string
	for rows.Next() { // Recorriendo las membresias
		err := rows.Scan(&tier, &user)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err1 := Database.Exec("CALL RENOVAR_MEMBRESIA(:1,:2,:3)", tier, nombre, user) // por cada row se renueva la membresia para la jornada nueva 2021-Q1
		if err1 != nil {
			fmt.Println(err1)
			return
		}
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
			fmt.Println("1")
			_, err := Database.Exec(`CALL INSERT_SEASON(:1,:2 )`,
				element.Season, arrayResultados[0]+"-"+arrayResultados[1])
			commitDB(err)
			season = element.Season
			tier = element.Tier
			fmt.Println("2")
			_, er := Database.Exec(`CALL INSERT_SEASON_DETAIL(:1,:2, :3)`,
				user, season, tier)
			commitDB(er)

			fmt.Println("Jornadas")
			for _, element := range element.Journeys {
				journey = element.Journey
				arrayJornada := strings.Split(element.Journey, "J")
				fmt.Println("3")
				_, er := Database.Exec(`CALL INSERT_JOURNEY_CARGA(:1,:2,:3,:4)`,
					element.Journey, season, arrayResultados[0]+"-"+arrayResultados[1], arrayJornada[1])
				commitDB(er)

				fmt.Println("Predicciones")
				for _, element := range element.Predictions {
					fmt.Println("4")
					_, err := Database.Exec(`CALL INSERT_SPORT(:1,:2,:3)`,
						element.Sport, nil, nil)
					commitDB(err)
					fmt.Println("7")

					_, er := Database.Exec(`CALL INSERT_EVENT_CARGA(:1,:2,:3,:4,:5,:6,:7,:8)`,
						element.Local, element.Visit, element.Date, element.Result.R_local,
						element.Result.R_visitant, journey, season, element.Sport)
					commitDB(er)
					fmt.Println("6")

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

func punteo(r_local, p_local, r_visita, p_visita int) int {
	if r_local == p_local && r_visita == p_visita {
		return 10
	} else if (r_local >= r_visita && p_local >= p_visita) || (r_local <= r_visita && p_local <= p_visita) {
		if (r_local == r_visita && p_local != p_visita) || (r_local != r_visita && p_local == p_visita) {
			return 0
		}
		results := float64(r_local + r_visita)
		predics := float64(p_local + p_visita)
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

func registrarEvento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var evento NuevoEvento

	_ = json.NewDecoder(r.Body).Decode(&evento)

	fmt.Println(evento.Home)
	fmt.Println(evento.Visit)
	fmt.Println(evento.I_Date)
	fmt.Println(evento.Journey)
	fmt.Println(evento.Sport)
	_, err := Database.Query("CALL INSERT_EVENT(:1,:2,:3,:4,:5)",
		evento.Home, evento.Visit, evento.I_Date, evento.Journey, evento.Sport)
	commitDB(err)
	json.NewEncoder(w).Encode(evento)
}

func registrarPrediccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prediccion NuevaPrediccion

	_ = json.NewDecoder(r.Body).Decode(&prediccion)
	fmt.Println(prediccion)
	// fmt.Println(evento.Home)

	_, err := Database.Query(` INSERT INTO PREDICCION(LOCAL, VISITANTE, ID_EVENTO, ID_USUARIO)
    VALUES (:1,:2,:3,:4)`, prediccion.P_local, prediccion.P_visita, prediccion.Event, prediccion.User)
	commitDB(err)
	json.NewEncoder(w).Encode(prediccion)
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

func actualizarEvento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resultados ResultadoEvento

	_ = json.NewDecoder(r.Body).Decode(&resultados)
	fmt.Println(resultados.Id, resultados.Home, resultados.Visit)
	rows, err := Database.Query(`UPDATE EVENTO 
							SET MARCADOR_LOCAL=:1,
							MARCADOR_VISITA=:2
							WHERE ID_EVENTO = :3`,
		resultados.Home, resultados.Visit, resultados.Id)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	} else {
		Database.Exec("COMMIT;")
	}
	defer rows.Close()
	json.NewEncoder(w).Encode(resultados)
}
func actualizarMembresia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lastSeason string
	err := Database.QueryRow(`SELECT ID_TEMPORADA 
		FROM TEMPORADA ORDER BY FECHA_INICIO DESC 
		FETCH FIRST 1 ROWS ONLY`).Scan(&lastSeason)
	commitDB(err)
	var cambioMember SetTier
	_ = json.NewDecoder(r.Body).Decode(&cambioMember)
	fmt.Println(cambioMember)
	var cantidadMembresias int
	err2 := Database.QueryRow(`
	SELECT count(*) FROM DETALLE_MEMBRESIA where ID_USUARIO = :1`, cambioMember.User).Scan(&cantidadMembresias)
	commitDB(err2)
	fmt.Println(cantidadMembresias)
	if cantidadMembresias == 0 {
		_, err1 := Database.Exec("CALL BUY_MEMBERSHIP(:1,:2,:3,'Y')", cambioMember.Tier, lastSeason, cambioMember.User) // por cada row se renueva la membresia para la jornada nueva 2021-Q1
		if err1 != nil {
			fmt.Println(err1)
			return
		} else {
			Database.Exec("COMMIT;")
		}

	} else {

		rows, err := Database.Query(`UPDATE DETALLE_MEMBRESIA 
							SET ID_PENDIENTE=:1,
							SUSCRIPCION=:2
							WHERE ID_USUARIO = :3
							AND ID_TEMPORADA =:4`,
			cambioMember.Tier, cambioMember.Suscribed, cambioMember.User, lastSeason)
		if err != nil {
			fmt.Println("Error running query")
			fmt.Println(err)
			return
		} else {
			Database.Exec("COMMIT;")
		}
		defer rows.Close()
	}
	json.NewEncoder(w).Encode(cambioMember)
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
	rows, err := Database.Query(`SELECT 
									E.ID_EVENTO,
									E.LOCAL,
									E.VISITANTE,
									E.FECHA_INICIO,
									TO_CHAR(E.MARCADOR_LOCAL),
									TO_CHAR(E.MARCADOR_VISITA),
									E.ID_JORNADA,
									E.ID_DEPORTE,
									E.FECHA_FIN
							 	FROM EVENTO E`)
	commitDB(err)

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

	var event EventoRetorno
	var lista = arrayEventReturn{}

	rows, err := Database.Query(`SELECT
									E.ID_EVENTO,
									E.LOCAL,
									E.VISITANTE,
									E.FECHA_INICIO,
									NVL(to_char(E.MARCADOR_LOCAL),'/'),
									NVL(to_char(E.MARCADOR_VISITA),'/'),
									E.ID_JORNADA,
									E.ID_DEPORTE,
									E.FECHA_FIN,
									to_char(P.LOCAL),
									to_char(P.VISITANTE)
									
								FROM
									EVENTO E
									LEFT JOIN PREDICCION P ON E.ID_EVENTO = P.ID_EVENTO    
									AND P.ID_USUARIO = :1`,
		strconv.Itoa(eventUser.IdUser))
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	for rows.Next() {

		rows.Scan(&event.Id, &event.Home, &event.Visit,
			&event.I_Date, &event.S_Home, &event.S_Visit,
			&event.Journey, &event.Sport, &event.F_Date, &event.P_Home, &event.P_Visit)
		Inicio := strings.Split(event.I_Date, "Z")
		event.I_Date = Inicio[0]
		Final := strings.Split(event.F_Date, "Z")
		event.F_Date = Final[0]
		lista = append(lista, event)

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

func obtenerPerdedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var season Temporada
	_ = json.NewDecoder(r.Body).Decode(&season)
	var membresia string
	var porcentaje int
	var pastel GraficaPie
	rows, _ := Database.Query(`SELECT
								MEMBERS,
								COUNT(*)
								FROM
									(
										SELECT
											MEMBERS,(
												DENSE_RANK() OVER(
													PARTITION BY TEMPORADA
													ORDER BY
														PUNTOS DESC
												)
											) AS RANKING
										FROM(
												SELECT
													SUM(P.PUNTAJE) PUNTOS,
													T.NOMBRE TEMPORADA,
													M.NOMBRE MEMBERS
												FROM
													PREDICCION P
													INNER JOIN USUARIO U ON P.ID_USUARIO = U.ID_USUARIO
													INNER JOIN DETALLE_MEMBRESIA DM ON DM.ID_USUARIO = U.ID_USUARIO
													INNER JOIN TEMPORADA T ON T.ID_TEMPORADA = DM.ID_TEMPORADA
													INNER JOIN JORNADA J ON J.ID_TEMPORADA = T.ID_TEMPORADA
													INNER JOIN EVENTO E ON E.ID_EVENTO = P.ID_EVENTO
													AND E.ID_JORNADA = J.ID_JORNADA
													INNER JOIN MEMBRESIA M ON M.ID_MEMBRESIA = DM.ID_MEMBRESIA
												GROUP BY
													U.ID_USUARIO,
													M.NOMBRE,
													T.NOMBRE
													HAVING T.NOMBRE=:1
											)
									)
									WHERE
									RANKING > 3
								GROUP BY
								MEMBERS`, season.Season)
	for rows.Next() {
		rows.Scan(&membresia, &porcentaje)

		if membresia == "bronze" {
			pastel.Color[0] = "#D2691E"
			pastel.Data[0] = porcentaje
		} else if membresia == "silver" {
			pastel.Color[1] = "#C0C0C0"
			pastel.Data[1] = porcentaje
		} else if membresia == "gold" {
			pastel.Color[2] = "#FFD700"
			pastel.Data[2] = porcentaje
		}

	}
	fmt.Println(pastel)
	json.NewEncoder(w).Encode(pastel)
}

func obtenerPosiciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var season Temporada
	var posicion Posiciones
	var lista = arrayPosiciones{}
	_ = json.NewDecoder(r.Body).Decode(&season)
	fmt.Println(season.Season)
	rows, _ := Database.Query(`SELECT
								U.ID_USUARIO,
								U.USUARIO,
								U.NOMBRE,
								U.APELLIDO,
								MEMBRESIA.NOMBRE,
								DM.ID_TEMPORADA,
								(
									SELECT
										COUNT(*)
									FROM
									PREDICCION P
									INNER JOIN EVENTO E
									ON P.ID_EVENTO=E.ID_EVENTO
									INNER JOIN JORNADA J
									ON J.ID_JORNADA = E.ID_JORNADA
									INNER JOIN TEMPORADA T
									ON J.ID_TEMPORADA =T.ID_TEMPORADA
									
									WHERE
										P.PUNTAJE = 10
										AND P.ID_USUARIO = U.ID_USUARIO
										AND T.NOMBRE = :1
								) AS P10,
								(
									SELECT
										COUNT(*)
									FROM
									PREDICCION P
									INNER JOIN EVENTO E
									ON P.ID_EVENTO=E.ID_EVENTO
									INNER JOIN JORNADA J
									ON J.ID_JORNADA = E.ID_JORNADA
									INNER JOIN TEMPORADA T
									ON J.ID_TEMPORADA =T.ID_TEMPORADA
									WHERE
										P.PUNTAJE = 5
										AND P.ID_USUARIO = U.ID_USUARIO
										AND T.NOMBRE = :1
								) AS P5,
								(
									SELECT
										COUNT(*)
									FROM
									PREDICCION P
									INNER JOIN EVENTO E
									ON P.ID_EVENTO=E.ID_EVENTO
									INNER JOIN JORNADA J
									ON J.ID_JORNADA = E.ID_JORNADA
									INNER JOIN TEMPORADA T
									ON J.ID_TEMPORADA =T.ID_TEMPORADA
									WHERE
										P.PUNTAJE = 3
										AND P.ID_USUARIO = U.ID_USUARIO
										AND T.NOMBRE = :1
								) AS P3,
								(
									SELECT
										COUNT(*)
									FROM
									PREDICCION P
									INNER JOIN EVENTO E
									ON P.ID_EVENTO=E.ID_EVENTO
									INNER JOIN JORNADA J
									ON J.ID_JORNADA = E.ID_JORNADA
									INNER JOIN TEMPORADA T
									ON J.ID_TEMPORADA =T.ID_TEMPORADA
									WHERE
										P.PUNTAJE = 0
										AND P.ID_USUARIO = U.ID_USUARIO
										AND T.NOMBRE = :1
								) AS P0,
								(
									SELECT
										SUM(P.PUNTAJE)
									FROM
									PREDICCION P
									INNER JOIN EVENTO E
									ON P.ID_EVENTO=E.ID_EVENTO
									INNER JOIN JORNADA J
									ON J.ID_JORNADA = E.ID_JORNADA
									INNER JOIN TEMPORADA T
									ON J.ID_TEMPORADA =T.ID_TEMPORADA
									WHERE
										P.ID_USUARIO = U.ID_USUARIO
										AND T.NOMBRE = :1
									GROUP BY
										ID_USUARIO
								) as TOTAL
							FROM
								PREDICCION P
								INNER JOIN USUARIO U ON P.ID_USUARIO = U.ID_USUARIO
								INNER JOIN DETALLE_MEMBRESIA DM ON U.ID_USUARIO = DM.ID_USUARIO
								INNER JOIN MEMBRESIA ON DM.ID_MEMBRESIA = MEMBRESIA.ID_MEMBRESIA
								INNER JOIN TEMPORADA T ON DM.ID_TEMPORADA = T.ID_TEMPORADA
							WHERE
								T.NOMBRE = :1
								AND DM.ID_USUARIO = U.ID_USUARIO
							GROUP BY
								U.ID_USUARIO,
								U.USUARIO,
								U.NOMBRE,
								U.APELLIDO,
								MEMBRESIA.NOMBRE,
								DM.ID_TEMPORADA   

							ORDER BY
								TOTAL DESC`, season.Season)

	for rows.Next() {
		rows.Scan(&posicion.Id, &posicion.User, &posicion.Name, &posicion.Last, &posicion.Tier,
			&posicion.Season, &posicion.P10, &posicion.P5, &posicion.P3, &posicion.P0, &posicion.Total)

		lista = append(lista, posicion)

	}
	json.NewEncoder(w).Encode(lista)
}

func obtenerGanadores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datos [3]GraficaStack

	var ranking int
	var membresia string
	var conteo int

	rows, _ := Database.Query(`SELECT
									RANKING,
									MEMBERS,
									COUNT(*)
								FROM
									(
										SELECT
											MEMBERS,(
												DENSE_RANK() OVER(
													PARTITION BY TEMPORADA
													ORDER BY
														PUNTOS DESC
												)
											) AS RANKING
										FROM(
												SELECT
													SUM(P.PUNTAJE) PUNTOS,
													T.ID_TEMPORADA TEMPORADA,
													M.NOMBRE MEMBERS
												FROM
													PREDICCION P
													INNER JOIN USUARIO U ON P.ID_USUARIO = U.ID_USUARIO
													INNER JOIN DETALLE_MEMBRESIA DM ON DM.ID_USUARIO = U.ID_USUARIO
													INNER JOIN TEMPORADA T ON T.ID_TEMPORADA = DM.ID_TEMPORADA
													INNER JOIN JORNADA J ON J.ID_TEMPORADA = T.ID_TEMPORADA
													INNER JOIN EVENTO E ON E.ID_EVENTO = P.ID_EVENTO
													AND E.ID_JORNADA = J.ID_JORNADA
													INNER JOIN MEMBRESIA M ON M.ID_MEMBRESIA = DM.ID_MEMBRESIA
												GROUP BY
													U.ID_USUARIO,
													M.NOMBRE,
													T.ID_TEMPORADA
											)
									)
								GROUP BY
									RANKING,
									MEMBERS
								HAVING
									RANKING <= 3
								ORDER BY
									RANKING ASC`)
	for rows.Next() {
		rows.Scan(&ranking, &membresia, &conteo)
		fmt.Println(ranking, membresia, conteo)
		if ranking == 1 {
			if membresia == "bronze" {
				datos[0].Color = "#D2691E"
				datos[0].Stack = "Stack"
				datos[0].Count[0] = conteo
				datos[0].Membership = membresia
			} else if membresia == "silver" {
				datos[1].Color = "#C0C0C0"
				datos[1].Stack = "Stack"
				datos[1].Count[0] = conteo
				datos[1].Membership = membresia
			} else if membresia == "gold" {
				datos[2].Color = "#FFD700"
				datos[2].Stack = "Stack"
				datos[2].Count[0] = conteo
				datos[2].Membership = membresia
			}
		} else if ranking == 2 {
			if membresia == "bronze" {
				datos[0].Color = "#D2691E"
				datos[0].Stack = "Stack"
				datos[0].Count[1] = conteo
				datos[0].Membership = membresia
			} else if membresia == "silver" {
				datos[1].Color = "#C0C0C0"
				datos[1].Stack = "Stack"
				datos[1].Count[1] = conteo
				datos[1].Membership = membresia
			} else if membresia == "gold" {
				datos[2].Color = "#FFD700"
				datos[2].Stack = "Stack"
				datos[2].Count[1] = conteo
				datos[2].Membership = membresia
			}
		} else if ranking == 3 {
			if membresia == "bronze" {
				datos[0].Color = "#D2691E"
				datos[0].Stack = "Stack"
				datos[0].Count[2] = conteo
				datos[0].Membership = membresia
			} else if membresia == "silver" {
				datos[1].Color = "#C0C0C0"
				datos[1].Stack = "Stack"
				datos[1].Count[2] = conteo
				datos[1].Membership = membresia
			} else if membresia == "gold" {
				datos[2].Color = "#FFD700"
				datos[2].Stack = "Stack"
				datos[2].Count[2] = conteo
				datos[2].Membership = membresia
			}
		}
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
	GROUP BY TEMPORADA.FECHA_INICIO,TEMPORADA.NOMBRE 
	ORDER BY TEMPORADA.FECHA_INICIO ASC`)
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
	var datos Grafica
	var x string
	var y int

	rows, _ := Database.Query(`SELECT SUM(MEMBRESIA.PRECIO)*0.20 GANANCIA, TEMPORADA.NOMBRE 
	FROM TEMPORADA 
	INNER JOIN DETALLE_MEMBRESIA 
	ON DETALLE_MEMBRESIA.ID_TEMPORADA = TEMPORADA.ID_TEMPORADA
	INNER JOIN MEMBRESIA 
	ON MEMBRESIA.ID_MEMBRESIA = DETALLE_MEMBRESIA.ID_MEMBRESIA
	WHERE EXTRACT(YEAR FROM TEMPORADA.FECHA_INICIO)=:1
	GROUP BY TEMPORADA.FECHA_INICIO,TEMPORADA.NOMBRE 
	ORDER BY TEMPORADA.FECHA_INICIO ASC`, year.Year)
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

func obtenerTemporadas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var season Temporada
	var lista = arraySeason{}
	rows, err := Database.Query("SELECT NOMBRE FROM TEMPORADA ORDER BY  FECHA_INICIO ASC")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	for rows.Next() {

		rows.Scan(&season.Season)
		lista = append(lista, season)
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
	router.HandleFunc("/perdedores/", obtenerPerdedores).Methods("POST")
	router.HandleFunc("/temporadas/", obtenerTemporadas).Methods("GET")
	router.HandleFunc("/u_results/", actualizarEvento).Methods("POST")
	router.HandleFunc("/crear_evento/", registrarEvento).Methods("POST")
	router.HandleFunc("/NewPred/", registrarPrediccion).Methods("POST")
	router.HandleFunc("/cambioMember/", actualizarMembresia).Methods("POST")
	router.HandleFunc("/tablaPosiciones/", obtenerPosiciones).Methods("POST")
	router.HandleFunc("/newSeason/", generarTemporada).Methods("GET")
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
