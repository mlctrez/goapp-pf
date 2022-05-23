package applicationlauncher

// from file "react-core/src/components/ApplicationLauncher/ApplicationLauncherItem.tsx"

type ItemProps struct {
	// Icon - Icon rendered before the text. Optional.
	Icon any // React.ReactNode 
	// IsExternal - If clicking on the item should open the page in a separate window. Optional.
	IsExternal bool
	// Tooltip - Tooltip to display when hovered over the item. Optional.
	Tooltip any // React.ReactNode 
	// TooltipProps - Additional tooltip props forwarded to the Tooltip component. Optional.
	TooltipProps any
	// Component - A ReactElement to render, or a string to use as the component tag. Example: component={<Link
	// to="/components/alert/">Alert</Link>} Example: component="button". Optional.
	Component any // React.ReactNode 
	// IsFavorite - Flag indicating if the item is favorited. Optional.
	IsFavorite bool
	// AriaIsFavoriteLabel - Aria label text for favoritable button when favorited. Optional.
	AriaIsFavoriteLabel string
	// AriaIsNotFavoriteLabel - Aria label text for favoritable button when not favorited. Optional.
	AriaIsNotFavoriteLabel string
	// Id - ID of the item. Required for tracking favorites. Optional.
	Id string
	// CustomChild - Optional.
	CustomChild any // React.ReactNode 
	// EnterTriggersArrowDown - Flag indicating if hitting enter triggers an arrow down key press. Automatically
	// passed to favorites list items. Optional.
	EnterTriggersArrowDown bool
}
