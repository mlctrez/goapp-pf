package toolbar

// from file "react-core/src/components/Toolbar/ToolbarFilter.tsx"

type ChipGroup struct {
	// Key - A unique key to identify this chip group category.
	Key string
	// Name - The category name to display for the chip group.
	Name string
}

type Chip struct {
	// Key - A unique key to identify this chip.
	Key string
	// Node - The ReactNode to display in the chip.
	Node any // React.ReactNode 
}

type FilterProps struct {
	// Chips - An array of strings to be displayed as chips in the expandable content. Optional.
	Chips []( any /* string | any // ToolbarChip  */ )
	// DeleteChipGroup - Callback passed by consumer used to close the entire chip group. Optional.
	DeleteChipGroup func(category any /* string | any // ToolbarChipGroup  */)
	// DeleteChip - Callback passed by consumer used to delete a chip from the chips[]. Optional.
	DeleteChip func(category any /* string | any // ToolbarChipGroup  */, chip any /* any // ToolbarChip  | string */)
	// ChipGroupExpandedText - Customizable "Show Less" text string for the chip group. Optional.
	ChipGroupExpandedText string
	// ChipGroupCollapsedText - Customizeable template string for the chip group. Use variable "${remaining}"
	// for the overflow chip count. Optional.
	ChipGroupCollapsedText string
	// Children - Content to be rendered inside the data toolbar item associated with the chip group.
	Children any // React.ReactNode 
	// CategoryName - Unique category name to be used as a label for the chip group.
	CategoryName any /* string | any // ToolbarChipGroup  */
	// ShowToolbarItem - Flag to show the toolbar item. Optional.
	ShowToolbarItem bool
}

type FilterState struct {
	// IsMounted - 
	IsMounted bool
}
