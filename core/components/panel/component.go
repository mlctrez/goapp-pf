package panel

// from file "react-core/src/components/Panel/Panel.tsx"

type Props struct {
	// Children - Content rendered inside the panel. Optional.
	Children any // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
	// Variant - Adds panel variant styles. Optional.
	Variant any /* "raised" | "bordered" */
	// IsScrollable - Flag to add scrollable styling to the panel. Optional.
	IsScrollable bool
	// InnerRef - Optional.
	InnerRef any // React.Ref 
}
