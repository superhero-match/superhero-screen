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
package main

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-screen/cmd/api/controller"
	"github.com/superhero-match/superhero-screen/cmd/api/service"
	"github.com/superhero-match/superhero-screen/internal/config"
	"github.com/superhero-match/superhero-screen/internal/es"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	esClient, err := elastic.NewClient(
		elastic.SetURL(
			fmt.Sprintf(
				"http://%s:%s",
				cfg.ES.Host,
				cfg.ES.Port,
			),
		),
	)
	if err != nil {
		panic(err)
	}

	e := es.New(esClient, cfg.ES.Index)

	srv := service.New(e)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	ctrl := controller.New(srv, logger, cfg.App.TimeFormat)
	if err != nil {
		panic(err)
	}
	r := ctrl.RegisterRoutes()

	err = r.Run(cfg.App.Port)
	if err != nil {
		panic(err)
	}
}
