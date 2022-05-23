package text

// from file "react-core/src/components/Text/TextList.tsx"

type ListProps struct {
	// Children - Content rendered within the TextList. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the TextList. Optional.
	ClassName string
	// Component - The text list component. Optional.
	Component any /* "ul" | "ol" | "dl" */
}
