package toolbar

// from file "react-core/src/components/Toolbar/ToolbarChipGroupContent.tsx"

type ChipGroupContentProps struct {
	// ClassName - Classes applied to root element of the data toolbar content row. Optional.
	ClassName string
	// IsExpanded - Flag indicating if a data toolbar toggle group's expandable content is expanded. Optional.
	IsExpanded bool
	// ChipGroupContentRef - Chip group content reference for passing to data toolbar children. Optional.
	ChipGroupContentRef any // RefObject 
	// ClearAllFilters - optional callback for clearing all filters in the toolbar. Optional.
	ClearAllFilters func()
	// ShowClearFiltersButton - Flag indicating that the clear all filters button should be visible.
	ShowClearFiltersButton bool
	// ClearFiltersButtonText - Text to display in the clear all filters button. Optional.
	ClearFiltersButtonText string
	// NumberOfFilters - Total number of filters currently being applied across all ToolbarFilter components.
	NumberOfFilters int
	// NumberOfFiltersText - Text to display in the total number of applied filters ToolbarFilter. Optional.
	NumberOfFiltersText func(numberOfFilters int) string
	// CollapseListedFiltersBreakpoint - The breakpoint at which the listed filters in chip groups are collapsed
	// down to a summary. Optional.
	CollapseListedFiltersBreakpoint any /* "all" | "md" | "lg" | "xl" | "2xl" */
	// CustomChipGroupContent - Custom additional content appended to the generated chips. To maintain spacing
	// and styling, each node should be a ToolbarItem or ToolbarGroup. This property will remove the built in
	// "Clear all filters" button. Optional.
	CustomChipGroupContent any // React.ReactNode 
}
