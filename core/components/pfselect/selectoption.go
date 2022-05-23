package pfselect

// from file "react-core/src/components/Select/SelectOption.tsx"

type SelectOptionObject struct {
	// ToString - Function returns a string to represent the select option object.
	ToString string
	// CompareTo - Function returns a true if the passed in select option is equal to this select option object,
	// false otherwise. Optional.
	CompareTo bool
}

type SelectOptionProps struct {
	// Children - Optional alternate display for the option. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Select Option. Optional.
	ClassName string
	// Description - Description of the item for single and both typeahead select variants. Optional.
	Description any // React.ReactNode 
	// ItemCount - Number of items matching the option. Optional.
	ItemCount int
	// Index - Optional.
	Index int
	// Component - Indicates which component will be used as select item. Optional.
	Component any // React.ReactNode 
	// Value - The value for the option, can be a string or select option object.
	Value any /* string | any // SelectOptionObject  */
	// IsDisabled - Flag indicating if the option is disabled. Optional.
	IsDisabled bool
	// IsPlaceholder - Flag indicating if the option acts as a placeholder. Optional.
	IsPlaceholder bool
	// IsNoResultsOption - Flag indicating if the option acts as a "no results" indicator. Optional.
	IsNoResultsOption bool
	// IsSelected - Optional.
	IsSelected bool
	// IsChecked - Optional.
	IsChecked bool
	// IsFocused - Flag forcing the focused state. Optional.
	IsFocused bool
	// SendRef - Optional.
	SendRef func(ref any // React.ReactNode , favoriteRef any // React.ReactNode , optionContainerRef any // React.ReactNode , index int)
	// KeyHandler - Optional.
	KeyHandler func(index int, innerIndex int, position string)
	// OnClick - Optional callback for click event. Optional.
	OnClick func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  */)
	// InputId - Id of the checkbox input. Optional.
	InputId string
	// IsFavorite - Optional.
	IsFavorite bool
	// AriaIsFavoriteLabel - Aria label text for favoritable button when favorited. Optional.
	AriaIsFavoriteLabel string
	// AriaIsNotFavoriteLabel - Aria label text for favoritable button when not favorited. Optional.
	AriaIsNotFavoriteLabel string
	// Id - ID of the item. Required for tracking favorites. Optional.
	Id string
	// IsLoad - Optional.
	IsLoad bool
	// IsLoading - Optional.
	IsLoading bool
	// SetViewMoreNextIndex - Optional.
	SetViewMoreNextIndex func()
	// IsLastOptionBeforeFooter - Optional.
	IsLastOptionBeforeFooter func(index int) bool
	// IsGrouped - Optional.
	IsGrouped bool
}
