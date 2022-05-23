package optionsmenu

// from file "react-core/src/components/OptionsMenu/OptionsMenuToggleWithText.tsx"

type ToggleWithTextProps struct {
	// ParentId - Id of the parent options menu component. Optional.
	ParentId string
	// ToggleText - Content to be rendered inside the options menu toggle as text or another non-interactive
	// element.
	ToggleText any // React.ReactNode 
	// ToggleTextClassName - classes to be added to the options menu toggle text. Optional.
	ToggleTextClassName string
	// ToggleButtonContents - Content to be rendered inside the options menu toggle button. Optional.
	ToggleButtonContents any // React.ReactNode 
	// ToggleButtonContentsClassName - Classes to be added to the options menu toggle button. Optional.
	ToggleButtonContentsClassName string
	// OnToggle - Callback for when this options menu is toggled. Optional.
	OnToggle func(event bool)
	// OnEnter - Inner function to indicate open on Enter. Optional.
	OnEnter func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  */)
	// IsOpen - Flag to indicate if menu is open. Optional.
	IsOpen bool
	// IsPlain - Flag to indicate if the button is plain. Optional.
	IsPlain bool
	// IsActive - Forces display of the active state of the options menu button. Optional.
	IsActive bool
	// IsDisabled - Disables the options menu toggle. Optional.
	IsDisabled bool
	// ParentRef - Optional.
	ParentRef any // React.RefObject 
	// AriaHaspopup - Indicates that the element has a popup context menu or sub-level menu. Optional.
	AriaHaspopup any /* bool | "dialog" | "menu" | "listbox" | "tree" | "grid" */
	// AriaLabel - Provides an accessible name for the button when an icon is used instead of text. Optional.
	AriaLabel string
	// IsText - Optional.
	IsText bool
	// GetMenuRef - Optional.
	GetMenuRef func() any // HTMLElement 
}
