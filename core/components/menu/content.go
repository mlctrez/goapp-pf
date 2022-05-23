package menu

// from file "react-core/src/components/Menu/MenuContent.tsx"

type ContentProps struct {
	// Children - Items within group. Optional.
	Children any // React.ReactNode 
	// InnerRef - Forwarded ref. Optional.
	InnerRef any // React.Ref 
	// MenuHeight - Height of the menu content. Optional.
	MenuHeight string
	// MaxMenuHeight - Maximum height of menu content. Optional.
	MaxMenuHeight string
	// GetHeight - Callback to return the height of the menu content. Optional.
	GetHeight func(height string)
}
