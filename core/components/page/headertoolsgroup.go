package page

// from file "react-core/src/components/Page/PageHeaderToolsGroup.tsx"

type HeaderToolsGroupProps struct {
	// Children - Content rendered in the page header tools group.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the page header tools group. Optional.
	ClassName string
	// Visibility - Visibility at various breakpoints. Optional.
	Visibility map[string]string /* { default:{ hidden, visible }, sm:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
}
