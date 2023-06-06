package v2models

type TreeModel interface {
	Model
	getParentID() *uint
}

type TreeModelStore[M TreeModel] struct {
	*ModelStore[M]
	*internalTreeModelStore[M]
}

func (s TreeModelStore[M]) GetChildrenPathToParent(originChild string, destinationParent string) (path []*M, connected bool, err error) {
	originChildModel, err := s.Get(originChild)
	if err != nil {
		return nil, false, err
	}
	destinationParentModel, err := s.Get(destinationParent)
	if err != nil {
		return nil, false, err
	}
	destinationParentID := destinationParentModel.getID()
	return s.getChildrenPathToParent(s.db, &originChildModel, &destinationParentID)
}
