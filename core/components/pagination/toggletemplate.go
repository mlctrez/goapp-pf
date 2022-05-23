package pagination

// from file "react-core/src/components/Pagination/ToggleTemplate.tsx"

type ToggleTemplateProps struct {
	// FirstIndex - The first index of the items being paginated. Optional.
	FirstIndex int
	// LastIndex - The last index of the items being paginated. Optional.
	LastIndex int
	// ItemCount - The total number of items being paginated. Optional.
	ItemCount int
	// ItemsTitle - The type or title of the items being paginated. Optional.
	ItemsTitle string
	// OfWord - The word that joins the index and itemCount/itemsTitle. Optional.
	OfWord any // React.ReactNode 
}
