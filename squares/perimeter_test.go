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
			name: "testOKSquare",
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
			name: "testOKSingle1",
			args: args{
				fields: []Point{
					{1, 1},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKSingle2",
			args: args{
				fields: []Point{
					{1, 1},
				},
				startEdge: Edge{0, 1},
			},
			want: []Edge{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 0},
			},
			wantErr: false,
		},
		{
			name: "testOKStraightHorizontal",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{4, 1},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
				{3, 1},
				{3, 2},
				{2, 2},
				{1, 2},
				{0, 2},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKStraightVertical",
			args: args{
				fields: []Point{
					{1, 1},
					{1, 2},
					{1, 3},
					{1, 4},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{3, 1},
				{3, 2},
				{3, 3},
				{2, 3},
				{1, 3},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKT",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{2, 2},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
				{3, 1},
				{3, 2},
				{3, 3},
				{0, 2},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKL",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{1, 2},
					{1, 3},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{1, 0},
				{1, 1},
				{1, 2},
				{2, 1},
				{3, 1},
				{3, 2},
				{3, 3},
				{2, 3},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKHoleOutside",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					//{2, 2},
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
				{4, 1},
				{7, 1},
				{7, 2},
				{6, 2},
				{5, 2},
				{5, 3},
				{3, 3},
				{0, 3},
			},
			wantErr: false,
		},
		{
			//hole inside traverses clockwise
			name: "testOKHoleInside",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					//{2, 2},
					{3, 2},
					{1, 3},
					{2, 3},
					{3, 3},
				},
				startEdge: Edge{1, 2},
			},
			want: []Edge{
				{1, 2},
				{3, 1},
				{6, 0},
				{4, 3},
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

func Test_traverseMatrix(t *testing.T) {
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
			name: "testOKSquare",
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
			name: "testOKSingle1",
			args: args{
				fields: []Point{
					{1, 1},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKSingle2",
			args: args{
				fields: []Point{
					{1, 1},
				},
				startEdge: Edge{0, 1},
			},
			want: []Edge{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 0},
			},
			wantErr: false,
		},
		{
			name: "testOKStraightHorizontal",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{4, 1},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
				{3, 1},
				{3, 2},
				{2, 2},
				{1, 2},
				{0, 2},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKStraightVertical",
			args: args{
				fields: []Point{
					{1, 1},
					{1, 2},
					{1, 3},
					{1, 4},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{3, 1},
				{3, 2},
				{3, 3},
				{2, 3},
				{1, 3},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKT",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{2, 2},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
				{3, 1},
				{3, 2},
				{3, 3},
				{0, 2},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKL",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{1, 2},
					{1, 3},
				},
				startEdge: Edge{0, 0},
			},
			want: []Edge{
				{0, 0},
				{1, 0},
				{1, 1},
				{1, 2},
				{2, 1},
				{3, 1},
				{3, 2},
				{3, 3},
				{2, 3},
				{0, 3},
			},
			wantErr: false,
		},
		{
			name: "testOKHoleOutside",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					//{2, 2},
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
				{4, 1},
				{7, 1},
				{7, 2},
				{6, 2},
				{5, 2},
				{5, 3},
				{3, 3},
				{0, 3},
			},
			wantErr: false,
		},
		{
			//hole inside traverses clockwise
			name: "testOKHoleInside",
			args: args{
				fields: []Point{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					//{2, 2},
					{3, 2},
					{1, 3},
					{2, 3},
					{3, 3},
				},
				startEdge: Edge{1, 2},
			},
			want: []Edge{
				{1, 2},
				{3, 1},
				{6, 0},
				{4, 3},
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
			got, err := traverseMatrix(tt.args.fields, tt.args.startEdge)
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
