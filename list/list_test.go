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
	} {
		list := NewList[int64]()
		for _, value := range testCase.inputValues {
			err := list.PushBack(value)
			assert.NotNil(t, err)
		}

		for _, pos := range testCase.erasePositions {
			_, err := list.Erase(pos)
			assert.NotNil(t, err)
		}

		if testCase.modifyList != nil {
			testCase.modifyList(t, list)
		}

		if testCase.assertFunc != nil {
			testCase.assertFunc(t, list)
		}
	}
}
