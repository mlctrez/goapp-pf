package duallistselector

// from file "react-core/src/components/DualListSelector/DualListSelectorTreeItem.tsx"

type TreeItemProps struct {
	// Children - Content rendered inside the dual list selector. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes applied to the dual list selector. Optional.
	ClassName string
	// DefaultExpanded - Flag indicating this option is expanded by default. Optional.
	DefaultExpanded bool
	// HasBadge - Flag indicating this option has a badge. Optional.
	HasBadge bool
	// OnOptionCheck - Callback fired when an option is checked. Optional.
	OnOptionCheck func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, isChecked bool, itemData any // DualListSelectorTreeItemData )
	// Id - ID of the option.
	Id string
	// Text - Text of the option.
	Text string
	// IsChecked - Flag indicating if this open is checked. Optional.
	IsChecked bool
	// CheckProps - Additional properties to pass to the option checkbox. Optional.
	CheckProps any
	// BadgeProps - Additional properties to pass to the option badge. Optional.
	BadgeProps any
	// ItemData - Raw data of the option. Optional.
	ItemData any // DualListSelectorTreeItemData 
	// IsDisabled - Flag indicating whether the component is disabled. Optional.
	IsDisabled bool
	// UseMemo - Flag indicating the DualListSelector tree should utilize memoization to help render large data
	// sets. Optional.
	UseMemo bool
}
