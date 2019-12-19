package activitypub

import (
	"reflect"
	"testing"
	"time"
)

func Test_marshalXSD(t *testing.T) {
	tests := []struct {
		name    string
		d       time.Duration
		want    []byte
		wantErr bool
	}{
		{
			name:    "Zero duration",
			d:       0,
			want:    []byte("PT0S"),
			wantErr: false,
		},
		{
			name:    "One year",
			d:       year,
			want:    []byte("P1Y"),
			wantErr: false,
		},
		{
			name:    "XSD:duration example 1st",
			d:       2*year+6*month+5*day+12*time.Hour+35*time.Minute+30*time.Second,
			want:    []byte("P2Y6M5DT12H35M30S"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := marshalXSD(tt.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("marshalXSD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marshalXSD() got = %s, want %s", got, tt.want)
			}
		})
	}
}

