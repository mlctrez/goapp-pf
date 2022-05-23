package accordion

// from file "react-core/src/components/Accordion/AccordionContent.tsx"

type ContentProps struct {
	// Children - Content rendered inside the Accordion. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Accordion content. Optional.
	ClassName string
	// Id - Identify the AccordionContent item. Optional.
	Id string
	// IsHidden - Flag to show if the expanded content of the Accordion item is visible. Optional.
	IsHidden bool
	// IsFixed - Flag to indicate Accordion content is fixed. Optional.
	IsFixed bool
	// AriaLabel - Adds accessible text to the Accordion content. Optional.
	AriaLabel string
	// Component - Component to use as content container. Optional.
	Component any // React.ElementType 
	// IsCustomContent - Flag indicating content is custom. Expanded content Body wrapper will be removed from
	// children.  This allows multiple bodies to be rendered as content. Optional.
	IsCustomContent any // React.ReactNode 
}
