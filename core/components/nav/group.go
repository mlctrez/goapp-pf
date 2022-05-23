package nav

// from file "react-core/src/components/Nav/NavGroup.tsx"

type GroupProps struct {
	// Title - Title shown for the group. Optional.
	Title string
	// Children - Anything that can be rendered inside of the group. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the container. Optional.
	ClassName string
	// Id - Identifier to use for the section aria label. Optional.
	Id string
}
