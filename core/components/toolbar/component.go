package toolbar

// from file "react-core/src/components/Toolbar/Toolbar.tsx"

type Props struct {
	// ClearAllFilters - Optional callback for clearing all filters in the toolbar. Optional.
	ClearAllFilters func()
	// ClearFiltersButtonText - Text to display in the clear all filters button. Optional.
	ClearFiltersButtonText string
	// CustomChipGroupContent - Custom content appended to the filter generated chip group. To maintain spacing
	// and styling, each node should be wrapped in a ToolbarItem or ToolbarGroup. This property will remove the
	// default "Clear all filters" button. Optional.
	CustomChipGroupContent any // React.ReactNode 
	// CollapseListedFiltersBreakpoint - The breakpoint at which the listed filters in chip groups are collapsed
	// down to a summary. Optional.
	CollapseListedFiltersBreakpoint any /* "all" | "md" | "lg" | "xl" | "2xl" */
	// IsExpanded - Flag indicating if a data toolbar toggle group's expandable content is expanded. Optional.
	IsExpanded bool
	// ToggleIsExpanded - A callback for setting the isExpanded flag. Optional.
	ToggleIsExpanded func()
	// ClassName - Classes applied to root element of the data toolbar. Optional.
	ClassName string
	// Children - Content to be rendered as rows in the data toolbar. Optional.
	Children any // React.ReactNode 
	// Id - Id of the data toolbar. Optional.
	Id string
	// IsFullHeight - Flag indicating the toolbar height should expand to the full height of the container. Optional.
	IsFullHeight bool
	// IsStatic - Flag indicating the toolbar is static. Optional.
	IsStatic bool
	// UsePageInsets - Flag indicating the toolbar should use the Page insets. Optional.
	UsePageInsets bool
	// IsSticky - Flag indicating the toolbar should stick to the top of its container. Optional.
	IsSticky bool
	// Inset - Insets at various breakpoints. Optional.
	Inset map[string]string /* { default:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, sm:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, md:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, lg:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, xl:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, 2xl:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl } } */
	// NumberOfFiltersText - Text to display in the total number of applied filters ToolbarFilter. Optional.
	NumberOfFiltersText func(numberOfFilters int) string
}

type State struct {
	// IsManagedToggleExpanded - Flag used if the user has opted NOT to manage the 'isExpanded' state of the
	// toggle group. Indicates whether or not the toggle group is expanded.
	IsManagedToggleExpanded bool
	// FilterInfo - Object managing information about how many chips are in each chip group.
	FilterInfo any // FilterInfo 
	// WindowWidth - Used to keep track of window width so we can collapse expanded content when window is resizing.
	WindowWidth int
	// OuiaStateId - 
	OuiaStateId string
}

type FilterInfo map[string]int
