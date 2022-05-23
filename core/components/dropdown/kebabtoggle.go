package dropdown

// from file "react-core/src/components/Dropdown/KebabToggle.tsx"

type KebabToggleProps struct {
	// Id - HTML ID of dropdown toggle. Optional.
	Id string
	// Children - Anything which can be rendered as dropdown toggle. Optional.
	Children any // React.ReactNode 
	// ClassName - Classess applied to root element of dropdown toggle. Optional.
	ClassName string
	// IsOpen - Flag to indicate if menu is opened. Optional.
	IsOpen bool
	// AriaLabel - Label Toggle button. Optional.
	AriaLabel string
	// OnToggle - Callback called when toggle is clicked. Optional.
	OnToggle func(value bool, event any)
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
	// Type - Type to put on the button. Optional.
	Type any /* "button" | "submit" | "reset" */
	// BubbleEvent - Allows selecting toggle to select parent. Optional.
	BubbleEvent bool
}
