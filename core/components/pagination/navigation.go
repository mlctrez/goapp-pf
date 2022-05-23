package pagination

// from file "react-core/src/components/Pagination/Navigation.tsx"

type NavigationProps struct {
	// ClassName - Additional classes for the container. Optional.
	ClassName string
	// IsDisabled - Flag indicating if the pagination is disabled. Optional.
	IsDisabled bool
	// IsCompact - Flag indicating if the pagination is compact. Optional.
	IsCompact bool
	// ItemCount - Total number of items. Optional.
	ItemCount int
	// LastPage - The number of the last page. Optional.
	LastPage int
	// FirstPage - The number of first page where pagination starts. Optional.
	FirstPage int
	// PagesTitle - The title of a page displayed beside the page number. Optional.
	PagesTitle string
	// PagesTitlePlural - The title of a page displayed beside the page number (the plural form). Optional.
	PagesTitlePlural string
	// ToLastPage - Accessible label for the button which moves to the last page. Optional.
	ToLastPage string
	// ToPreviousPage - Accessible label for the button which moves to the previous page. Optional.
	ToPreviousPage string
	// ToNextPage - Accessible label for the button which moves to the next page. Optional.
	ToNextPage string
	// ToFirstPage - Accessible label for the button which moves to the first page. Optional.
	ToFirstPage string
	// CurrPage - Accessible label for the input displaying the current page. Optional.
	CurrPage string
	// PaginationTitle - Accessible label for the pagination component. Optional.
	PaginationTitle string
	// OfWord - Accessible label for the English word "of". Optional.
	OfWord string
	// Page - The number of the current page.
	Page any // React.ReactText 
	// PerPage - Number of items per page. Optional.
	PerPage int
	// OnSetPage - Function called when page is changed.
	OnSetPage any // OnSetPage 
	// OnNextClick - Function called when user clicks to navigate to next page. Optional.
	OnNextClick func(event any // React.SyntheticEvent , page int)
	// OnPreviousClick - Function called when user clicks to navigate to previous page. Optional.
	OnPreviousClick func(event any // React.SyntheticEvent , page int)
	// OnFirstClick - Function called when user clicks to navigate to first page. Optional.
	OnFirstClick func(event any // React.SyntheticEvent , page int)
	// OnLastClick - Function called when user clicks to navigate to last page. Optional.
	OnLastClick func(event any // React.SyntheticEvent , page int)
	// OnPageInput - Function called when user inputs page number. Optional.
	OnPageInput func(event any // React.SyntheticEvent , page int)
}

type NavigationState struct {
	// UserInputPage - Optional.
	UserInputPage any // React.ReactText 
}
