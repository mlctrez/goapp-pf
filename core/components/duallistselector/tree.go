package duallistselector

// from file "react-core/src/components/DualListSelector/DualListSelectorTree.tsx"

type TreeItemData struct {
	// Children - Content rendered inside the dual list selector. Optional.
	Children []any // DualListSelectorTreeItemData 
	// ClassName - Additional classes applied to the dual list selector. Optional.
	ClassName string
	// DefaultExpanded - Flag indicating this option is expanded by default. Optional.
	DefaultExpanded bool
	// HasBadge - Flag indicating this option has a badge. Optional.
	HasBadge bool
	// OnOptionCheck - Callback fired when an option is checked. Optional.
	OnOptionCheck func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, isChecked bool, isChosen bool, itemData any // DualListSelectorTreeItemData )
	// Id - ID of the option.
	Id string
	// Text - Text of the option.
	Text string
	// ParentId - Parent id of an option. Optional.
	ParentId string
	// IsChecked - Checked state of the option.
	IsChecked bool
	// CheckProps - Additional properties to pass to the option checkbox. Optional.
	CheckProps any
	// BadgeProps - Additional properties to pass to the option badge. Optional.
	BadgeProps any
	// IsDisabled - Flag indicating whether the component is disabled. Optional.
	IsDisabled bool
}

type TreeProps struct {
	// Data - Data of the tree view.
	Data any /* []any // DualListSelectorTreeItemData  | func() []any // DualListSelectorTreeItemData  */
	// Id - ID of the tree view. Optional.
	Id string
	// IsNested - Optional.
	IsNested bool
	// HasBadges - Flag indicating if all options should have badges. Optional.
	HasBadges bool
	// DefaultAllExpanded - Sets the default expanded behavior. Optional.
	DefaultAllExpanded bool
	// IsDisabled - Callback fired when an option is checked. Optional.
	IsDisabled bool
	// OnOptionCheck - Optional.
	OnOptionCheck func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, isChecked bool, itemData any // DualListSelectorTreeItemData )
}
