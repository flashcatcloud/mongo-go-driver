// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package mongo

import (
	"testing"

	"go.mongodb.org/mongo-driver/internal/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestChangeStream(t *testing.T) {
	t.Run("nil cursor", func(t *testing.T) {
		cs := &ChangeStream{}

		id := cs.ID()
		assert.Equal(t, int64(0), id, "expected ID 0, got %v", id)
		assert.False(t, cs.Next(bgCtx), "expected Next to return false, got true")
		err := cs.Decode(nil)
		assert.Equal(t, ErrNilCursor, err, "expected error %v, got %v", ErrNilCursor, err)
		err = cs.Err()
		assert.Nil(t, err, "change stream error: %v", err)
		err = cs.Close(bgCtx)
		assert.Nil(t, err, "Close error: %v", err)
	})
}

func TestMergeChangeStreamOptions(t *testing.T) {
	t.Parallel()

	fullDocumentP := func(x options.FullDocument) *options.FullDocument { return &x }
	int32P := func(x int32) *int32 { return &x }

	testCases := []struct {
		description string
		input       []*options.ChangeStreamOptions
		want        *options.ChangeStreamOptions
	}{
		{
			description: "nil",
			input:       nil,
			want:        &options.ChangeStreamOptions{},
		},
		{
			description: "empty",
			input:       []*options.ChangeStreamOptions{},
			want:        &options.ChangeStreamOptions{},
		},
		{
			description: "many ChangeStreamOptions with one configuration each",
			input: []*options.ChangeStreamOptions{
				options.ChangeStream().SetFullDocumentBeforeChange(options.Required),
				options.ChangeStream().SetFullDocument(options.Required),
				options.ChangeStream().SetBatchSize(10),
			},
			want: &options.ChangeStreamOptions{
				FullDocument:             fullDocumentP(options.Required),
				FullDocumentBeforeChange: fullDocumentP(options.Required),
				BatchSize:                int32P(10),
			},
		},
		{
			description: "single ChangeStreamOptions with many configurations",
			input: []*options.ChangeStreamOptions{
				options.ChangeStream().
					SetFullDocumentBeforeChange(options.Required).
					SetFullDocument(options.Required).
					SetBatchSize(10),
			},
			want: &options.ChangeStreamOptions{
				FullDocument:             fullDocumentP(options.Required),
				FullDocumentBeforeChange: fullDocumentP(options.Required),
				BatchSize:                int32P(10),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable.

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			got := mergeChangeStreamOptions(tc.input...)
			assert.Equal(t, tc.want, got, "expected and actual ChangeStreamOptions are different")
		})
	}
}
