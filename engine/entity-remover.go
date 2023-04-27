package engine

type EntityRemover struct {
	Parent IEntity
	Id     string
}

func (remover EntityRemover) Remove() bool {
	for _, child := range remover.Parent.GetChildren() {
		isIt := child.GetId() == remover.Id

		if isIt {
			child.Die()
			return true
		}
	}

	wasRemovedFromChildren := false

	for _, child := range remover.Parent.GetChildren() {
		wasRemovedFromChildren = EntityRemover{
			Parent: child,
			Id:     remover.Id,
		}.Remove()

		if wasRemovedFromChildren {
			break
		}
	}

	return wasRemovedFromChildren
}
