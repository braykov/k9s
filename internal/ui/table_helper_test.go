// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	uu := map[string]struct {
		s, e string
	}{
		"empty": {},
		"max": {
			s: "/app.kubernetes.io/instance=prom,app.kubernetes.io/name=prometheus,app.kubernetes.io/component=server",
			e: "/app.kubernetes.io/instance=prom,app.kubernetes.io...",
		},
		"less": {
			s: "app=fred,env=blee",
			e: "app=fred,env=blee",
		},
	}

	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, truncate(u.s, 50))
		})
	}
}

func TestIsLabelSelector(t *testing.T) {
	uu := map[string]struct {
		s  string
		ok bool
	}{
		"empty":       {s: ""},
		"cool":        {s: "-l app=fred,env=blee", ok: true},
		"no-flag":     {s: "app=fred,env=blee", ok: true},
		"no-space":    {s: "-lapp=fred,env=blee", ok: true},
		"wrong-flag":  {s: "-f app=fred,env=blee"},
		"missing-key": {s: "=fred"},
		"missing-val": {s: "fred="},
	}

	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.ok, IsLabelSelector(u.s))
		})
	}
}

func TestTrimLabelSelector(t *testing.T) {
	uu := map[string]struct {
		sel, e string
	}{
		"cool":    {"-l app=fred,env=blee", "app=fred,env=blee"},
		"noSpace": {"-lapp=fred,env=blee", "app=fred,env=blee"},
	}

	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, TrimLabelSelector(u.sel))
		})
	}
}
