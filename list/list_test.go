package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PushBack(t *testing.T) {
	for _, testCase := range []struct {
		inputValues    []int64
		erasePositions []int64
		modifyList     func(t *testing.T, list *List[int64])
		assertFunc     func(t *testing.T, list *List[int64])
	}{
		{
			inputValues:    []int64{1, 2, 3, 4, 5},
			erasePositions: []int64{},
			modifyList:     nil,
			assertFunc: func(t *testing.T, list *List[int64]) {
				assert.Equal(t, list.Size(), 5)
			},
		},
		{
			inputValues:    []int64{1, 2, 3},
			erasePositions: []int64{0, 1, 2},
			modifyList:     nil,
			assertFunc: func(t *testing.T, list *List[int64]) {
				assert.Equal(t, list.Size(), 0)
			},
		},
		{
			inputValues:    []int64{2, 4, 6, 8, 10, 12},
			erasePositions: []int64{},
			modifyList: func(t *testing.T, list *List[int64]) {
				val, err := list.Erase(0)
				assert.Equal(t, val, 2)
				assert.Nil(t, err)

				val, err = list.Erase(1)
				assert.Equal(t, val, 6)
				assert.Nil(t, err)

				val, err = list.Erase(10)
				assert.NotNil(t, err)
				assert.ErrorIs(t, err, ErrEmptyList)
			},
			assertFunc: func(t *testing.T, list *List[int64]) {
				assert.Equal(t, list.Size(), 4)
			},
		},
	} {
		list := NewList[int64]()
		for _, value := range testCase.inputValues {
			err := list.PushBack(value)
			assert.Nil(t, err)
		}

		for _, pos := range testCase.erasePositions {
			_, err := list.Erase(pos)
			assert.Nil(t, err)
		}

		if testCase.modifyList != nil {
			testCase.modifyList(t, list)
		}

		if testCase.assertFunc != nil {
			testCase.assertFunc(t, list)
		}
	}
}
