//----------------------------------------------------------------------------
// ZGork - a playful imitation of Zork.
// Copyright (C) 2018 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//----------------------------------------------------------------------------

// Package xii contains helpers for 12 Factor.
package xii

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// Int is
type Int struct {
	Required     bool
	DefaultValue int
	Help         string // short help message if required and not found
}

// String is
type String struct {
	Required     bool
	DefaultValue string
	Help         string // short help message if required and not found
}

// AsInt is
func AsInt(key string, opts Int) (int, error) {
	val, found := os.LookupEnv(key)
	if !found {
		if opts.Required {
			return 0, errors.Errorf("%s:\n\t%s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("%s:\n\t%s", key, "value is not an integer"))
	}
	return i, nil
}

// AsString is
func AsString(key string, opts String) (string, error) {
	val, found := os.LookupEnv(key)
	if !found {
		if opts.Required {
			return "", errors.Errorf("%s:\n\t%s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}
	return val, nil
}
