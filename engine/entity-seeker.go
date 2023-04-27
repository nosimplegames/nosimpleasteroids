package engine

type EntitySeeker struct {
	Entity IEntity
}

func (seeker EntitySeeker) IsThereOfType(entityType string) bool {
	hasSameType := seeker.Entity.GetType() == entityType

	if hasSameType {
		return true
	}

	for _, child := range seeker.Entity.GetChildren() {
		childSeeker := EntitySeeker{
			Entity: child,
		}

		hasSameType = childSeeker.IsThereOfType(entityType)
		if hasSameType {
			return true
		}
	}

	return false
}
