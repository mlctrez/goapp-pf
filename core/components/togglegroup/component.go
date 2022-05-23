package togglegroup

// from file "react-core/src/components/ToggleGroup/ToggleGroup.tsx"

type Props struct {
	// Children - Content rendered inside the toggle group. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the toggle group. Optional.
	ClassName string
	// IsCompact - Modifies the toggle group to include compact styling. Optional.
	IsCompact bool
	// AreAllGroupsDisabled - Disable all toggle group items under this component. Optional.
	AreAllGroupsDisabled bool
	// AriaLabel - Accessible label for the toggle group. Optional.
	AriaLabel string
}
