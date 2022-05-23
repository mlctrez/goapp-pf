package card

// from file "react-core/src/components/Card/CardHeader.tsx"

type HeaderProps struct {
	// Children - Content rendered inside the CardHeader. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the CardHeader. Optional.
	ClassName string
	// Id - ID of the card header. Optional.
	Id string
	// OnExpand - Callback expandable card. Optional.
	OnExpand func(event any // React.MouseEvent , id string)
	// ToggleButtonProps - Additional props for expandable toggle button. Optional.
	ToggleButtonProps any
	// IsToggleRightAligned - Whether to right-align expandable toggle button. Optional.
	IsToggleRightAligned bool
}
