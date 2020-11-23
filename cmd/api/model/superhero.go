/*
  Copyright (C) 2019 - 2021 MWSOFT
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
package model

type Superhero struct {
	ID                    string   `json:"id" struct:"id"`
	Email                 string   `json:"email" struct:"email"`
	Name                  string   `json:"name" struct:"name"`
	SuperheroName         string   `json:"superHeroName" struct:"superHeroName"`
	MainProfilePicURL     string   `json:"mainProfilePicUrl" struct:"mainProfilePicUrl"`
	ProfilePicsUrls       []string `json:"profilePicsUrls" struct:"profilePicsUrls"`
	Gender                int      `json:"gender" struct:"gender"`
	LookingForGender      int      `json:"lookingForGender" struct:"lookingForGender"`
	Age                   int      `json:"age" struct:"age"`
	LookingForAgeMin      int      `json:"lookingForAgeMin" struct:"lookingForAgeMin"`
	LookingForAgeMax      int      `json:"lookingForAgeMax" struct:"lookingForAgeMax"`
	LookingForDistanceMax int      `json:"lookingForDistanceMax" struct:"lookingForDistanceMax"`
	DistanceUnit          string   `json:"distanceUnit" struct:"distanceUnit"`
	Lat                   float64  `json:"lat" struct:"lat"`
	Lon                   float64  `json:"lon" struct:"lon"`
	Birthday              string   `json:"birthday" struct:"birthday"`
	Country               string   `json:"country" struct:"country"`
	City                  string   `json:"city" struct:"city"`
	SuperPower            string   `json:"superPower" struct:"superPower"`
	AccountType           string   `json:"accountType" struct:"accountType"`
	CreatedAt             string   `json:"createdAt"`
}
