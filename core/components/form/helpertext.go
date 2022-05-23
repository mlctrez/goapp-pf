package form

// from file "react-core/src/components/Form/FormHelperText.tsx"

type HelperTextProps struct {
	// Children - Content rendered inside the Helper Text Item. Optional.
	Children any // React.ReactNode 
	// IsError - Adds error styling to the Helper Text  *. Optional.
	IsError bool
	// IsHidden - Hides the helper text *. Optional.
	IsHidden bool
	// ClassName - Additional classes added to the Helper Text Item. Optional.
	ClassName string
	// Icon - Icon displayed to the left of the helper text. Optional.
	Icon any // React.ReactNode 
	// Component - Component type of the form helper text. Optional.
	Component any /* "p" | "div" */
}
