package dropdown

// from file "react-core/src/components/Dropdown/Dropdown.tsx"

type Props struct {
	// Children - Anything which can be rendered in a dropdown. Optional.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of dropdown. Optional.
	ClassName string
	// DropdownItems - Array of DropdownItem nodes that will be rendered in the dropdown Menu list. Optional.
	DropdownItems []any
	// IsOpen - Flag to indicate if menu is opened. Optional.
	IsOpen bool
	// IsPlain - Display the toggle with no border or background. Optional.
	IsPlain bool
	// IsText - Display the toggle in text only mode. Optional.
	IsText bool
	// IsFullHeight - Flag indicating that the dropdown should expand to full height. Optional.
	IsFullHeight bool
	// Position - Indicates where menu will be aligned horizontally. Optional.
	Position any /* any // DropdownPosition  | "right" | "left" */
	// Alignments - Indicates how the menu will align at screen size breakpoints. Default alignment is set via
	// the position property. Optional.
	Alignments map[string]string /* { sm:{ right, left }, md:{ right, left }, lg:{ right, left }, xl:{ right, left }, 2xl:{ right, left } } */
	// Direction - Display menu above or below dropdown toggle. Optional.
	Direction any /* any // DropdownDirection  | "up" | "down" */
	// IsGrouped - Flag to indicate if dropdown has groups. Optional.
	IsGrouped bool
	// Toggle - Toggle for the dropdown, examples: <DropdownToggle> or <DropdownToggleCheckbox>.
	Toggle any // React.ReactElement 
	// OnSelect - Function callback called when user selects item. Optional.
	OnSelect func(event any // React.SyntheticEvent )
	// AutoFocus - Flag to indicate if the first dropdown item should gain initial focus, set false when adding
	// a specific auto-focus item (like a current selection) otherwise leave as true. Optional.
	AutoFocus bool
	// ContextProps - Props for extreme customization of dropdown. Optional.
	ContextProps typeof DropdownContext
}
