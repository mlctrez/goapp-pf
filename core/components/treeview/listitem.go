package treeview

// from file "react-core/src/components/TreeView/TreeViewListItem.tsx"

type CheckProps struct {
	// Checked - Optional.
	Checked any /* bool | "" */
}

type ListItemProps struct {
	// Name - Internal content of a tree view item.
	Name any // React.ReactNode 
	// Title - Title a tree view item.
	Title any // React.ReactNode 
	// Id - ID of a tree view item. Optional.
	Id string
	// IsExpanded - Flag indicating if the node is expanded, overrides internal state. Optional.
	IsExpanded bool
	// DefaultExpanded - Flag indicating if node is expanded by default. Optional.
	DefaultExpanded bool
	// Children - Child nodes of a tree view item. Optional.
	Children any // React.ReactNode 
	// OnSelect - Callback for item selection. Note: calling event.preventDefault() will prevent the node from
	// toggling. Optional.
	OnSelect func(event any // React.MouseEvent , item any // TreeViewDataItem , parent any // TreeViewDataItem )
	// OnCheck - Callback for item checkbox selection. Optional.
	OnCheck func(event any // React.ChangeEvent , item any // TreeViewDataItem , parent any // TreeViewDataItem )
	// HasCheck - Flag indicating if a tree view item has a checkbox. Optional.
	HasCheck bool
	// CheckProps - Additional properties of the tree view item checkbox. Optional.
	CheckProps any // TreeViewCheckProps 
	// HasBadge - Flag indicating if a tree view item has a badge. Optional.
	HasBadge bool
	// CustomBadgeContent - Optional prop for custom badge. Optional.
	CustomBadgeContent any // React.ReactNode 
	// BadgeProps - Additional properties of the tree view item badge. Optional.
	BadgeProps any
	// IsCompact - Flag indicating if the tree view is using a compact variation. Optional.
	IsCompact bool
	// ActiveItems - Active items of tree view. Optional.
	ActiveItems []any // TreeViewDataItem 
	// ItemData - Data structure of tree view item. Optional.
	ItemData any // TreeViewDataItem 
	// ParentItem - Parent item of tree view item. Optional.
	ParentItem any // TreeViewDataItem 
	// Icon - Default icon of a tree view item. Optional.
	Icon any // React.ReactNode 
	// ExpandedIcon - Expanded icon of a tree view item. Optional.
	ExpandedIcon any // React.ReactNode 
	// Action - Action of a tree view item, can be a Button or Dropdown. Optional.
	Action any // React.ReactNode 
	// CompareItems - Callback for item comparison function. Optional.
	CompareItems func(item any // TreeViewDataItem , itemToCheck any // TreeViewDataItem ) bool
	// UseMemo - Flag indicating the TreeView should utilize memoization to help render large data sets. Setting
	// this property requires that `activeItems` pass in an array containing every node in the selected item's
	// path. Optional.
	UseMemo bool
}
