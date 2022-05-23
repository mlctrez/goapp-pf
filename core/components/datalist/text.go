package datalist

// from file "react-core/src/components/DataList/DataListText.tsx"

type TextProps struct {
	// Children - Content rendered within the data list text. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the data list text. Optional.
	ClassName string
	// Component - Determines which element to render as a data list text. Usually div or span. Optional.
	Component any // React.ReactNode 
	// WrapModifier - Determines which wrapping modifier to apply to the data list text. Optional.
	WrapModifier any /* any // DataListWrapModifier  | "nowrap" | "truncate" | "breakWord" */
	// Tooltip - text to display on the tooltip. Optional.
	Tooltip string
	// OnMouseEnter - callback used to create the tooltip if text is truncated. Optional.
	OnMouseEnter func(event any)
}
