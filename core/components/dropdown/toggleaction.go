package dropdown

// from file "react-core/src/components/Dropdown/DropdownToggleAction.tsx"

type ToggleActionProps struct {
	// ClassName - Additional classes added to the DropdownToggleAction. Optional.
	ClassName string
	// IsDisabled - Flag to show if the action button is disabled. Optional.
	IsDisabled bool
	// OnClick - A callback for when the action button is clicked. Optional.
	OnClick func(event any // React.MouseEvent )
	// Children - Element to be rendered inside the <button>. Optional.
	Children any // React.ReactNode 
	// Id - Id of the action button. Optional.
	Id string
	// AriaLabel - Aria-label of the action button. Optional.
	AriaLabel string
}
