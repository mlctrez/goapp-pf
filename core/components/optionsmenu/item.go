package optionsmenu

// from file "react-core/src/components/OptionsMenu/OptionsMenuItem.tsx"

type ItemProps struct {
	// Children - Anything which can be rendered as an options menu item. Optional.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of an options menu item. Optional.
	ClassName string
	// IsSelected - Render options menu item as selected. Optional.
	IsSelected bool
	// IsDisabled - Render options menu item as disabled option. Optional.
	IsDisabled bool
	// OnSelect - Callback for when this options menu item is selected. Optional.
	OnSelect func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  */)
	// Id - Unique id of this options menu item. Optional.
	Id string
}
