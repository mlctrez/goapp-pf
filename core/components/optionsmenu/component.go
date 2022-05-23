package optionsmenu

// from file "react-core/src/components/OptionsMenu/OptionsMenu.tsx"

type Props struct {
	// ClassName - Classes applied to root element of the options menu. Optional.
	ClassName string
	// Id - Id of the root element of the options menu.
	Id string
	// MenuItems - Array of OptionsMenuItem and/or OptionMenuGroup nodes that will be rendered in the options
	// menu list.
	MenuItems []any // React.ReactNode 
	// Toggle - Either an OptionsMenuToggle or an OptionsMenuToggleWithText to use to toggle the options menu.
	Toggle any // React.ReactElement 
	// IsPlain - Flag to indicate the toggle has no border or background. Optional.
	IsPlain bool
	// IsOpen - Flag to indicate if menu is open. Optional.
	IsOpen bool
	// IsText - Flag to indicate if toggle is textual toggle. Optional.
	IsText bool
	// IsGrouped - Flag to indicate if menu is groupped. Optional.
	IsGrouped bool
	// Position - Indicates where menu will be aligned horizontally. Optional.
	Position any /* "right" | "left" */
	// Direction - Menu will open up or open down from the options menu toggle. Optional.
	Direction any /* "up" | "down" */
}
