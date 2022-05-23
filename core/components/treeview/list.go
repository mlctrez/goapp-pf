package treeview

// from file "react-core/src/components/TreeView/TreeViewList.tsx"

type ListProps struct {
	// IsNested - Flag indicating if the tree view is nested under another tree view. Optional.
	IsNested bool
	// Toolbar - Toolbar to display above the tree view. Optional.
	Toolbar any // React.ReactNode 
	// Children - Child nodes of the current tree view.
	Children any // React.ReactNode 
}
