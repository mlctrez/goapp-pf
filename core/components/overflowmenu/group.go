package overflowmenu

// from file "react-core/src/components/OverflowMenu/OverflowMenuGroup.tsx"

type GroupProps struct {
	// Children - Any elements that can be rendered in the menu. Optional.
	Children any
	// ClassName - Additional classes added to the OverflowMenuGroup. Optional.
	ClassName string
	// IsPersistent - Modifies the overflow menu group visibility. Optional.
	IsPersistent bool
	// GroupType - Indicates a button or icon group. Optional.
	GroupType any /* "button" | "icon" */
}
