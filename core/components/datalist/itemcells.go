package datalist

// from file "react-core/src/components/DataList/DataListItemCells.tsx"

type ItemCellsProps struct {
	// ClassName - Additional classes added to the DataList item Content Wrapper.  Children should be one ore
	// more <DataListCell> nodes. Optional.
	ClassName string
	// DataListCells - Array of <DataListCell> nodes that are rendered one after the other. Optional.
	DataListCells any // React.ReactNode 
	// Rowid - Id for the row. Optional.
	Rowid string
}
