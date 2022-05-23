package searchinput

// from file "react-core/src/components/SearchInput/SearchInput.tsx"

type SearchAttribute struct {
	// Attr - The search attribute's value to be provided in the search input's query string. It should have
	// no spaces and be unique for every attribute.
	Attr string
	// Display - The search attribute's display name. It is used to label the field in the advanced search menu.
	Display any // React.ReactNode 
}

type Props struct {
	// ClassName - Additional classes added to the banner. Optional.
	ClassName string
	// Value - Value of the search input. Optional.
	Value string
	// IsDisabled - Flag indicating if search input is disabled. Optional.
	IsDisabled bool
	// AriaLabel - An accessible label for the search input. Optional.
	AriaLabel string
	// Placeholder - placeholder text of the search input. Optional.
	Placeholder string
	// InnerRef - Optional.
	InnerRef any // React.RefObject 
	// OnChange - A callback for when the input value changes. Optional.
	OnChange func(value string, event any // React.FormEvent )
	// Hint - A suggestion for autocompleting. Optional.
	Hint string
	// OnSearch - A callback for when the search button clicked changes. Optional.
	OnSearch func(value string, event any // React.SyntheticEvent , attrValueMap map[string]string /* { [key: string]:string } */)
	// OnClear - A callback for when the user clicks the clear button. Optional.
	OnClear func(event any // React.SyntheticEvent )
	// ResetButtonLabel - Label for the buttons which reset the advanced search form and clear the search input.
	// Optional.
	ResetButtonLabel string
	// SubmitSearchButtonLabel - Label for the buttons which called the onSearch event handler. Optional.
	SubmitSearchButtonLabel string
	// OnToggleAdvancedSearch - A callback for when the open advanced search button is clicked. Optional.
	OnToggleAdvancedSearch func(event any // React.SyntheticEvent , isOpen bool)
	// IsAdvancedSearchOpen - A flag for controlling the open state of a custom advanced search implementation.
	// Optional.
	IsAdvancedSearchOpen bool
	// OpenMenuButtonAriaLabel - Label for the button which opens the advanced search form menu. Optional.
	OpenMenuButtonAriaLabel string
	// OnNextClick - Function called when user clicks to navigate to next result. Optional.
	OnNextClick func(event any // React.SyntheticEvent )
	// OnPreviousClick - Function called when user clicks to navigate to previous result. Optional.
	OnPreviousClick func(event any // React.SyntheticEvent )
	// ResultsCount - The number of search results returned. Either a total number of results, or a string representing
	// the current result over the total number of results. i.e. "1 / 5". Optional.
	ResultsCount any /* int | string */
	// Attributes - Array of attribute values used for dynamically generated advanced search. Optional.
	Attributes any /* []string | []any // SearchAttribute  */
	// FormAdditionalItems - Optional.
	FormAdditionalItems any // React.ReactNode 
	// HasWordsAttrLabel - Attribute label for strings unassociated with one of the provided listed attributes.
	// Optional.
	HasWordsAttrLabel any // React.ReactNode 
	// AdvancedSearchDelimiter - Delimiter in the query string for pairing attributes with search values. Required
	// whenever attributes are passed as props. Optional.
	AdvancedSearchDelimiter string
}
