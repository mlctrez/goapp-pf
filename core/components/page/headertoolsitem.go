package page

// from file "react-core/src/components/Page/PageHeaderToolsItem.tsx"

type HeaderToolsItemProps struct {
	// Children - Content rendered in page header tools item.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the page header tools item. Optional.
	ClassName string
	// Id - HTML id of the PageHeaderToolsItem. Optional.
	Id string
	// Visibility - Visibility at various breakpoints. Optional.
	Visibility map[string]string /* { default:{ hidden, visible }, sm:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// IsSelected - True to make an icon button appear selected. Optional.
	IsSelected bool
}
