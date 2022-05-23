package accordion

// from file "react-core/src/components/Accordion/AccordionToggle.tsx"

type ToggleProps struct {
	// Children - Content rendered inside the Accordion toggle. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Accordion Toggle. Optional.
	ClassName string
	// IsExpanded - Flag to show if the expanded content of the Accordion item is visible. Optional.
	IsExpanded bool
	// Id - Identify the Accordion toggle number.
	Id string
	// Component - Container to override the default for toggle. Optional.
	Component any // React.ElementType 
}
