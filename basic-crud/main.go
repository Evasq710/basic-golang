package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	// carga de driver de mysql
	_ "github.com/go-sql-driver/mysql"
)

type Empleado struct {
	Id int
	Nombre string
	Correo string
}

// en la variable se almacenarán todas las plantillas que se creen en la carpeta
var templates = template.Must(template.ParseGlob("templates/*"))

func main(){
	// ROUTES
	http.HandleFunc("/", home)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert_ACTION)
	http.HandleFunc("/delete", delete_ACTION)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/update", update_ACTION)
	log.Println("Running server...")

	// START THE SERVER
	http.ListenAndServe(":3000", nil)
}

func conectionDB()(conection *sql.DB){
	// datos de la database
	Driver := "mysql"
	User := "test_user"
	Password := "test_user"
	NombreDB := "sistema"

	conection, err := sql.Open(Driver, User + ":" + Password + "@tcp(127.0.0.1)/" + NombreDB)

	if err != nil {
		panic(err.Error())
	}

	return conection
}

func home(w http.ResponseWriter, req *http.Request){
	// http.ServeFile(w, req, "./inicio.html") // <-- para mostrar páginas html específicas

	establishedConection := conectionDB()
	rows, err := establishedConection.Query("SELECT * FROM Empleados;") // Query generalmente para SELECT, almacena las ROWS en var rows
	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	arrayEmpleados := []Empleado{}

	// recorriendo las ROWS de la DB
	for rows.Next() {
		var id int
		var nombre, email string
		err = rows.Scan(&id, &nombre, &email) // el número de destinos debe ser igual al número de columnas de la tabla
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = email
		arrayEmpleados = append(arrayEmpleados, empleado)
	}

	// mandamos el array al TEMPLATE, para recorrerlo con RANGE
	// VER templates/index.tmpl
	templates.ExecuteTemplate(w, "index", arrayEmpleados)
}

func create(w http.ResponseWriter, req *http.Request){
	templates.ExecuteTemplate(w, "create", nil)
}

func insert_ACTION(w http.ResponseWriter, req *http.Request){
	// en el req viene lo ingresado en el formulario
	// VER templates/create.tmpl, ya que aquí es donde se hace llamado a esta ruta

	if req.Method == "POST" {
		// id's de los inputs del form
		nombre := req.FormValue("nombre")
		correo := req.FormValue("correo")

		establishedConection := conectionDB()
		insert, err := establishedConection.Prepare("INSERT INTO Empleados(nombre, email) VALUES(?, ?)") // en los ? van los valores del req
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(nombre, correo) // los parámetros son los ? en la instrucción para mysql

		http.Redirect(w, req, "/", 301)
	}
}

func delete_ACTION(w http.ResponseWriter, req *http.Request){
	// VER templates/index.tmpl, ya que aquí es donde se hace llamado a esta ruta

	id := req.URL.Query().Get("idEmpleado") // aquí capturamos el id mandado en la URL como parámetro

	establishedConection := conectionDB()
	delete, err := establishedConection.Prepare("DELETE FROM Empleados WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)

	http.Redirect(w, req, "/", 301)
}

func edit(w http.ResponseWriter, req *http.Request){
	// En templates/index.tmpl se llama a esta ruta

	id := req.URL.Query().Get("idEmpleado")

	establishedConection := conectionDB()
	edit, err := establishedConection.Query("SELECT * FROM Empleados WHERE id=?", id)

	// obteniendo los datos del empleado a editar
	empleado := Empleado{}	
	for edit.Next() {
		var id int
		var nombre, email string
		err = edit.Scan(&id, &nombre, &email)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = email
	}

	// mandamos el empleado al TEMPLATE
	// VER templates/edit.tmpl
	templates.ExecuteTemplate(w, "edit", empleado)
}

func update_ACTION(w http.ResponseWriter, req *http.Request){
	// en el req viene lo ingresado en el formulario
	// VER templates/edit.tmpl, ya que aquí es donde se hace llamado a esta ruta

	if req.Method == "POST" {
		// id's de los inputs del form
		id := req.FormValue("idEmpleado")
		nombre := req.FormValue("nombre")
		correo := req.FormValue("correo")

		establishedConection := conectionDB()
		update, err := establishedConection.Prepare("UPDATE Empleados SET nombre=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		update.Exec(nombre, correo, id)

		http.Redirect(w, req, "/", 301)
	}
}