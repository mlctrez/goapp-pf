package optionsmenu

// from file "react-core/src/components/OptionsMenu/OptionsMenuItemGroup.tsx"

type ItemGroupProps struct {
	// Children - Content to be rendered in the options menu items component. Optional.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of the options menu items group. Optional.
	ClassName string
	// AriaLabel - Provides an accessible name for the options menu items group. Optional.
	AriaLabel string
	// GroupTitle - Optional title for the options menu items group. Optional.
	GroupTitle any /* string | any // React.ReactNode  */
	// HasSeparator - Flag indicating this options menu items group will be followed by a horizontal separator.
	// Optional.
	HasSeparator bool
}
