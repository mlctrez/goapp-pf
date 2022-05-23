package dragdrop

// from file "react-core/src/components/DragDrop/DragDrop.tsx"

type DraggableItemPosition struct {
	// DroppableId - Parent droppableId.
	DroppableId string
	// Index - Index of item in parent Droppable.
	Index int
}

type Props struct {
	// Children - Potentially Droppable and Draggable children. Optional.
	Children any // React.ReactNode 
	// OnDrag - Callback for drag event. Return true to allow drag, false to disallow. Optional.
	OnDrag func(source any // DraggableItemPosition ) bool
	// OnDragMove - Callback on mouse move while dragging. Optional.
	OnDragMove func(source any // DraggableItemPosition , dest any // DraggableItemPosition )
	// OnDrop - Callback for drop event. Return true to allow drop, false to disallow. Optional.
	OnDrop func(source any // DraggableItemPosition , dest any // DraggableItemPosition ) bool
}
