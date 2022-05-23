package datalist

// from file "react-core/src/components/DataList/DataListCell.tsx"

type CellProps struct {
	// Children - Content rendered inside the DataList cell. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the DataList cell. Optional.
	ClassName string
	// Width - Width (from 1-5) to the DataList cell. Optional.
	Width any /* "1" | "2" | "3" | "4" | "5" */
	// IsFilled - Enables the body Content to fill the height of the card. Optional.
	IsFilled bool
	// AlignRight - Aligns the cell content to the right of its parent. Optional.
	AlignRight bool
	// IsIcon - Set to true if the cell content is an Icon. Optional.
	IsIcon bool
	// WrapModifier - Determines which wrapping modifier to apply to the DataListCell. Optional.
	WrapModifier any /* any // DataListWrapModifier  | "nowrap" | "truncate" | "breakWord" */
}
