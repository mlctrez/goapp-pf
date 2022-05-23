package pagination

// from file "react-core/src/components/Pagination/PaginationOptionsMenu.tsx"

type OptionsMenuProps struct {
	// ClassName - Custom class name added to the pagination options menu. Optional.
	ClassName string
	// WidgetId - Id added to the title of the Pagination options menu. Optional.
	WidgetId string
	// IsDisabled - Flag indicating if pagination options menu is disabled. Optional.
	IsDisabled bool
	// DropDirection - Menu will open up or open down from the options menu toggle. Optional.
	DropDirection any /* "up" | "down" */
	// PerPageOptions - Array of titles and values which will be the options on the options menu dropdown. Optional.
	PerPageOptions []any // PerPageOptions 
	// ItemsPerPageTitle - The title of the pagination options menu. Optional.
	ItemsPerPageTitle string
	// Page - Current page number. Optional.
	Page int
	// PerPageSuffix - The suffix to be displayed after each option on the options menu dropdown. Optional.
	PerPageSuffix string
	// ItemsTitle - The type or title of the items being paginated. Optional.
	ItemsTitle string
	// OptionsToggle - Accessible label for the options toggle. Optional.
	OptionsToggle string
	// ItemCount - The total number of items being paginated. Optional.
	ItemCount int
	// FirstIndex - The first index of the items being paginated. Optional.
	FirstIndex int
	// LastIndex - The last index of the items being paginated. Optional.
	LastIndex int
	// DefaultToFullPage - Flag to show last full page of results if perPage selected > remaining rows. Optional.
	DefaultToFullPage bool
	// PerPage - The number of items to be displayed per page. Optional.
	PerPage int
	// LastPage - The number of the last page. Optional.
	LastPage int
	// ToggleTemplate - This will be shown in pagination toggle span. You can use firstIndex, lastIndex, itemCount,
	// itemsTitle props.
	ToggleTemplate any /* func(props any // ToggleTemplateProps ) any // React.ReactElement  | string */
	// OnPerPageSelect - Function called when user selects number of items per page. Optional.
	OnPerPageSelect any // OnPerPageSelect 
	// OfWord - Label for the English word "of". Optional.
	OfWord string
	// PerPageComponent - Component to be used for wrapping the toggle contents. Use 'button' when you want all
	// of the toggle text to be clickable. Optional.
	PerPageComponent any /* "div" | "button" */
}

type OptionsMenuState struct {
	// IsOpen - 
	IsOpen bool
}
