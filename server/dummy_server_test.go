// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartDummyServer(t *testing.T) {
	r := require.New(t)
	s, err := NewDummyServer(32223, nil)
	r.NoError(err)
	r.True(s != nil)
	ctx := context.Background()
	r.NoError(s.Start(ctx))
}
