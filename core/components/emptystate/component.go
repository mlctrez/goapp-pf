package emptystate

// from file "react-core/src/components/EmptyState/EmptyState.tsx"

type Props struct {
	// ClassName - Additional classes added to the EmptyState. Optional.
	ClassName string
	// Children - Content rendered inside the EmptyState.
	Children any // React.ReactNode 
	// Variant - Modifies EmptyState max-width. Optional.
	Variant any /* "xs" | "small" | "large" | "xl" | "full" */
	// IsFullHeight - Cause component to consume the available height of its container. Optional.
	IsFullHeight bool
}
