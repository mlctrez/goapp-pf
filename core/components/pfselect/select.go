package pfselect

// from file "react-core/src/components/Select/Select.tsx"

type SelectViewMoreObject struct {
	// Text - View more text.
	Text string
	// OnClick - Callback for when the view more button is clicked.
	OnClick func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  */)
}

type SelectProps struct {
	// Children - Content rendered inside the Select. Must be React.ReactElement<SelectGroupProps>[]. Optional.
	Children []any // React.ReactElement 
	// ClassName - Classes applied to the root of the Select. Optional.
	ClassName string
	// Position - Indicates where menu will be aligned horizontally. Optional.
	Position any /* any // SelectPosition  | "right" | "left" */
	// Direction - Flag specifying which direction the Select menu expands. Optional.
	Direction any /* "up" | "down" */
	// IsOpen - Flag to indicate if select is open. Optional.
	IsOpen bool
	// IsGrouped - Flag to indicate if select options are grouped. Optional.
	IsGrouped bool
	// IsPlain - Display the toggle with no border or background. Optional.
	IsPlain bool
	// IsDisabled - Flag to indicate if select is disabled. Optional.
	IsDisabled bool
	// IsCreatable - Flag to indicate if the typeahead select allows new items. Optional.
	IsCreatable bool
	// HasPlaceholderStyle - Flag indicating if placeholder styles should be applied. Optional.
	HasPlaceholderStyle bool
	// IsCreateSelectOptionObject - Optional.
	IsCreateSelectOptionObject bool
	// Validated - Value to indicate if the select is modified to show that validation state. If set to success,
	// select will be modified to indicate valid state. If set to error, select will be modified to indicate
	// error state. If set to warning, select will be modified to indicate warning state. Optional.
	Validated any /* "success" | "warning" | "error" | "default" */
	// LoadingVariant - Optional.
	LoadingVariant any /* "spinner" | any // SelectViewMoreObject  */
	// CreateText - Text displayed in typeahead select to prompt the user to create an item. Optional.
	CreateText string
	// PlaceholderText - Title text of Select. Optional.
	PlaceholderText any /* string | any // React.ReactNode  */
	// NoResultsFoundText - Text to display in typeahead select when no results are found. Optional.
	NoResultsFoundText string
	// Selections - Array of selected items for multi select variants. Optional.
	Selections any /* string | any // SelectOptionObject  | []( any /* string | any // SelectOptionObject  */ ) */
	// IsCheckboxSelectionBadgeHidden - Flag indicating if selection badge should be hidden for checkbox variant,default
	// false. Optional.
	IsCheckboxSelectionBadgeHidden bool
	// ToggleId - Id for select toggle element. Optional.
	ToggleId string
	// AriaLabel - Adds accessible text to Select. Optional.
	AriaLabel string
	// AriaLabelledby - Id of label for the Select aria-labelledby. Optional.
	AriaLabelledby string
	// AriaDescribedby - Id of div for the select aria-labelledby. Optional.
	AriaDescribedby string
	// AriaInvalid - Flag indicating if the select is an invalid state. Optional.
	AriaInvalid bool
	// TypeAheadAriaLabel - Label for input field of type ahead select variants. Optional.
	TypeAheadAriaLabel string
	// ClearSelectionsAriaLabel - Label for clear selection button of type ahead select variants. Optional.
	ClearSelectionsAriaLabel string
	// ToggleAriaLabel - Label for toggle of type ahead select variants. Optional.
	ToggleAriaLabel string
	// RemoveSelectionAriaLabel - Label for remove chip button of multiple type ahead select variant. Optional.
	RemoveSelectionAriaLabel string
	// Favorites - ID list of favorited select items. Optional.
	Favorites []string
	// FavoritesLabel - Label for the favorites group. Optional.
	FavoritesLabel string
	// OnFavorite - Enables favorites. Callback called when a select options's favorite button is clicked. Optional.
	OnFavorite func(itemId string, isFavorite bool)
	// OnSelect - Callback for selection behavior. Optional.
	OnSelect func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  */, value any /* string | any // SelectOptionObject  */, isPlaceholder bool)
	// OnToggle - Callback for toggle button behavior.
	OnToggle func(isExpanded bool, event any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  | any // Event  */)
	// OnBlur - Callback for toggle blur. Optional.
	OnBlur func(event any)
	// OnClear - Callback for typeahead clear button. Optional.
	OnClear func(event any // React.MouseEvent )
	// OnFilter - Optional callback for custom filtering. Optional.
	OnFilter func(e any /* any // React.ChangeEvent  | "" */, value string) any /* []any // React.ReactElement  | undefined */
	// OnCreateOption - Optional callback for newly created options. Optional.
	OnCreateOption func(newOptionValue string)
	// OnTypeaheadInputChanged - Optional event handler called each time the value in the typeahead input changes.
	// Optional.
	OnTypeaheadInputChanged func(value string)
	// Variant - Variant of rendered Select. Optional.
	Variant any /* "single" | "checkbox" | "typeahead" | "typeaheadmulti" */
	// Width - Width of the select container as a number of px or string percentage. Optional.
	Width any /* string | int */
	// MaxHeight - Max height of the select container as a number of px or string percentage. Optional.
	MaxHeight any /* string | int */
	// ToggleIcon - Icon element to render inside the select toggle. Optional.
	ToggleIcon any // React.ReactElement 
	// CustomContent - Custom content to render in the select menu.  If this prop is defined, the variant prop
	// will be ignored and the select will render with a single select toggle. Optional.
	CustomContent any // React.ReactNode 
	// HasInlineFilter - Flag indicating if select should have an inline text input for filtering. Optional.
	HasInlineFilter bool
	// InlineFilterPlaceholderText - Placeholder text for inline filter. Optional.
	InlineFilterPlaceholderText string
	// CustomBadgeText - Custom text for select badge. Optional.
	CustomBadgeText any /* string | int */
	// InputIdPrefix - Prefix for the id of the input in the checkbox select variant. Optional.
	InputIdPrefix string
	// InputAutoComplete - Value for the typeahead and inline filtering input autocomplete attribute. When targeting
	// Chrome this property should be a random string. Optional.
	InputAutoComplete string
	// ChipGroupProps - Optional props to pass to the chip group in the typeaheadmulti variant. Optional.
	ChipGroupProps any // Omit 
	// ChipGroupComponent - Optional props to render custom chip group in the typeaheadmulti variant. Optional.
	ChipGroupComponent any // React.ReactNode 
	// IsInputValuePersisted - Flag for retaining keyboard-entered value in typeahead text field when focus leaves
	// input away. Optional.
	IsInputValuePersisted bool
	// IsInputFilterPersisted - Optional.
	IsInputFilterPersisted bool
	// ShouldResetOnSelect - Flag indicating the typeahead input value should reset upon selection. Optional.
	ShouldResetOnSelect bool
	// Footer - Content rendered in the footer of the select menu. Optional.
	Footer any // React.ReactNode 
}

type SelectState struct {
	// FocusFirstOption - 
	FocusFirstOption bool
	// TypeaheadInputValue - 
	TypeaheadInputValue any /* string | "" */
	// TypeaheadFilteredChildren - 
	TypeaheadFilteredChildren []any // React.ReactNode 
	// FavoritesGroup - 
	FavoritesGroup []any // React.ReactNode 
	// TypeaheadCurrIndex - 
	TypeaheadCurrIndex int
	// CreatableValue - 
	CreatableValue string
	// TabbedIntoFavoritesMenu - 
	TabbedIntoFavoritesMenu bool
	// TypeaheadStoredIndex - 
	TypeaheadStoredIndex int
	// OuiaStateId - 
	OuiaStateId string
	// ViewMoreNextIndex - 
	ViewMoreNextIndex int
}
