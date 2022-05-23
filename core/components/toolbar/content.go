package toolbar

// from file "react-core/src/components/Toolbar/ToolbarContent.tsx"

type ContentProps struct {
	// ClassName - Classes applied to root element of the data toolbar content row. Optional.
	ClassName string
	// Visibility - Visibility at various breakpoints. Optional.
	Visibility map[string]string /* { default:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// Visiblity - Optional.
	Visiblity map[string]string /* { default:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// Alignment - Alignment at various breakpoints. Optional.
	Alignment map[string]string /* { default:{ alignRight, alignLeft }, md:{ alignRight, alignLeft }, lg:{ alignRight, alignLeft }, xl:{ alignRight, alignLeft }, 2xl:{ alignRight, alignLeft } } */
	// Children - Content to be rendered as children of the content row. Optional.
	Children any // React.ReactNode 
	// IsExpanded - Flag indicating if a data toolbar toggle group's expandable content is expanded. Optional.
	IsExpanded bool
	// ClearAllFilters - Optional callback for clearing all filters in the toolbar. Optional.
	ClearAllFilters func()
	// ShowClearFiltersButton - Flag indicating that the clear all filters button should be visible. Optional.
	ShowClearFiltersButton bool
	// ClearFiltersButtonText - Text to display in the clear all filters button. Optional.
	ClearFiltersButtonText string
	// ToolbarId - Id of the parent Toolbar component. Optional.
	ToolbarId string
}
