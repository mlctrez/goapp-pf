package breadcrumb

// from file "react-core/src/components/Breadcrumb/BreadcrumbItem.tsx"

type ItemRenderArgs struct {
	// ClassName - 
	ClassName string
	// AriaCurrent - 
	AriaCurrent any /* "page" | undefined */
}

type ItemProps struct {
	// Children - Content rendered inside the breadcrumb item. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the breadcrumb item. Optional.
	ClassName string
	// To - HREF for breadcrumb link. Optional.
	To string
	// IsActive - Flag indicating whether the item is active. Optional.
	IsActive bool
	// IsDropdown - Flag indicating whether the item contains a dropdown. Optional.
	IsDropdown bool
	// ShowDivider - Internal prop set by Breadcrumb on all but the first crumb. Optional.
	ShowDivider bool
	// Target - Target for breadcrumb link. Optional.
	Target string
	// Component - Sets the base component to render. Defaults to <a>. Optional.
	Component any // React.ElementType 
	// Render - A render function to render the component inside the breadcrumb item. Optional.
	Render func(props any // BreadcrumbItemRenderArgs ) any // React.ReactNode 
}
