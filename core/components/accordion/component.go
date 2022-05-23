package accordion

// from file "react-core/src/components/Accordion/Accordion.tsx"

type Props struct {
	// Children - Content rendered inside the Accordion. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Accordion. Optional.
	ClassName string
	// AriaLabel - Adds accessible text to the Accordion. Optional.
	AriaLabel string
	// HeadingLevel - Heading level to use. Optional.
	HeadingLevel any /* "h1" | "h2" | "h3" | "h4" | "h5" | "h6" */
	// AsDefinitionList - Flag to indicate whether use definition list or div. Optional.
	AsDefinitionList bool
	// IsBordered - Flag to indicate the accordion had a border. Optional.
	IsBordered bool
	// DisplaySize - Display size variant. Optional.
	DisplaySize any /* "default" | "large" */
}
