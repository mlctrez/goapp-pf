package text

// from file "react-core/src/components/Text/Text.tsx"

type Props struct {
	// Component - The text component. Optional.
	Component any /* "h1" | "h2" | "h3" | "h4" | "h5" | "h6" | "p" | "a" | "small" | "blockquote" | "pre" */
	// Children - Content rendered within the Text. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Text. Optional.
	ClassName string
	// IsVisitedLink - Flag to indicate the link has visited styles applied if the browser determines the link
	// has been visited. Optional.
	IsVisitedLink bool
}
