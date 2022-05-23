package applicationlauncher

// from file "react-core/src/components/ApplicationLauncher/ApplicationLauncher.tsx"

type Props struct {
	// ClassName - Additional element css classes. Optional.
	ClassName string
	// Direction - Display menu above or below dropdown toggle. Optional.
	Direction any /* any // DropdownDirection  | "up" | "down" */
	// Items - Array of application launcher items. Optional.
	Items []any // React.ReactNode 
	// IsDisabled - Render Application launcher toggle as disabled icon. Optional.
	IsDisabled bool
	// IsOpen - open bool. Optional.
	IsOpen bool
	// Position - Indicates where menu will be alligned horizontally. Optional.
	Position any /* any // DropdownPosition  | "right" | "left" */
	// OnSelect - Function callback called when user selects item. Optional.
	OnSelect func(event any)
	// OnToggle - Callback called when application launcher toggle is clicked. Optional.
	OnToggle func(value bool)
	// AriaLabel - Adds accessible text to the button. Required for plain buttons. Optional.
	AriaLabel string
	// IsGrouped - Flag to indicate if application launcher has groups. Optional.
	IsGrouped bool
	// ToggleIcon - Toggle Icon, optional to override the icon used for the toggle. Optional.
	ToggleIcon any // React.ReactNode 
	// Favorites - ID list of favorited ApplicationLauncherItems. Optional.
	Favorites []string
	// OnFavorite - Enables favorites. Callback called when an ApplicationLauncherItem's favorite button is clicked.
	// Optional.
	OnFavorite func(itemId string, isFavorite bool)
	// OnSearch - Enables search. Callback called when text input is entered into search box. Optional.
	OnSearch func(textInput string)
	// SearchPlaceholderText - Placeholder text for search input. Optional.
	SearchPlaceholderText string
	// SearchNoResultsText - Text for search input when no results are found. Optional.
	SearchNoResultsText string
	// SearchProps - Additional properties for search input. Optional.
	SearchProps any
	// FavoritesLabel - Label for the favorites group. Optional.
	FavoritesLabel string
	// ToggleId - ID of toggle. Optional.
	ToggleId string
}
