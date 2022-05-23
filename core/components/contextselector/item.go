package contextselector

// from file "react-core/src/components/ContextSelector/ContextSelectorItem.tsx"

type ItemProps struct {
	// Children - Anything which can be rendered as Context Selector item. Optional.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of the Context Selector item. Optional.
	ClassName string
	// IsDisabled - Render Context  Selector item as disabled. Optional.
	IsDisabled bool
	// OnClick - Callback for click event.
	OnClick func(event any // React.MouseEvent )
	// Index - 
	Index int
	// SendRef - Internal callback for ref tracking.
	SendRef func(index int, current any)
	// Href - Link href, indicates item should render as anchor tag. Optional.
	Href string
}
