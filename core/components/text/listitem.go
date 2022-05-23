package text

// from file "react-core/src/components/Text/TextListItem.tsx"

type ListItemProps struct {
	// Children - Content rendered within the TextListItem. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the TextListItem. Optional.
	ClassName string
	// Component - The text list item component. Optional.
	Component any /* "li" | "dt" | "dd" */
}
