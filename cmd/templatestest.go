package main

import (
	"testing"
	"time"

	"snippetbox.arsenzhamshitov.net/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{
			name: "UTC",
			t:    time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15",
		}, {
			name: "Empty",
			t:    time.Time{},
			want: "",
		}, {
			name: "CET",
			t:    time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 60*60)),
			want: "17 Mar 2022 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.t)
			assert.Equal(t, hd, tt.want)
		})
	}
}
