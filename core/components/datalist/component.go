package datalist

// from file "react-core/src/components/DataList/DataList.tsx"

type Props struct {
	// Children - Content rendered inside the DataList list. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the DataList list. Optional.
	ClassName string
	// AriaLabel - Adds accessible text to the DataList list.
	AriaLabel string
	// OnSelectDataListItem - Optional callback to make DataList selectable, fired when DataListItem selected.
	// Optional.
	OnSelectDataListItem func(id string)
	// OnDragFinish - Optional.
	OnDragFinish func(newItemOrder []string)
	// OnDragStart - Optional.
	OnDragStart func(id string)
	// OnDragMove - Optional.
	OnDragMove func(oldIndex int, newIndex int)
	// OnDragCancel - Optional.
	OnDragCancel func()
	// SelectedDataListItemId - Id of DataList item currently selected. Optional.
	SelectedDataListItemId string
	// IsCompact - Flag indicating if DataList should have compact styling. Optional.
	IsCompact bool
	// GridBreakpoint - Specifies the grid breakpoints. Optional.
	GridBreakpoint any /* "none" | "always" | "sm" | "md" | "lg" | "xl" | "2xl" */
	// WrapModifier - Determines which wrapping modifier to apply to the DataList. Optional.
	WrapModifier any /* any // DataListWrapModifier  | "nowrap" | "truncate" | "breakWord" */
	// ItemOrder - Optional.
	ItemOrder []string
}

type State struct {
	// DraggedItemId - 
	DraggedItemId string
	// DraggingToItemIndex - 
	DraggingToItemIndex int
	// Dragging - 
	Dragging bool
	// TempItemOrder - 
	TempItemOrder []string
}

type ContextProps struct {
	// IsSelectable - 
	IsSelectable bool
	// SelectedDataListItemId - 
	SelectedDataListItemId string
	// UpdateSelectedDataListItem - 
	UpdateSelectedDataListItem func(id string)
	// IsDraggable - 
	IsDraggable bool
	// DragStart - 
	DragStart func(e any // React.DragEvent )
	// DragEnd - 
	DragEnd func(e any // React.DragEvent )
	// Drop - 
	Drop func(e any // React.DragEvent )
	// DragKeyHandler - 
	DragKeyHandler func(e any // React.KeyboardEvent )
}
