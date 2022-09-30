package squares

import (
	"reflect"
	"testing"
)

func Test_traverse(t *testing.T) {
	type args struct {
		fields    []Point
		startEdge Edge
	}
	tests := []struct {
		name    string
		args    args
		want    []Edge
		wantErr bool
	}{
		{
			name: "testOK",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					{2, 2},
					{3, 2},
					{1, 3},
					{2, 3},
					{3, 3},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{5, 1},
				{8, 1},
				{8, 2},
				{7, 2},
				{6, 2},
				{6, 3},
				{3, 3},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testErr",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					{2, 2},
					{3, 2},
					{1, 3},
					{2, 3},
					{3, 3},
				},
				startEdge: Edge{4, 2},
			},
			want:    []Edge{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := traverse(tt.args.fields, tt.args.startEdge)
			if (err != nil) != tt.wantErr {
				t.Errorf("traverse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("traverse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
