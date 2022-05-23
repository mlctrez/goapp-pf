package dropdown

// from file "react-core/src/components/Dropdown/Toggle.tsx"

type ToggleProps struct {
	// Id - HTML ID of dropdown toggle.
	Id string
	// Type - Type to put on the button. Optional.
	Type any /* "button" | "submit" | "reset" */
	// Children - Anything which can be rendered as dropdown toggle. Optional.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of dropdown toggle. Optional.
	ClassName string
	// IsOpen - Flag to indicate if menu is opened. Optional.
	IsOpen bool
	// OnToggle - Callback called when toggle is clicked. Optional.
	OnToggle func(isOpen bool, event any /* any // MouseEvent  | any // TouchEvent  | any // KeyboardEvent  | any // React.KeyboardEvent  | any // React.MouseEvent  */)
	// OnEnter - Callback called when the Enter key is pressed. Optional.
	OnEnter func()
	// ParentRef - Element which wraps toggle. Optional.
	ParentRef any
	// GetMenuRef - The menu element. Optional.
	GetMenuRef func() any // HTMLElement 
	// IsActive - Forces active state. Optional.
	IsActive bool
	// IsDisabled - Disables the dropdown toggle. Optional.
	IsDisabled bool
	// IsPlain - Display the toggle with no border or background. Optional.
	IsPlain bool
	// IsText - Display the toggle in text only mode. Optional.
	IsText bool
	// IsPrimary - Optional.
	IsPrimary bool
	// IsSplitButton - Style the toggle as a child of a split button. Optional.
	IsSplitButton bool
	// ToggleVariant - Alternate styles for the dropdown toggle button. Optional.
	ToggleVariant any /* "primary" | "secondary" | "default" */
	// AriaHaspopup - Flag for aria popup. Optional.
	AriaHaspopup any /* bool | "listbox" | "menu" | "dialog" | "grid" | "tree" */
	// BubbleEvent - Allows selecting toggle to select parent. Optional.
	BubbleEvent bool
}
