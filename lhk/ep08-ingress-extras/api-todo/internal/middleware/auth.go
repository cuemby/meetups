package middleware

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
)

func CustomJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hostAccounts := os.Getenv("HOST_ACCOUNTS")
		token := c.Request().Header["Authorization"]
		if len(token) == 0 {
			return echo.ErrUnauthorized
		}
		req, _ := http.NewRequest("GET", "http://"+hostAccounts+"/api/v1/accounts/user/", nil)
		req.Header.Set("Authorization", token[0])
		client := &http.Client{}
		resp, err := client.Do(req)
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			responseData, err := ioutil.ReadAll(resp.Body)
			var data map[string]map[string]string
			err = json.Unmarshal(responseData, &data)
			if err == nil {
				c.Set("user", data["user"]["id"])
				return next(c)
			} else {
				print(err.Error())
				return echo.ErrUnauthorized
			}
		} else {
			return echo.ErrUnauthorized
		}
	}
}
