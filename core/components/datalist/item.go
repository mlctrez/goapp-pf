package datalist

// from file "react-core/src/components/DataList/DataListItem.tsx"

type ItemProps struct {
	// IsExpanded - Flag to show if the expanded content of the DataList item is visible. Optional.
	IsExpanded bool
	// Children - Content rendered inside the DataList item.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the DataList item should be either <DataListItemRow> or <DataListContent>.
	// Optional.
	ClassName string
	// AriaLabelledby - Adds accessible text to the DataList item.
	AriaLabelledby string
	// Id - Unique id for the DataList item. Optional.
	Id string
}

type ItemChildProps struct {
	// Rowid - Id for the row.
	Rowid string
}
