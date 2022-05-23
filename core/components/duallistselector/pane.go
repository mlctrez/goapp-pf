package duallistselector

// from file "react-core/src/components/DualListSelector/DualListSelectorPane.tsx"

type PaneProps struct {
	// ClassName - Additional classes applied to the dual list selector pane. Optional.
	ClassName string
	// Children - A dual list selector list or dual list selector tree to be rendered in the pane. Optional.
	Children any // React.ReactNode 
	// IsChosen - Flag indicating if this pane is the chosen pane. Optional.
	IsChosen bool
	// Status - Status to display above the pane. Optional.
	Status string
	// Title - Title of the pane. Optional.
	Title any // React.ReactNode 
	// SearchInput - A search input placed above the list at the top of the pane, before actions. Optional.
	SearchInput any // React.ReactNode 
	// Actions - Actions to place above the pane. Optional.
	Actions []any // React.ReactNode 
	// Id - Id of the pane. Optional.
	Id string
	// Options - Optional.
	Options []any // React.ReactNode 
	// SelectedOptions - Optional.
	SelectedOptions any /* []string | []int */
	// OnOptionSelect - Optional.
	OnOptionSelect func(e any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, index int, isChosen bool, id string, itemData any, parentData any)
	// OnOptionCheck - Optional.
	OnOptionCheck func(evt any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, isChecked bool, itemData any // DualListSelectorTreeItemData )
	// IsSearchable - Optional.
	IsSearchable bool
	// IsDisabled - Flag indicating whether the component is disabled. Optional.
	IsDisabled bool
	// OnSearch - Callback for search input. To be used when isSearchable is true. Optional.
	OnSearch func(event any // React.ChangeEvent )
	// OnSearchInputChanged - Optional.
	OnSearchInputChanged func(value string, event any // React.FormEvent )
	// FilterOption - Optional.
	FilterOption func(option any // React.ReactNode , input string) bool
	// SearchInputAriaLabel - Optional.
	SearchInputAriaLabel string
	// OnFilterUpdate - Optional.
	OnFilterUpdate func(newFilteredOptions []any // React.ReactNode , paneType string, isSearchReset bool)
}
