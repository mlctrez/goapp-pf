package datalist

// from file "react-core/src/components/DataList/DataListToggle.tsx"

type ToggleProps struct {
	// ClassName - Additional classes added to the DataList cell. Optional.
	ClassName string
	// IsExpanded - Flag to show if the expanded content of the DataList item is visible. Optional.
	IsExpanded bool
	// Id - Identify the DataList toggle number.
	Id string
	// Rowid - Id for the row. Optional.
	Rowid string
	// AriaLabelledby - Adds accessible text to the DataList toggle. Optional.
	AriaLabelledby string
	// AriaLabel - Adds accessible text to the DataList toggle. Optional.
	AriaLabel string
	// AriaControls - Allows users of some screen readers to shift focus to the controlled element. Should be
	// used when the controlled contents are not adjacent to the toggle that controls them. Optional.
	AriaControls string
}
