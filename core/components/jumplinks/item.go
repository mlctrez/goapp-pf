package jumplinks

// from file "react-core/src/components/JumpLinks/JumpLinksItem.tsx"

type ItemProps struct {
	// IsActive - Whether this item is active. Parent JumpLinks component sets this when passed a `scrollableSelector`.
	// Optional.
	IsActive bool
	// Href - Href for this link. Optional.
	Href string
	// Node - Selector or HTMLElement to spy on. Optional.
	Node any /* string | any // HTMLElement  */
	// Children - Text to be rendered inside span. Optional.
	Children any // React.ReactNode 
	// OnClick - Click handler for anchor tag. Parent JumpLinks components tap into this. Optional.
	OnClick func(ev any // React.MouseEvent )
	// ClassName - Class to add to li. Optional.
	ClassName string
}
