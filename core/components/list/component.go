package list

// from file "react-core/src/components/List/List.tsx"

type Props struct {
	// Children - Anything that can be rendered inside of the list. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the list. Optional.
	ClassName string
	// Variant - Adds list variant styles. Optional.
	Variant any // ListVariant.inline 
	// IsBordered - Modifies the list to add borders between items. Optional.
	IsBordered bool
	// IsPlain - Modifies the list to include plain styling. Optional.
	IsPlain bool
	// IconSize - Modifies the size of the icons in the list. Optional.
	IconSize any /* "default" | "large" */
	// Type - Sets the way items are numbered if variant is set to ordered. Optional.
	Type any // OrderType 
	// Component - Optional.
	Component any /* "ol" | "ul" */
}
