package alert

// from file "react-core/src/components/Alert/AlertIcon.tsx"

type IconProps struct {
	// Variant - variant.
	Variant any /* "success" | "danger" | "warning" | "info" | "default" */
	// ClassName - className. Optional.
	ClassName string
	// CustomIcon - A custom icon. If not set the icon is set according to the variant. Optional.
	CustomIcon any // React.ReactNode 
}
