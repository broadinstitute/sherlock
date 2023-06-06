package v2controllers

import "github.com/broadinstitute/sherlock/internal/models/v2models"

type TreeModelController[M v2models.TreeModel, R Readable[M], C Creatable[M], E Editable[M]] struct {
	*ModelController[M, R, C, E]
	treeModelStore *v2models.TreeModelStore[M]
}

func (c TreeModelController[M, R, C, E]) GetChildrenPathToParent(originChild string, destinationParent string) (path []R, connected bool, err error) {
	results, connected, err := c.treeModelStore.GetChildrenPathToParent(originChild, destinationParent)
	readables := make([]R, 0)
	for _, result := range results {
		if result != nil {
			readables = append(readables, *c.modelToReadable(result))
		}
	}
	return readables, connected, err
}
