package tabs

// from file "react-core/src/components/Tabs/TabContent.tsx"

type TabContentProps struct {
	// Children - content rendered inside the tab content area if used outside Tabs component. Optional.
	Children any
	// Child - Child to show in the content area. Optional.
	Child any // React.ReactElement 
	// ClassName - class of tab content area if used outside Tabs component. Optional.
	ClassName string
	// ActiveKey - Identifies the active Tab. Optional.
	ActiveKey any /* int | string */
	// EventKey - uniquely identifies the controlling Tab if used outside Tabs component. Optional.
	EventKey any /* int | string */
	// InnerRef - Callback for the section ref. Optional.
	InnerRef any // React.Ref 
	// Id - id passed from parent to identify the content section.
	Id string
	// AriaLabel - title of controlling Tab if used outside Tabs component. Optional.
	AriaLabel string
}
