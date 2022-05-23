package dragdrop

// from file "react-core/src/components/DragDrop/Draggable.tsx"

type DraggableProps struct {
	// Children - Content rendered inside DragDrop. Optional.
	Children any // React.ReactNode 
	// HasNoWrapper - Don't wrap the component in a div. Requires passing a single child. Optional.
	HasNoWrapper bool
	// ClassName - Class to add to outer div. Optional.
	ClassName string
}

type DroppableItem struct {
	// Node - 
	Node any // HTMLElement 
	// Rect - 
	Rect any // DOMRect 
	// IsDraggingHost - 
	IsDraggingHost bool
	// DraggableNodes - 
	DraggableNodes []any // HTMLElement 
	// DraggableNodesRects - 
	DraggableNodesRects []any // DOMRect 
}
