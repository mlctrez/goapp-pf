package datalist

// from file "react-core/src/components/DataList/DataListAction.tsx"

type ActionProps struct {
	// Children - Content rendered as DataList Action  (e.g <Button> or <Dropdown>).
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the DataList Action. Optional.
	ClassName string
	// Id - Identify the DataList toggle number.
	Id string
	// AriaLabelledby - Adds accessible text to the DataList Action.
	AriaLabelledby string
	// AriaLabel - Adds accessible text to the DataList Action.
	AriaLabel string
	// Visibility - What breakpoints to hide/show the data list action. Optional.
	Visibility map[string]string /* { default:{ hidden, visible }, sm:{ hidden, visible }, md:{ hidden, visible }, lg:{ hidden, visible }, xl:{ hidden, visible }, 2xl:{ hidden, visible } } */
	// IsPlainButtonAction - Flag to indicate that the action is a plain button (e.g. kebab dropdown toggle)
	// so that styling is applied to align the button. Optional.
	IsPlainButtonAction bool
}
