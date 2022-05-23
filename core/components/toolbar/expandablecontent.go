package toolbar

// from file "react-core/src/components/Toolbar/ToolbarExpandableContent.tsx"

type ExpandableContentProps struct {
	// ClassName - Classes added to the root element of the data toolbar expandable content. Optional.
	ClassName string
	// IsExpanded - Flag indicating the expandable content is expanded. Optional.
	IsExpanded bool
	// ExpandableContentRef - Expandable content reference for passing to data toolbar children. Optional.
	ExpandableContentRef any // RefObject 
	// ChipContainerRef - Chip container reference for passing to data toolbar children. Optional.
	ChipContainerRef any // RefObject 
	// ClearAllFilters - optional callback for clearing all filters in the toolbar. Optional.
	ClearAllFilters func()
	// ClearFiltersButtonText - Text to display in the clear all filters button. Optional.
	ClearFiltersButtonText string
	// ShowClearFiltersButton - Flag indicating that the clear all filters button should be visible.
	ShowClearFiltersButton bool
}
