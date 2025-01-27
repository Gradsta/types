// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import "testing"

func TestPipeline_Worker_Empty(t *testing.T) {
	// setup tests
	tests := []struct {
		worker *Worker
		want   bool
	}{
		{
			worker: &Worker{Flavor: "foo"},
			want:   false,
		},
		{
			worker: new(Worker),
			want:   true,
		},
	}

	// run tests
	for _, test := range tests {
		got := test.worker.Empty()

		if got != test.want {
			t.Errorf("Empty is %v, want %t", got, test.want)
		}
	}
}
