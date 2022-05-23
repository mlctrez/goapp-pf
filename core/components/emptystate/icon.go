package emptystate

// from file "react-core/src/components/EmptyState/EmptyStateIcon.tsx"

type IconProps struct {
	// Color - Changes the color of the icon. Optional.
	Color string
}

type IconProps struct {
	// ClassName - Additional classes added to the EmptyState. Optional.
	ClassName string
	// Icon - Icon component to be rendered inside the EmptyState on icon variant Usually a CheckCircleIcon,
	// ExclamationCircleIcon, LockIcon, PlusCircleIcon, RocketIcon SearchIcon, or WrenchIcon. Optional.
	Icon any // React.ComponentType 
	// Component - Component to be rendered inside the EmptyState on container variant. Optional.
	Component any // React.ComponentType 
	// Variant - Adds empty state icon variant styles. Optional.
	Variant any /* "icon" | "container" */
}
