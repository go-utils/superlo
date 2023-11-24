package superlo_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-utils/superlo"
)

func TestMap(t *testing.T) {
	type args[T any, R any] struct {
		collection []T
		iteratee   func(item T, index int) (R, error)
	}
	type testCase[T any, R any] struct {
		name                 string
		args                 args[T, R]
		wantMappedCollection []R
		wantErr              bool
	}
	tests := []testCase[int, int]{
		{
			name: "OK",
			args: args[int, int]{
				collection: []int{1, 2, 3, 4, 5},
				iteratee: func(item int, index int) (int, error) {
					return item * index, nil
				},
			},
			wantMappedCollection: []int{0, 2, 6, 12, 20},
		},
		{
			name: "NG",
			args: args[int, int]{
				collection: []int{0, 2, 3, 4, 6},
				iteratee: func(item int, index int) (int, error) {
					if item%2 != 0 {
						return 0, errors.New("error")
					}
					return item * index, nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMappedCollection, err := superlo.Map(tt.args.collection, tt.args.iteratee)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMappedCollection, tt.wantMappedCollection) {
				t.Errorf("Map() gotMappedCollection = %v, want %v", gotMappedCollection, tt.wantMappedCollection)
			}
		})
	}
}

func TestParallelMap(t *testing.T) {
	type args[T any, R any] struct {
		collection []T
		iteratee   func(item T, index int) (R, error)
	}
	type testCase[T any, R any] struct {
		name                 string
		args                 args[T, R]
		wantMappedCollection []R
		wantErr              bool
	}
	tests := []testCase[int, int]{
		{
			name: "OK",
			args: args[int, int]{
				collection: []int{1, 2, 3, 4, 5},
				iteratee: func(item int, index int) (int, error) {
					return item * index, nil
				},
			},
			wantMappedCollection: []int{0, 2, 6, 12, 20},
		},
		{
			name: "NG",
			args: args[int, int]{
				collection: []int{0, 2, 3, 4, 6},
				iteratee: func(item int, index int) (int, error) {
					if item%2 != 0 {
						return 0, errors.New("error")
					}
					return item * index, nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMappedCollection, err := superlo.ParallelMap(tt.args.collection, tt.args.iteratee)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParallelMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMappedCollection, tt.wantMappedCollection) {
				t.Errorf("ParallelMap() gotMappedCollection = %v, want %v", gotMappedCollection, tt.wantMappedCollection)
			}
		})
	}
}
