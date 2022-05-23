package tile

// from file "react-core/src/components/Tile/Tile.tsx"

type Props struct {
	// Children - Content rendered inside the banner. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the banner. Optional.
	ClassName string
	// Title - Title of the tile.
	Title string
	// Icon - Icon in the tile title. Optional.
	Icon any // React.ReactNode 
	// IsSelected - Flag indicating if the tile is selected. Optional.
	IsSelected bool
	// IsDisabled - Flag indicating if the tile is disabled. Optional.
	IsDisabled bool
	// IsStacked - Flag indicating if the tile header is stacked. Optional.
	IsStacked bool
	// IsDisplayLarge - Flag indicating if the stacked tile icon is large. Optional.
	IsDisplayLarge bool
}
