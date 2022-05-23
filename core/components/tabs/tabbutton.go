package tabs

// from file "react-core/src/components/Tabs/TabButton.tsx"

type TabButtonProps struct {
	// Children - content rendered inside the Tab content area. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Tab. Optional.
	ClassName string
	// Href - URL associated with the Tab. A Tab with an href will render as an <a> instead of a <button>. A
	// Tab inside a <Tabs component="nav"> should have an href. Optional.
	Href string
	// TabContentRef - child reference for case in which a TabContent section is defined outside of a Tabs component.
	// Optional.
	TabContentRef any // React.Ref 
	// ParentInnerRef - Parents' innerRef passed down for properly displaying Tooltips. Optional.
	ParentInnerRef any // React.Ref 
}
