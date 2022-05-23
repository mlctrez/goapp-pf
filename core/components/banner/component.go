package banner

// from file "react-core/src/components/Banner/Banner.tsx"

type Props struct {
	// Children - Content rendered inside the banner. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the banner. Optional.
	ClassName string
	// Variant - Variant styles for the banner. Optional.
	Variant any /* "default" | "info" | "danger" | "success" | "warning" */
	// IsSticky - If set to true, the banner sticks to the top of its container. Optional.
	IsSticky bool
}
