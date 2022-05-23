package button

// from file "react-core/src/components/Button/Button.tsx"

type Props struct {
	// Children - Content rendered inside the button. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the button. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to button. Optional.
	Component any /* any // React.ElementType  | any // React.ComponentType  */
	// IsActive - Adds active styling to button. Optional.
	IsActive bool
	// IsBlock - Adds block styling to button. Optional.
	IsBlock bool
	// IsDisabled - Adds disabled styling and disables the button using the disabled html attribute. Optional.
	IsDisabled bool
	// IsAriaDisabled - Adds disabled styling and communicates that the button is disabled using the aria-disabled
	// html attribute. Optional.
	IsAriaDisabled bool
	// IsLoading - Adds progress styling to button. Optional.
	IsLoading bool
	// SpinnerAriaValueText - Text describing that current loading status or progress. Optional.
	SpinnerAriaValueText string
	// SpinnerAriaLabel - Accessible label for the spinner to describe what is loading. Optional.
	SpinnerAriaLabel string
	// SpinnerAriaLabelledBy - Id of element which describes what is being loaded. Optional.
	SpinnerAriaLabelledBy string
	// InoperableEvents - Events to prevent when the button is in an aria-disabled state. Optional.
	InoperableEvents []string
	// IsInline - Adds inline styling to a link button. Optional.
	IsInline bool
	// Type - Sets button type. Optional.
	Type any /* "button" | "submit" | "reset" */
	// Variant - Adds button variant styles. Optional.
	Variant any /* "primary" | "secondary" | "tertiary" | "danger" | "warning" | "link" | "plain" | "control" */
	// IconPosition - Sets position of the link icon. Optional.
	IconPosition any /* "left" | "right" */
	// AriaLabel - Adds accessible text to the button. Optional.
	AriaLabel string
	// Icon - Icon for the button. Usable by all variants except for plain. Optional.
	Icon any /* any // React.ReactNode  | "" */
	// TabIndex - Sets the button tabindex. Optional.
	TabIndex int
	// IsSmall - Adds small styling to the button. Optional.
	IsSmall bool
	// IsLarge - Adds large styling to the button. Optional.
	IsLarge bool
	// IsDanger - Adds danger styling to secondary or link button variants. Optional.
	IsDanger bool
	// InnerRef - Forwarded ref. Optional.
	InnerRef any // React.Ref 
}
