package alert

// from file "react-core/src/components/Alert/Alert.tsx"

type Props struct {
	// Variant - Adds alert variant styles. Optional.
	Variant any /* "success" | "danger" | "warning" | "info" | "default" */
	// IsInline - Flag to indicate if the alert is inline. Optional.
	IsInline bool
	// IsPlain - Flag to indicate if the alert is plain. Optional.
	IsPlain bool
	// Title - Title of the alert.
	Title any // React.ReactNode 
	// TitleHeadingLevel - Sets the heading level to use for the alert title. Default is h4. Optional.
	TitleHeadingLevel any /* "h1" | "h2" | "h3" | "h4" | "h5" | "h6" */
	// ActionClose - Close button; use the alertActionCloseButton component. Optional.
	ActionClose any // React.ReactNode 
	// ActionLinks - Action links; use a single alertActionLink component or multiple wrapped in an array or
	// React.Fragment. Optional.
	ActionLinks any // React.ReactNode 
	// Children - Content rendered inside the alert. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the alert. Optional.
	ClassName string
	// AriaLabel - Adds accessible text to the alert. Optional.
	AriaLabel string
	// VariantLabel - Variant label text for screen readers. Optional.
	VariantLabel string
	// IsLiveRegion - Flag to indicate if the alert is in a live region. Optional.
	IsLiveRegion bool
	// Timeout - If set to true, the timeout is 8000 milliseconds. If a number is provided, alert will be dismissed
	// after that amount of time in milliseconds. Optional.
	Timeout any /* int | bool */
	// TimeoutAnimation - If the user hovers over the alert and `timeout` expires, this is how long to wait before
	// finally dismissing the alert. Optional.
	TimeoutAnimation int
	// OnTimeout - Function to be executed on alert timeout. Relevant when the timeout prop is set. Optional.
	OnTimeout func()
	// TruncateTitle - Truncate title to number of lines. Optional.
	TruncateTitle int
	// TooltipPosition - Position of the tooltip which is displayed if text is truncated. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// CustomIcon - Set a custom icon to the alert. If not set the icon is set according to the variant. Optional.
	CustomIcon any // React.ReactNode 
	// IsExpandable - Flag indicating that the alert is expandable. Optional.
	IsExpandable bool
	// ToggleAriaLabel - Adds accessible text to the alert Toggle. Optional.
	ToggleAriaLabel string
}
