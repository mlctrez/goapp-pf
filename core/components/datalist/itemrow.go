package datalist

// from file "react-core/src/components/DataList/DataListItemRow.tsx"

type ItemRowProps struct {
	// Children - Content rendered inside the DataListItemRow.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the DataListItemRow. Optional.
	ClassName string
	// Rowid - Id for the row item. Optional.
	Rowid string
	// WrapModifier - Determines which wrapping modifier to apply to the DataListItemRow. Optional.
	WrapModifier any /* any // DataListWrapModifier  | "nowrap" | "truncate" | "breakWord" */
}
