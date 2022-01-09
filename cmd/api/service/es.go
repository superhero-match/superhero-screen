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
package service

import (
	"github.com/superhero-match/superhero-screen/internal/es/model"
)

// CheckEmailExists checks if document of type superhero in superheros index exists
// and if so, if the superhero is blocked.
func (srv *service) CheckEmailExists(email string) (rsp *model.CheckEmailResponse, err error) {
	return srv.ES.CheckEmailExists(email)
}

// CreateIndex crates Elasticsearch index if index does not exist yet.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!! DO NOT CALL THIS METHOD THIS IS ONLY FOR LOCAL DEVELOPMENT PURPOSES !!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (srv *service) CreateIndex() error {
	return srv.ES.CreateIndex()
}

// DeleteIndex deletes the whole index.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!! DO NOT CALL THIS METHOD THIS IS ONLY FOR LOCAL DEVELOPMENT PURPOSES !!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (srv *service) DeleteIndex() error {
	return srv.ES.DeleteIndex()
}

// Ping the Elasticsearch server to make sure that ES is running.
func (srv *service) Ping() error {
	return srv.ES.Ping()
}
