package expandablesection

// from file "react-core/src/components/ExpandableSection/ExpandableSection.tsx"

type Props struct {
	// Children - Content rendered inside the Expandable Component. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Expandable Component. Optional.
	ClassName string
	// IsExpanded - Flag to indicate if the content is expanded. Optional.
	IsExpanded bool
	// ToggleText - Text that appears in the attached toggle. Optional.
	ToggleText string
	// ToggleTextExpanded - Text that appears in the attached toggle when expanded (will override toggleText
	// if both are specified; used for uncontrolled expandable with dynamic toggle text). Optional.
	ToggleTextExpanded string
	// ToggleTextCollapsed - Text that appears in the attached toggle when collapsed (will override toggleText
	// if both are specified; used for uncontrolled expandable with dynamic toggle text). Optional.
	ToggleTextCollapsed string
	// ToggleContent - React node that appears in the attached toggle in place of toggle text. Optional.
	ToggleContent any // React.ReactNode 
	// OnToggle - Callback function to toggle the expandable content. Detached expandable sections should use
	// the onToggle property of ExpandableSectionToggle. Optional.
	OnToggle func(isExpanded bool)
	// IsActive - Forces active state. Optional.
	IsActive bool
	// IsDetached - Indicates the expandable section has a detached toggle. Optional.
	IsDetached bool
	// ContentId - ID of the content of the expandable section. Optional.
	ContentId string
	// DisplaySize - Display size variant. Set to large for disclosure styling. Optional.
	DisplaySize any /* "default" | "large" */
	// IsWidthLimited - Flag to indicate the width of the component is limited. Set to true for disclosure styling.
	// Optional.
	IsWidthLimited bool
	// IsIndented - Flag to indicate if the content is indented. Optional.
	IsIndented bool
}

type State struct {
	// IsExpanded - 
	IsExpanded bool
}
