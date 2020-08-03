package main

import (
	"flag"
	"net/http"

	"github.com/jinzhu/gorm"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *gorm.DB

func main() {
	var conf Config
	var flagConfigPath = flag.String("config", "", "path to config file")
	var err error

	flag.Parse()
	err = initConfig(*flagConfigPath, &conf)
	if err != nil {
		panic("Unable to load config: " + err.Error())
	}

	db, err = initDatabase(&conf)
	if err != nil {
		panic("Unable to connect to the database: " + err.Error())
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/users", handleGetUsers)
	e.DELETE("/users/:user_id", handleDeleteUser)
	e.POST("/users", handleAddUser)

	e.Logger.Fatal(e.Start(":" + conf.Port))
}

// Fetches all users from the database and return them in 'application/json' format.
// It would take one optional 'query parameter' called 'active'(boolean type).
// If provided returns only users with 'active' field set to given value
// true or false, if not - returns all users.
func handleGetUsers(c echo.Context) error {
	filter := newUsersFilter()
	activeParam := c.Request().URL.Query().Get("active")
	if err := filter.SetActiveFilter(activeParam); err != nil {
		return c.JSON(http.StatusBadRequest, responseError{"filtering is not valid"})
	}

	users := findAllUsers(db, filter)

	return c.JSON(http.StatusOK, users)
}

// Deletes user with 'user_id' provided as 'path parameter'
func handleDeleteUser(c echo.Context) error {
	userID := c.Param("user_id")
	user := &User{UserID: userID}
	if !user.isValidID() {
		return c.JSON(http.StatusBadRequest, responseError{"user id is not valid"})
	}
	if err := user.load(db); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, responseError{"no such user found"})
		}
		return c.JSON(http.StatusInternalServerError, responseError{"unable to load the user"})
	}
	if err := user.delete(db); err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{"unable to delete the user"})
	}

	return c.String(http.StatusOK, "")
}

// Creates new user. request body: {name string, active boolean},
// and returns full user data{name, user_id, active}
func handleAddUser(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusBadRequest, responseError{"only application/json content type is allowed"})
	}
	req := new(requestAddUser)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, responseError{"unable to create user"})
	}
	u := req.user()
	if !u.isValid() {
		return c.JSON(http.StatusBadRequest, responseError{"user's data is not valid"})
	}
	if err := u.add(db); err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{"unable to save the data"})
	}

	return c.JSON(http.StatusCreated, u)
}

type requestAddUser struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func (t *requestAddUser) user() *User {
	return &User{
		Name:   t.Name,
		Active: t.Active,
	}
}

type responseError struct {
	Msg string `json:"msg"`
}
