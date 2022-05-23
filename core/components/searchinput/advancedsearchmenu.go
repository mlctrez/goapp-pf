package searchinput

// from file "react-core/src/components/SearchInput/AdvancedSearchMenu.tsx"

type AdvancedSearchMenuProps struct {
	// Value - Value of the search input. Optional.
	Value string
	// ParentRef - Ref of the div wrapping the whole search input *. Optional.
	ParentRef any // React.RefObject 
	// ParentInputRef - Ref of the input element within the search input*. Optional.
	ParentInputRef any // React.RefObject 
	// GetAttrValueMap - Function which builds an attribute-value map by parsing the value in the search input.
	// Optional.
	GetAttrValueMap func() map[string]string /* { [key: string]:string } */
	// OnSearch - A callback for when the search button clicked changes. Optional.
	OnSearch func(value string, event any // React.SyntheticEvent , attrValueMap map[string]string /* { [key: string]:string } */)
	// OnClear - A callback for when the user clicks the clear button. Optional.
	OnClear func(event any // React.SyntheticEvent )
	// OnChange - A callback for when the input value changes. Optional.
	OnChange func(value string, event any // React.FormEvent )
	// OnToggleAdvancedMenu - Function called to toggle the advanced search menu. Optional.
	OnToggleAdvancedMenu func(e any // React.SyntheticEvent )
	// IsSearchMenuOpen - Flag for toggling the open/close state of the advanced search menu. Optional.
	IsSearchMenuOpen bool
	// ResetButtonLabel - Label for the buttons which reset the advanced search form and clear the search input.
	// Optional.
	ResetButtonLabel string
	// SubmitSearchButtonLabel - Label for the buttons which called the onSearch event handler. Optional.
	SubmitSearchButtonLabel string
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
