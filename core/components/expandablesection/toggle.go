package expandablesection

// from file "react-core/src/components/ExpandableSection/ExpandableSectionToggle.tsx"

type ToggleProps struct {
	// Children - Content rendered inside the expandable toggle. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the expandable toggle. Optional.
	ClassName string
	// IsExpanded - Flag indicating if the expandable section is expanded. Optional.
	IsExpanded bool
	// OnToggle - Callback function to toggle the expandable content. Optional.
	OnToggle func(isExpanded bool)
	// ContentId - ID of the toggle's respective expandable section content. Optional.
	ContentId string
	// Direction - Direction the toggle arrow should point when the expandable section is expanded. Optional.
	Direction any /* "up" | "down" */
}
