package menu

// from file "react-core/src/components/Menu/DrilldownMenu.tsx"

type DrilldownMenuProps struct {
	// Children - Items within drilldown sub-menu. Optional.
	Children any // React.ReactNode 
	// Id - ID of the drilldown sub-menu. Optional.
	Id string
	// IsMenuDrilledIn - Flag indicating whether the menu is drilled in. Optional.
	IsMenuDrilledIn bool
	// GetHeight - Optional callback to get the height of the sub menu. Optional.
	GetHeight func(height string)
}
