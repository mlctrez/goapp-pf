package dropdown

// from file "react-core/src/components/Dropdown/DropdownMenu.tsx"

type MenuProps struct {
	// Children - Anything which can be rendered as dropdown items. Optional.
	Children any // React.ReactNode 
	// ClassName - Classess applied to root element of dropdown menu. Optional.
	ClassName string
	// IsOpen - Flag to indicate if menu is opened. Optional.
	IsOpen bool
	// OpenedOnEnter - Optional.
	OpenedOnEnter bool
	// AutoFocus - Flag to indicate if the first dropdown item should gain initial focus, set false when adding
	// a specific auto-focus item (like a current selection) otherwise leave as true. Optional.
	AutoFocus bool
	// Component - Indicates which component will be used as dropdown menu. Optional.
	Component any // React.ReactNode 
	// Position - Indicates where menu will be alligned horizontally. Optional.
	Position any /* any // DropdownPosition  | "right" | "left" */
	// Alignments - Indicates how the menu will align at screen size breakpoints. Optional.
	Alignments map[string]string /* { sm:{ right, left }, md:{ right, left }, lg:{ right, left }, xl:{ right, left }, 2xl:{ right, left } } */
	// IsGrouped - Flag to indicate if menu is grouped. Optional.
	IsGrouped bool
	// SetMenuComponentRef - Optional.
	SetMenuComponentRef any
}

type MenuItem struct {
	// IsDisabled - 
	IsDisabled bool
	// Disabled - 
	Disabled bool
	// IsHovered - 
	IsHovered bool
	// Ref - 
	Ref any // HTMLElement 
}
