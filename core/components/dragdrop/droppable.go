package dragdrop

// from file "react-core/src/components/DragDrop/Droppable.tsx"

type DroppableProps struct {
	// Children - Content rendered inside DragDrop. Optional.
	Children any // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
	// Zone - Name of zone that items can be dragged between. Should specify if there is more than one Droppable
	// on the page. Optional.
	Zone string
	// DroppableId - Id to be passed back on drop events. Optional.
	DroppableId string
	// HasNoWrapper - Don't wrap the component in a div. Requires passing a single child. Optional.
	HasNoWrapper bool
}
