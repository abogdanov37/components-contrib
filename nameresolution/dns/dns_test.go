// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation and Dapr Contributors.
// Licensed under the MIT License.
// ------------------------------------------------------------

package dns

import (
	"testing"

	"github.com/dapr/components-contrib/nameresolution"
	"github.com/dapr/dapr/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestResolve(t *testing.T) {
	resolver := NewResolver(logger.NewLogger("test"))
	request := nameresolution.ResolveRequest{ID: "myid", Namespace: "abc", Port: 1234}

	u := "myid-dapr:1234"
	target, err := resolver.ResolveID(request)

	assert.Nil(t, err)
	assert.Equal(t, target, u)
}