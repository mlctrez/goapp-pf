package contextselector

// from file "react-core/src/components/ContextSelector/ContextSelectorToggle.tsx"

type ToggleProps struct {
	// Id - HTML ID of toggle.
	Id string
	// ClassName - Classes applied to root element of toggle. Optional.
	ClassName string
	// ToggleText - Text that appears in the Context Selector Toggle. Optional.
	ToggleText string
	// IsOpen - Flag to indicate if menu is opened. Optional.
	IsOpen bool
	// OnToggle - Callback called when toggle is clicked. Optional.
	OnToggle func(event any, value bool)
	// OnEnter - Callback for toggle open on keyboard entry. Optional.
	OnEnter func()
	// ParentRef - Element which wraps toggle. Optional.
	ParentRef any
	// IsActive - Forces active state. Optional.
	IsActive bool
	// IsPlain - Flag to indicate the toggle has no border or background. Optional.
	IsPlain bool
	// IsText - Flag to indicate if toggle is textual toggle. Optional.
	IsText bool
}
