package optionsmenu

// from file "react-core/src/components/OptionsMenu/OptionsMenuToggle.tsx"

type ToggleProps struct {
	// ParentId - Id of the parent options menu component. Optional.
	ParentId string
	// OnToggle - Callback for when this options menu is toggled. Optional.
	OnToggle func(isOpen bool)
	// IsOpen - Flag to indicate if menu is open. Optional.
	IsOpen bool
	// IsPlain - Flag to indicate if the button is plain. Optional.
	IsPlain bool
	// IsSplitButton - Optional.
	IsSplitButton bool
	// IsActive - Forces display of the active state of the options menu. Optional.
	IsActive bool
	// IsDisabled - Disables the options menu toggle. Optional.
	IsDisabled bool
	// HideCaret - hide the toggle caret. Optional.
	HideCaret bool
	// AriaLabel - Provides an accessible name for the button when an icon is used instead of text. Optional.
	AriaLabel string
	// OnEnter - Optional.
	OnEnter func(event any // React.MouseEvent )
	// ParentRef - Optional.
	ParentRef any // HTMLElement 
	// ToggleTemplate - Content to be rendered in the options menu toggle button. Optional.
	ToggleTemplate any // React.ReactNode 
}
