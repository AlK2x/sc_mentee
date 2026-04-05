package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestMapReduce(t *testing.T) {
	cases := []struct {
		msg        string
		input      []Event
		numWorkers int
		output     []Event
	}{
		{
			msg:        "empty input",
			input:      []Event{},
			numWorkers: 1,
			output:     []Event{},
		},
		{
			msg:        "sort input events",
			input:      []Event{{Timestamp: 1}, {Timestamp: 10}, {Timestamp: 3}},
			numWorkers: 1,
			output:     []Event{{Timestamp: 1}, {Timestamp: 3}, {Timestamp: 10}},
		},
		{
			msg:        "numWorkers more than input events",
			input:      []Event{{Timestamp: 1}, {Timestamp: 10}, {Timestamp: 3}},
			numWorkers: 4,
			output:     []Event{{Timestamp: 1}, {Timestamp: 3}, {Timestamp: 10}},
		},
		{
			msg:        "numWorker not devided to input size",
			input:      []Event{{Timestamp: 1}, {Timestamp: 2}, {Timestamp: 3}},
			numWorkers: 2,
			output:     []Event{{Timestamp: 1}, {Timestamp: 2}, {Timestamp: 3}},
		},
		{
			msg:        "more events",
			input:      []Event{{Timestamp: 12}, {Timestamp: 85}, {Timestamp: 43}, {Timestamp: 7}, {Timestamp: 92}, {Timestamp: 56}, {Timestamp: 21}, {Timestamp: 38}, {Timestamp: 64}, {Timestamp: 10}, {Timestamp: 77}, {Timestamp: 49}, {Timestamp: 3}, {Timestamp: 88}, {Timestamp: 31}, {Timestamp: 60}, {Timestamp: 25}, {Timestamp: 99}, {Timestamp: 14}, {Timestamp: 52}, {Timestamp: 42}},
			numWorkers: 5,
			output:     []Event{{Timestamp: 3}, {Timestamp: 7}, {Timestamp: 10}, {Timestamp: 12}, {Timestamp: 14}, {Timestamp: 21}, {Timestamp: 25}, {Timestamp: 31}, {Timestamp: 38}, {Timestamp: 42}, {Timestamp: 43}, {Timestamp: 49}, {Timestamp: 52}, {Timestamp: 56}, {Timestamp: 60}, {Timestamp: 64}, {Timestamp: 77}, {Timestamp: 85}, {Timestamp: 88}, {Timestamp: 92}, {Timestamp: 99}},
		},
	}

	for i, c := range cases {
		result := make([]Event, 0, len(c.input))
		for v := range MapReduce(c.input, c.numWorkers) {
			result = append(result, v)
		}

		slices.SortFunc(c.input, func(lt, rt Event) int {
			if lt.Timestamp < rt.Timestamp {
				return -1
			} else {
				return 1
			}
		})
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("[%d] %s: failed got %v, expected %v", i, c.msg, result, c.output)
		}
	}
}
