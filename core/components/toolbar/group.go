package toolbar

// from file "react-core/src/components/Toolbar/ToolbarGroup.tsx"

type GroupProps struct {
	// ClassName - Classes applied to root element of the data toolbar group. Optional.
	ClassName string
	// Variant - A type modifier which modifies spacing specifically depending on the type of group. Optional.
	Variant any /* any // ToolbarGroupVariant  | "filter-group" | "icon-button-group" | "button-group" */
	// Visibility - Visibility at various breakpoints. Optional.
	Visibility map[string]string /* { default:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// Visiblity - Optional.
	Visiblity map[string]string /* { default:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// Alignment - Alignment at various breakpoints. Optional.
	Alignment map[string]string /* { default:{ alignRight, alignLeft }, md:{ alignRight, alignLeft }, lg:{ alignRight, alignLeft }, xl:{ alignRight, alignLeft }, 2xl:{ alignRight, alignLeft } } */
	// Spacer - Spacers at various breakpoints. Optional.
	Spacer map[string]string /* { default:{ spacerNone, spacerSm, spacerMd, spacerLg }, md:{ spacerNone, spacerSm, spacerMd, spacerLg }, lg:{ spacerNone, spacerSm, spacerMd, spacerLg }, xl:{ spacerNone, spacerSm, spacerMd, spacerLg }, 2xl:{ spacerNone, spacerSm, spacerMd, spacerLg } } */
	// SpaceItems - Space items at various breakpoints. Optional.
	SpaceItems map[string]string /* { default:{ spaceItemsNone, spaceItemsSm, spaceItemsMd, spaceItemsLg }, md:{ spaceItemsNone, spaceItemsSm, spaceItemsMd, spaceItemsLg }, lg:{ spaceItemsNone, spaceItemsSm, spaceItemsMd, spaceItemsLg }, xl:{ spaceItemsNone, spaceItemsSm, spaceItemsMd, spaceItemsLg }, 2xl:{ spaceItemsNone, spaceItemsSm, spaceItemsMd, spaceItemsLg } } */
	// Children - Content to be rendered inside the data toolbar group. Optional.
	Children any // React.ReactNode 
	// InnerRef - Reference to pass to this group if it has .pf-m-chip-container modifier. Optional.
	InnerRef any // React.RefObject 
}
