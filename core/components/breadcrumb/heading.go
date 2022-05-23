package breadcrumb

// from file "react-core/src/components/Breadcrumb/BreadcrumbHeading.tsx"

type HeadingProps struct {
	// Children - Content rendered inside the breadcrumb title. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the breadcrumb item. Optional.
	ClassName string
	// To - HREF for breadcrumb link. Optional.
	To string
	// Target - Target for breadcrumb link. Optional.
	Target string
	// Component - Sets the base component to render. Defaults to <a>. Optional.
	Component any // React.ReactNode 
	// ShowDivider - Internal prop set by Breadcrumb on all but the first crumb. Optional.
	ShowDivider bool
}
