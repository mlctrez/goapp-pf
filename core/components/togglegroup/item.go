package togglegroup

// from file "react-core/src/components/ToggleGroup/ToggleGroupItem.tsx"

type ItemProps struct {
	// Text - Text rendered inside the toggle group item. Optional.
	Text any // React.ReactNode 
	// Icon - Icon rendered inside the toggle group item. Optional.
	Icon any // React.ReactNode 
	// ClassName - Additional classes added to the toggle group item. Optional.
	ClassName string
	// IsDisabled - Flag indicating if the toggle group item is disabled. Optional.
	IsDisabled bool
	// IsSelected - Flag indicating if the toggle group item is selected. Optional.
	IsSelected bool
	// AriaLabel - required when icon is used with no supporting text. Optional.
	AriaLabel string
	// ButtonId - Optional id for the button within the toggle group item. Optional.
	ButtonId string
	// OnChange - A callback for when the toggle group item selection changes. Optional.
	OnChange func(selected bool, event any /* any // React.MouseEvent  | any // React.KeyboardEvent  | any // MouseEvent  */)
}
