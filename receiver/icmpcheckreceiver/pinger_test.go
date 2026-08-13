// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package icmpcheckreceiver

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDefaultPingerFactoryPrivilegedOnWindows(t *testing.T) {
	p, err := defaultPingerFactory(PingTarget{
		Host:         "127.0.0.1",
		PingCount:    1,
		PingTimeout:  time.Second,
		PingInterval: time.Second,
	})
	require.NoError(t, err)

	defaultP, ok := p.(*defaultPinger)
	require.True(t, ok)

	require.Equal(t, runtime.GOOS == "windows", defaultP.Privileged())
}
