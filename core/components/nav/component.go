package nav

// from file "react-core/src/components/Nav/Nav.tsx"

type Props struct {
	// Children - Anything that can be rendered inside of the nav. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the container. Optional.
	ClassName string
	// OnSelect - Callback for updating when item selection changes. Optional.
	OnSelect func(selectedItem map[string]string /* { groupId:{ int, string }, itemId:{ int, string }, to:string, "event":any // React.FormEvent  } */)
	// OnToggle - Callback for when a list is expanded or collapsed. Optional.
	OnToggle func(toggledItem map[string]string /* { groupId:{ int, string }, "isExpanded":bool, "event":any // React.MouseEvent  } */)
	// AriaLabel - Accessibility label. Optional.
	AriaLabel string
	// Theme - Indicates which theme color to use. Optional.
	Theme any /* "dark" | "light" */
	// Variant - For horizontal navs. Optional.
	Variant any /* "default" | "horizontal" | "tertiary" | "horizontal-subnav" */
}

type ContextProps struct {
	// OnSelect - Optional.
	OnSelect func(event any // React.FormEvent , groupId any /* int | string */, itemId any /* int | string */, to string, preventDefault bool, onClick func(e any // React.FormEvent , itemId any /* int | string */, groupId any /* int | string */, to string))
	// OnToggle - Optional.
	OnToggle func(event any // React.MouseEvent , groupId any /* int | string */, expanded bool)
	// UpdateIsScrollable - Optional.
	UpdateIsScrollable func(isScrollable bool)
	// IsHorizontal - Optional.
	IsHorizontal bool
	// FlyoutRef - Optional.
	FlyoutRef any // React.Ref 
	// SetFlyoutRef - Optional.
	SetFlyoutRef func(ref any // React.Ref )
}
