package treeview

// from file "react-core/src/components/TreeView/TreeViewSearch.tsx"

type SearchProps struct {
	// OnSearch - Callback for search input. Optional.
	OnSearch func(event any // React.ChangeEvent )
	// Id - Id for the search input. Optional.
	Id string
	// Name - Name for the search input. Optional.
	Name string
	// AriaLabel - Accessible label for the search input. Optional.
	AriaLabel string
	// ClassName - Classes applied to the wrapper for the search input. Optional.
	ClassName string
}
