package toolbar

// from file "react-core/src/components/Toolbar/ToolbarItem.tsx"

type ItemProps struct {
	// ClassName - Classes applied to root element of the data toolbar item. Optional.
	ClassName string
	// Variant - A type modifier which modifies spacing specifically depending on the type of item. Optional.
	Variant any /* any // ToolbarItemVariant  | "bulk-select" | "overflow-menu" | "pagination" | "search-filter" | "label" | "chip-group" | "separator" | "expand-all" */
	// Visibility - Visibility at various breakpoints. Optional.
	Visibility map[string]string /* { default:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// Visiblity - Optional.
	Visiblity map[string]string /* { default:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// Alignment - Alignment at various breakpoints. Optional.
	Alignment map[string]string /* { default:{ alignRight, alignLeft }, md:{ alignRight, alignLeft }, lg:{ alignRight, alignLeft }, xl:{ alignRight, alignLeft }, 2xl:{ alignRight, alignLeft } } */
	// Spacer - Spacers at various breakpoints. Optional.
	Spacer map[string]string /* { default:{ spacerNone, spacerSm, spacerMd, spacerLg }, md:{ spacerNone, spacerSm, spacerMd, spacerLg }, lg:{ spacerNone, spacerSm, spacerMd, spacerLg }, xl:{ spacerNone, spacerSm, spacerMd, spacerLg }, 2xl:{ spacerNone, spacerSm, spacerMd, spacerLg } } */
	// Widths - Widths at various breakpoints. Optional.
	Widths map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
	// Id - id for this data toolbar item. Optional.
	Id string
	// IsAllExpanded - Flag indicating if the expand-all variant is expanded or not. Optional.
	IsAllExpanded bool
	// Children - Content to be rendered inside the data toolbar item. Optional.
	Children any // React.ReactNode 
}
