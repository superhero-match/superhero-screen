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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	err := os.Setenv("TEST_CONFIG", "config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	// App
	assert.Equal(t, ":3200", cfg.App.Port, "The port should be :4000.")
	assert.Equal(t, "2006-01-02T15:04:05", cfg.App.TimeFormat, "The time format should be 2006-01-02T15:04:05.")

	// Elasticsearch
	assert.Equal(t, "localhost", cfg.ES.Host, "The host should be localhost.")
	assert.Equal(t, "9200", cfg.ES.Port, "The port should be 9200.")
	assert.Equal(t, "superheromatch", cfg.ES.Cluster, "The cluster should be superheromatch.")
	assert.Equal(t, "superhero", cfg.ES.Index, "The index should be superhero.")
}
