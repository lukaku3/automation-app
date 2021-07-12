package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := "1323"
	e := echo.New()

	// Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/hello", hello)
	// user
	e.POST("/users", createUser)
	e.GET("/users", findUser)
	e.PUT("/users", updateUser)
	e.DELETE("/users", deleteUser)
	// server
	e.GET("/hello", hello)
	e.POST("/server", createServer)
	//e.GET("/server", findServer)
	//e.GET("/server", commons.findServer)
	e.PUT("/server", updateServer)
	e.DELETE("/server", deleteServer)

	e.Logger.Fatal(e.Start(":"+port))
}

func loadConfig() (*config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal("loadConfig os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func findUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//func findServer(c echo.Context) error {
	//var (
	//	id int
	//	servername string
	//	user string
	//	password sql.NullString
	//	//created_at string
	//	//updated_at string
	//	//deleted_at sql.NullString
	//)
	//db, err := sql.Open("mysql",
	//	"db_user:db_password@tcp(127.0.0.1:3306)/automation_db")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//rows, err := db.Query("select id, servername, user, password from server where deleted_at is NULL")
	//defer rows.Close()
	//for rows.Next() {
	//	//err := rows.Scan(&id, &servername, &user, &password, &created_at, &updated_at, &deleted_at)
	//	err := rows.Scan(&id, &servername, &user, &password)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	//log.Println(id, servername)
	//	fmt.Println(id, servername, user, password)
	//}
	//if err = rows.Close(); err != nil {
	//	// but what should we do if there's an error?
	//	log.Println(err)
	//}
	//
	//fmt.Println(json.Marshal(rows))

//
//	return c.String(http.StatusOK, "Hello, World!")
//}

func updateServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func deleteServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// server func
