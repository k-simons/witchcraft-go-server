// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"context"
	"net/http"
	"testing"

	werror "github.com/palantir/witchcraft-go-error"
	"github.com/stretchr/testify/assert"
)

func TestStatusCodeMapper(t *testing.T) {
	for _, tc := range []struct {
		name         string
		err          error
		expectedCode int
	}{
		{
			name:         "not found rest error",
			err:          NewError(werror.Error("Test error"), StatusCode(http.StatusNotFound)),
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "werror wrapping not found error",
			err:          werror.WrapWithContextParams(context.Background(), NewError(werror.Error("inner"), StatusCode(http.StatusNotFound)), "outer"),
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "no rest error",
			err:          werror.ErrorWithContextParams(context.Background(), "werror"),
			expectedCode: http.StatusInternalServerError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, StatusCodeMapper(tc.err), tc.expectedCode)
		})
	}
}
