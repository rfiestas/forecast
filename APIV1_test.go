package p

import "testing"

func TestDegToCardinal(t *testing.T) {
	type args struct {
		deg float64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "north",
			args:    args{0},
			want:    "N",
			wantErr: false,
		},
		{
			name:    "south",
			args:    args{180},
			want:    "S",
			wantErr: false,
		},
		{
			name:    "south",
			args:    args{361},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DegToCardinal(tt.args.deg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DegToCardinal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DegToCardinal() = %v, want %v", got, tt.want)
			}
		})
	}
}
