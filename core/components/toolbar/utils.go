package toolbar

// from file "react-core/src/components/Toolbar/ToolbarUtils.tsx"

type ContextProps struct {
	// IsExpanded - 
	IsExpanded bool
	// ToggleIsExpanded - 
	ToggleIsExpanded func()
	// ChipGroupContentRef - 
	ChipGroupContentRef any // RefObject 
	// UpdateNumberFilters - 
	UpdateNumberFilters func(categoryName string, numberOfFilters int)
	// NumberOfFilters - 
	NumberOfFilters int
	// ClearAllFilters - Optional.
	ClearAllFilters func()
	// ClearFiltersButtonText - Optional.
	ClearFiltersButtonText string
	// ShowClearFiltersButton - Optional.
	ShowClearFiltersButton bool
	// ToolbarId - Optional.
	ToolbarId string
	// CustomChipGroupContent - Optional.
	CustomChipGroupContent any // React.ReactNode 
}

type ContentContextProps struct {
	// ExpandableContentRef - 
	ExpandableContentRef any // RefObject 
	// ExpandableContentId - 
	ExpandableContentId string
	// ChipContainerRef - 
	ChipContainerRef any // RefObject 
}
