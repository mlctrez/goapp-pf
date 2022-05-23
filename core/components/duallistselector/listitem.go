package duallistselector

// from file "react-core/src/components/DualListSelector/DualListSelectorListItem.tsx"

type ListItemProps struct {
	// Children - Content rendered inside the dual list selector. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes applied to the dual list selector. Optional.
	ClassName string
	// IsSelected - Flag indicating the list item is currently selected. Optional.
	IsSelected bool
	// OnOptionSelect - Callback fired when an option is selected. Optional.
	OnOptionSelect func(e any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, id string)
	// Id - ID of the option. Optional.
	Id string
	// OrderIndex - Optional.
	OrderIndex int
	// InnerRef - Optional.
	InnerRef any // React.RefObject 
	// IsDraggable - Flag indicating this item is draggable for reordring. Optional.
	IsDraggable bool
	// DraggableButtonAriaLabel - Accessible label for the draggable button on draggable list items. Optional.
	DraggableButtonAriaLabel string
	// IsDisabled - Flag indicating if the dual list selector is in a disabled state. Optional.
	IsDisabled bool
}
