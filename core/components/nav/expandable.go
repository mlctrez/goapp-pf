package nav

// from file "react-core/src/components/Nav/NavExpandable.tsx"

type ExpandableProps struct {
	// Title - Title shown for the expandable list.
	Title string
	// SrText - If defined, screen readers will read this text instead of the list title. Optional.
	SrText string
	// IsExpanded - Boolean to programatically expand or collapse section. Optional.
	IsExpanded bool
	// Children - Anything that can be rendered inside of the expandable list. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the container. Optional.
	ClassName string
	// GroupId - Group identifier, will be returned with the onToggle and onSelect callback passed to the Nav
	// component. Optional.
	GroupId any /* string | int */
	// IsActive - If true makes the expandable list title active. Optional.
	IsActive bool
	// Id - Identifier to use for the section aria label. Optional.
	Id string
	// OnExpand - allow consumer to optionally override this callback and manage expand state externally. if
	// passed will not call Nav's onToggle. Optional.
	OnExpand func(e any // React.MouseEvent , val bool)
	// ButtonProps - Additional props added to the NavExpandable <button>. Optional.
	ButtonProps any
}

type ExpandableState struct {
	// ExpandedState - 
	ExpandedState bool
	// OuiaStateId - 
	OuiaStateId string
}
