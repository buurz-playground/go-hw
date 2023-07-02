package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("copy", func(t *testing.T) {
		tests := []struct {
			name   string
			limit  int64
			offset int64
			isErr  bool
			err    error
		}{
			{
				name:   "success, without limit and offset",
				limit:  0,
				offset: 0,
			},
			{
				name:   "success, with limit and offset",
				limit:  2,
				offset: 2,
			},
			{
				name:   "error, offset exceeds file size",
				limit:  0,
				isErr:  true,
				offset: 10,
				err:    ErrOffsetExceedsFileSize,
			},
			{
				name:   "success, limit exceeds file size",
				limit:  10,
				offset: 0,
			},
		}

		for _, test := range tests {
			test := test

			t.Run(test.name, func(t *testing.T) {
				fromFile, err := os.CreateTemp("/tmp", "from")
				require.NoError(t, err)

				defer os.Remove(fromFile.Name())

				toFile, err := os.CreateTemp("/tmp", "to")
				require.NoError(t, err)

				defer os.Remove(toFile.Name())

				_, err = fromFile.Write([]byte("hello"))
				require.NoError(t, err)

				err = Copy(fromFile.Name(), toFile.Name(), test.offset, test.limit)

				if test.isErr {
					require.Error(t, err)
					require.ErrorIs(t, err, test.err)
				} else {
					require.NoError(t, err)
				}
			})
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		err := Copy("not_exist", "out", 0, 0)
		require.Error(t, err)
	})

	t.Run("not regular file", func(t *testing.T) {
		err := Copy("/tmp", "out", 0, 0)
		require.Error(t, err)
	})
}
