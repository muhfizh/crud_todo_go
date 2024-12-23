package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type ListToDo struct {
	id          int
	Judul_Task  string
	Jenis_Task  string
	Desc_Task   string
	Tgl_Task    string
	Tgl_Selesai string
	Total_Jam   int
	Status_Task string
	Created_At  string
	Create_Date string
	Update_At   string
	Update_Date string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "listtodo"
	db, err := sql.Open(dbDriver, dbUser+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM list_todo_tabs ORDER BY id DESC")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ltd := ListToDo{}
	res := []ListToDo{}
	for selDB.Next() {
		var id, total_jam int
		var judul_task, jenis_task, desc_task, status_task, created_at, update_at, tgl_task, tgl_selesai, create_date, update_date string
		err = selDB.Scan(&id,
			&judul_task,
			&jenis_task,
			&desc_task,
			&tgl_task,
			&tgl_selesai,
			&total_jam,
			&status_task,
			&created_at,
			&create_date,
			&update_at,
			&update_date)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ltd.id = id
		ltd.Judul_Task = judul_task
		ltd.Jenis_Task = jenis_task
		ltd.Desc_Task = desc_task
		ltd.Tgl_Task = tgl_task
		ltd.Tgl_Selesai = tgl_selesai
		ltd.Total_Jam = total_jam
		ltd.Status_Task = status_task
		ltd.Created_At = created_at
		ltd.Create_Date = create_date
		ltd.Update_At = update_at
		ltd.Update_Date = update_date
		res = append(res, ltd)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM list_todo_tabs WHERE id=?", nId)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ltd := ListToDo{}
	for selDB.Next() {
		var id, total_jam int
		var judul_task, jenis_task, desc_task, status_task, created_at, update_at, tgl_task, tgl_selesai, create_date, update_date string
		err = selDB.Scan(&id,
			&judul_task,
			&jenis_task,
			&desc_task,
			&tgl_task,
			&tgl_selesai,
			&total_jam,
			&status_task,
			&created_at,
			&create_date,
			&update_at,
			&update_date)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ltd.id = id
		ltd.Judul_Task = judul_task
		ltd.Jenis_Task = jenis_task
		ltd.Desc_Task = desc_task
		ltd.Tgl_Task = tgl_task
		ltd.Tgl_Selesai = tgl_selesai
		ltd.Total_Jam = total_jam
		ltd.Status_Task = status_task
		ltd.Created_At = created_at
		ltd.Create_Date = create_date
		ltd.Update_At = update_at
		ltd.Update_Date = update_date
	}
	tmpl.ExecuteTemplate(w, "Show", ltd)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM list_todo_tabs WHERE id=?", nId)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ltd := ListToDo{}
	for selDB.Next() {
		var id, total_jam int
		var judul_task, jenis_task, desc_task, status_task, created_at, update_at, tgl_task, tgl_selesai, create_date, update_date string
		err = selDB.Scan(&id,
			&judul_task,
			&jenis_task,
			&desc_task,
			&tgl_task,
			&tgl_selesai,
			&total_jam,
			&status_task,
			&created_at,
			&create_date,
			&update_at,
			&update_date)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ltd.id = id
		ltd.Judul_Task = judul_task
		ltd.Jenis_Task = jenis_task
		ltd.Desc_Task = desc_task
		ltd.Tgl_Task = tgl_task
		ltd.Tgl_Selesai = tgl_selesai
		ltd.Total_Jam = total_jam
		ltd.Status_Task = status_task
		ltd.Created_At = created_at
		ltd.Create_Date = create_date
		ltd.Update_At = update_at
		ltd.Update_Date = update_date
	}
	tmpl.ExecuteTemplate(w, "Edit", ltd)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("id")
		Judul_Task := r.FormValue("Judul_Task")
		Jenis_Task := r.FormValue("Jenis_Task")
		Desc_Task := r.FormValue("Desc_Task")
		Tgl_Task := r.FormValue("Tgl_Task")
		Tgl_Selesai := r.FormValue("Tgl_Selesai")
		Total_Jam := r.FormValue("Total_Jam")
		Status_Task := r.FormValue("Status_Task")
		Created_At := r.FormValue("Created_At")
		Create_Date := r.FormValue("Create_Date")
		Update_At := r.FormValue("Update_At")
		Update_Date := r.FormValue("Update_Date")
		insForm, err := db.Prepare("INSERT INTO list_todo_tabs(id, Judul_Task,Jenis_Task,Desc_Task,Tgl_Task,Tgl_Selesai,Total_Jam,Status_Task,Created_At,Create_Date,Update_At,Update_Date) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		insForm.Exec(id,
			Judul_Task,
			Jenis_Task,
			Desc_Task,
			Tgl_Task,
			Tgl_Selesai,
			Total_Jam,
			Status_Task,
			Created_At,
			Create_Date,
			Update_At,
			Update_Date)
		log.Println("INSERT: id : " + id + " | Judul_Task  : " + Judul_Task + " | Jenis_Task  : " + Jenis_Task + " | Desc_Task   : " + Desc_Task + " | Tgl_Task    : " + Tgl_Task + " | Tgl_Selesai : " + Tgl_Selesai + " | Total_Jam   : " + Total_Jam + " | Status_Task : " + Status_Task + " | Created_At  : " + Created_At + " | Create_Date : " + Create_Date + " | Update_At   : " + Update_At + " | Update_Date : " + Update_Date)
	}
	http.Redirect(w, r, "/", 301)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Judul_Task := r.FormValue("Judul_Task")
		Jenis_Task := r.FormValue("Jenis_Task")
		Desc_Task := r.FormValue("Desc_Task")
		Tgl_Task := r.FormValue("Tgl_Task")
		Tgl_Selesai := r.FormValue("Tgl_Selesai")
		Total_Jam := r.FormValue("Total_Jam")
		Status_Task := r.FormValue("Status_Task")
		Created_At := r.FormValue("Created_At")
		Create_Date := r.FormValue("Create_Date")
		Update_At := r.FormValue("Update_At")
		Update_Date := r.FormValue("Update_Date")
		id := r.FormValue("Id")
		insForm, err := db.Prepare("UPDATE list_todo_tabs SET id =?, Judul_Task =?,Jenis_Task =?,Desc_Task =?, Tgl_Task =?, Tgl_Selesai =?, Total_Jam =?, Status_Task =?, Created_At =?, Create_Date =?, Update_At =?, Update_Date =? WHERE id=?")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		insForm.Exec(id,
			Judul_Task,
			Jenis_Task,
			Desc_Task,
			Tgl_Task,
			Tgl_Selesai,
			Total_Jam,
			Status_Task,
			Created_At,
			Create_Date,
			Update_At,
			Update_Date)
		log.Println("UPDATE: id : " + id + " | Judul_Task  : " + Judul_Task + " | Jenis_Task  : " + Jenis_Task + " | Desc_Task   : " + Desc_Task + " | Tgl_Task    : " + Tgl_Task + " | Tgl_Selesai : " + Tgl_Selesai + " | Total_Jam   : " + Total_Jam + " | Status_Task : " + Status_Task + " | Created_At  : " + Created_At + " | Create_Date : " + Create_Date + " | Update_At   : " + Update_At + " | Update_Date : " + Update_Date)
	}
	http.Redirect(w, r, "/", 301)
	defer db.Close()

}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	ltd := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM list_todo_tabs WHERE id=?")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	delForm.Exec(ltd)
	log.Println("DELETE")
	http.Redirect(w, r, "/", 301)
	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
