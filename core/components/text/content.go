package text

// from file "react-core/src/components/Text/TextContent.tsx"

type ContentProps struct {
	// Children - Content rendered within the TextContent. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the TextContent. Optional.
	ClassName string
	// IsVisited - Flag to indicate the all links in a the content block have visited styles applied if the browser
	// determines the link has been visited. Optional.
	IsVisited bool
}
