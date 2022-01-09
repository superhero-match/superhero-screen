/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/superhero-match/superhero-screen/cmd/api/model"
)

// CheckEmail checks if email already registered.
// This check is done after the user has logged-in.
// If email is registered, return status OK and true for registered.
// If not, return status OK and false for registered, and user will be redirected to register service.
func (ctl *Controller) CheckEmail(c *gin.Context) {
	var isRegistered bool
	var isDeleted bool
	var isBlocked bool

	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":       http.StatusInternalServerError,
			"isRegistered": isRegistered,
			"isDeleted":    isDeleted,
			"isBlocked":    isBlocked,
			"superhero":    nil,
		})

		return
	}

	//err := ctl.Service.DeleteIndex()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = ctl.Service.CreateIndex()
	//if err != nil {
	//	fmt.Println(err)
	//}

	resp, err := ctl.Service.CheckEmailExists(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":       http.StatusInternalServerError,
			"isRegistered": isRegistered,
			"isDeleted":    isDeleted,
			"isBlocked":    isBlocked,
			"superhero":    nil,
		})

		return
	}

	fmt.Println()
	fmt.Printf("resp: %+v", resp)
	fmt.Println()

	var superhero model.Superhero

	if resp.IsRegistered && resp.Superhero != nil {
		superhero.ID = resp.Superhero.ID
		superhero.Email = resp.Superhero.Email
		superhero.Name = resp.Superhero.Name
		superhero.SuperheroName = resp.Superhero.SuperheroName
		superhero.MainProfilePicURL = resp.Superhero.MainProfilePicURL
		superhero.ProfilePicsUrls = make([]string, 0)
		superhero.Gender = resp.Superhero.Gender
		superhero.LookingForGender = resp.Superhero.LookingForGender
		superhero.Age = resp.Superhero.Age
		superhero.LookingForAgeMin = resp.Superhero.LookingForAgeMin
		superhero.LookingForAgeMax = resp.Superhero.LookingForAgeMax
		superhero.LookingForDistanceMax = resp.Superhero.LookingForDistanceMax
		superhero.DistanceUnit = resp.Superhero.DistanceUnit
		superhero.Lat = resp.Superhero.Location.Lat
		superhero.Lon = resp.Superhero.Location.Lon
		superhero.Birthday = resp.Superhero.Birthday
		superhero.Country = resp.Superhero.Country
		superhero.City = resp.Superhero.City
		superhero.SuperPower = resp.Superhero.SuperPower
		superhero.AccountType = resp.Superhero.AccountType
		superhero.CreatedAt = resp.Superhero.CreatedAt
	}

	fmt.Println()
	fmt.Printf("Superhero Before Marshalling: %+v", superhero)
	fmt.Println()

	res, err := json.Marshal(superhero)
	if err != nil {
		fmt.Println("json.Marshal")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":       http.StatusInternalServerError,
			"isRegistered": isRegistered,
			"isDeleted":    isDeleted,
			"isBlocked":    isBlocked,
			"superhero":    nil,
		})

		return
	}

	fmt.Println()
	fmt.Printf("Superhero After Marshalling: %s", string(res))
	fmt.Println()

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"isRegistered": resp.IsRegistered,
		"isDeleted":    resp.IsDeleted,
		"isBlocked":    resp.IsBlocked,
		"superhero":    superhero,
	})
}
