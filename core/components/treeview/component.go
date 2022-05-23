package treeview

// from file "react-core/src/components/TreeView/TreeView.tsx"

type DataItem struct {
	// Name - Internal content of a tree view item.
	Name any // React.ReactNode 
	// Title - Title a tree view item. Only used in Compact presentations. Optional.
	Title any // React.ReactNode 
	// Id - ID of a tree view item. Optional.
	Id string
	// Children - Child nodes of a tree view item. Optional.
	Children []any // TreeViewDataItem 
	// DefaultExpanded - Flag indicating if node is expanded by default. Optional.
	DefaultExpanded bool
	// Icon - Default icon of a tree view item. Optional.
	Icon any // React.ReactNode 
	// ExpandedIcon - Expanded icon of a tree view item. Optional.
	ExpandedIcon any // React.ReactNode 
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
	// Action - Action of a tree view item, can be a Button or Dropdown. Optional.
	Action any // React.ReactNode 
}

type Props struct {
	// Data - Data of the tree view.
	Data []any // TreeViewDataItem 
	// Id - ID of the tree view. Optional.
	Id string
	// IsNested - Flag indicating if the tree view is nested. Optional.
	IsNested bool
	// HasChecks - Flag indicating if all nodes in the tree view should have checkboxes. Optional.
	HasChecks bool
	// HasBadges - Flag indicating if all nodes in the tree view should have badges. Optional.
	HasBadges bool
	// HasGuides - Flag indicating if tree view has guide lines. Optional.
	HasGuides bool
	// Variant - Variant presentation styles for the tree view. Optional.
	Variant any /* "default" | "compact" | "compactNoBackground" */
	// Icon - Icon for all leaf or unexpanded node items. Optional.
	Icon any // React.ReactNode 
	// ExpandedIcon - Icon for all expanded node items. Optional.
	ExpandedIcon any // React.ReactNode 
	// AllExpanded - Sets the expanded state on all tree nodes, overriding default behavior and current internal
	// state. Optional.
	AllExpanded bool
	// DefaultAllExpanded - Sets the default expanded behavior. Optional.
	DefaultAllExpanded bool
	// OnSelect - Callback for item selection. Optional.
	OnSelect func(event any // React.MouseEvent , item any // TreeViewDataItem , parentItem any // TreeViewDataItem )
	// OnCheck - Callback for item checkbox selection. Optional.
	OnCheck func(event any // React.ChangeEvent , item any // TreeViewDataItem , parentItem any // TreeViewDataItem )
	// ActiveItems - Active items of tree view. Optional.
	ActiveItems []any // TreeViewDataItem 
	// ParentItem - Internal. Parent item of a TreeViewListItem. Optional.
	ParentItem any // TreeViewDataItem 
	// CompareItems - Comparison function for determining active items. Optional.
	CompareItems func(item any // TreeViewDataItem , itemToCheck any // TreeViewDataItem ) bool
	// ClassName - Class to add to add if not passed a parentItem. Optional.
	ClassName string
	// Toolbar - Toolbar to display above the tree view. Optional.
	Toolbar any // React.ReactNode 
	// UseMemo - Flag indicating the TreeView should utilize memoization to help render large data sets. Setting
	// this property requires that `activeItems` pass in an array containing every node in the selected item's
	// path. Optional.
	UseMemo bool
}
