package main

import (
	models "backend/Models"
	Repository "backend/Repository"
	"backend/storage"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)




func main(){
	err := godotenv.Load(".env")
	if err!=nil{
        log.Fatal(err)
	}

	config:= &storage.Config{
        Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	db,err := storage.NewConnection(config)



	if(err!=nil){
		log.Fatal(err)
	}


	err = models.MigrateStudent(db)

	if err!=nil{
		log.Fatal(err)
	}

	err = models.MigrateTeacher(db)
	if err!=nil{
		log.Fatal(err)
	}

	err = models.MigrateAttendence(db)
	if err!=nil{
		log.Fatal(err)
	}
    
	r:=Repository.Repository{
		DB:db,
	}

	app := mux.NewRouter();

	app.HandleFunc("/api/AddStudent",r.AddStudent).Methods("POST")
	app.HandleFunc("/api/AddTeacher",r.AddTeacher).Methods("POST")
	app.HandleFunc("/api/GetTeacherAttendence/{id}/{month}/{year}",r.GetTeacherAttendence).Methods("GET")
	app.HandleFunc("/api/PunchIn",r.PunchIn).Methods("POST")
	app.HandleFunc("/api/PunchOut",r.PunchOut).Methods("POST")
	app.HandleFunc("/api/GetClassAttendence/{id}/{day}/{month}/{year}",r.GetClassAttendence).Methods("GET")
	app.HandleFunc("/api/GetStudentAttendence/{id}/{month}/{year}",r.GetStudentAttendence).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000",app));
}