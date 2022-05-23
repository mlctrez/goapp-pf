package toolbar

// from file "react-core/src/components/Toolbar/ToolbarToggleGroup.tsx"

type ToggleGroupProps struct {
	// ToggleIcon - An icon to be rendered when the toggle group has collapsed down.
	ToggleIcon any // React.ReactNode 
	// Breakpoint - Controls when filters are shown and when the toggle button is hidden.
	Breakpoint any /* "md" | "lg" | "xl" | "2xl" */
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
}
