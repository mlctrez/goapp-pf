package datalist

// from file "react-core/src/components/DataList/DataListContent.tsx"

type ContentProps struct {
	// Children - Content rendered inside the DataList item. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the DataList cell. Optional.
	ClassName string
	// Id - Identify the DataListContent item. Optional.
	Id string
	// Rowid - Id for the row. Optional.
	Rowid string
	// IsHidden - Flag to show if the expanded content of the DataList item is visible. Optional.
	IsHidden bool
	// HasNoPadding - Flag to remove padding from the expandable content. Optional.
	HasNoPadding bool
	// AriaLabel - Adds accessible text to the DataList toggle.
	AriaLabel string
}
