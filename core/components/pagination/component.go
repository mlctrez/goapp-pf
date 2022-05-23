package pagination

// from file "react-core/src/components/Pagination/Pagination.tsx"

type PerPageOptions struct {
	// Title - option title. Optional.
	Title string
	// Value - option value. Optional.
	Value int
}

type Titles struct {
	// Page - The title of a page displayed beside the page number. Optional.
	Page string
	// Pages - The title of a page displayed beside the page number (plural form). Optional.
	Pages string
	// Items - The type or title of the items being paginated. Optional.
	Items string
	// ItemsPerPage - The title of the pagination options menu. Optional.
	ItemsPerPage string
	// PerPageSuffix - The suffix to be displayed after each option on the options menu dropdown. Optional.
	PerPageSuffix string
	// ToFirstPage - Accessible label for the button which moves to the first page. Optional.
	ToFirstPage string
	// ToPreviousPage - Accessible label for the button which moves to the previous page. Optional.
	ToPreviousPage string
	// ToLastPage - Accessible label for the button which moves to the last page. Optional.
	ToLastPage string
	// ToNextPage - Accessible label for the button which moves to the next page. Optional.
	ToNextPage string
	// OptionsToggle - Accessible label for the options toggle. Optional.
	OptionsToggle string
	// CurrPage - Accessible label for the input displaying the current page. Optional.
	CurrPage string
	// PaginationTitle - Accessible label for the pagination component. Optional.
	PaginationTitle string
	// OfWord - Accessible label for the English word "of". Optional.
	OfWord string
}

type Props struct {
	// Children - What should be rendered inside. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes for the container. Optional.
	ClassName string
	// ItemCount - Total number of items. Optional.
	ItemCount int
	// Variant - Position where pagination is rendered. Optional.
	Variant any /* "top" | "bottom" | any // PaginationVariant  */
	// IsDisabled - Flag indicating if pagination is disabled. Optional.
	IsDisabled bool
	// IsCompact - Flag indicating if pagination is compact. Optional.
	IsCompact bool
	// IsStatic - Flag indicating if pagination should not be sticky on mobile. Optional.
	IsStatic bool
	// IsSticky - Flag indicating if pagination should stick to its position (based on variant). Optional.
	IsSticky bool
	// PerPage - Number of items per page. Optional.
	PerPage int
	// PerPageOptions - Array of the number of items per page  options. Optional.
	PerPageOptions []any // PerPageOptions 
	// DefaultToFullPage - Indicate whether to show last full page of results when user selects perPage value
	// greater than remaining rows. Optional.
	DefaultToFullPage bool
	// FirstPage - Page we start at. Optional.
	FirstPage int
	// Page - Current page number. Optional.
	Page int
	// Offset - Start index of rows to display, used in place of providing page. Optional.
	Offset int
	// ItemsStart - First index of items on current page. Optional.
	ItemsStart int
	// ItemsEnd - Last index of items on current page. Optional.
	ItemsEnd int
	// WidgetId - ID to ideintify widget on page. Optional.
	WidgetId string
	// DropDirection - Direction of dropdown context menu. Optional.
	DropDirection any /* "up" | "down" */
	// Titles - Object with titles to display in pagination. Optional.
	Titles any // PaginationTitles 
	// ToggleTemplate - This will be shown in pagination toggle span. You can use firstIndex, lastIndex, itemCount,
	// itemsTitle, ofWord props. Optional.
	ToggleTemplate any /* func(props any // ToggleTemplateProps ) any // React.ReactElement  | string */
	// OnSetPage - Function called when user sets page. Optional.
	OnSetPage any // OnSetPage 
	// OnFirstClick - Function called when user clicks on navigate to first page. Optional.
	OnFirstClick func(event any // React.SyntheticEvent , page int)
	// OnPreviousClick - Function called when user clicks on navigate to previous page. Optional.
	OnPreviousClick func(event any // React.SyntheticEvent , page int)
	// OnNextClick - Function called when user clicks on navigate to next page. Optional.
	OnNextClick func(event any // React.SyntheticEvent , page int)
	// OnLastClick - Function called when user clicks on navigate to last page. Optional.
	OnLastClick func(event any // React.SyntheticEvent , page int)
	// OnPageInput - Function called when user inputs page number. Optional.
	OnPageInput func(event any // React.SyntheticEvent , page int)
	// OnPerPageSelect - Function called when user selects number of items per page. Optional.
	OnPerPageSelect any // OnPerPageSelect 
	// PerPageComponent - Component to be used for wrapping the toggle contents. Use 'button' when you want all
	// of the toggle text to be clickable. Optional.
	PerPageComponent any /* "div" | "button" */
}
