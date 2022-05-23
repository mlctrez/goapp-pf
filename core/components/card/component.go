package card

// from file "react-core/src/components/Card/Card.tsx"

type Props struct {
	// Children - Content rendered inside the Card. Optional.
	Children any // React.ReactNode 
	// Id - ID of the Card. Also passed back in the CardHeader onExpand callback. Optional.
	Id string
	// ClassName - Additional classes added to the Card. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to article. Optional.
	Component any /* keyof any // JSX.IntrinsicElements  */
	// IsHoverable - Optional.
	IsHoverable bool
	// IsCompact - Modifies the card to include compact styling. Should not be used with isLarge. Optional.
	IsCompact bool
	// IsSelectable - Modifies the card to include selectable styling. Optional.
	IsSelectable bool
	// IsSelectableRaised - Optional.
	IsSelectableRaised bool
	// IsSelected - Modifies the card to include selected styling. Optional.
	IsSelected bool
	// IsDisabledRaised - Optional.
	IsDisabledRaised bool
	// IsFlat - Modifies the card to include flat styling. Optional.
	IsFlat bool
	// IsRounded - Modifies the card to include rounded styling. Optional.
	IsRounded bool
	// IsLarge - Modifies the card to be large. Should not be used with isCompact. Optional.
	IsLarge bool
	// IsFullHeight - Cause component to consume the available height of its container. Optional.
	IsFullHeight bool
	// IsPlain - Modifies the card to include plain styling; this removes border and background. Optional.
	IsPlain bool
	// IsExpanded - Flag indicating if a card is expanded. Modifies the card to be expandable. Optional.
	IsExpanded bool
	// HasSelectableInput - Flag indicating that the card should render a hidden input to make it selectable.
	// Optional.
	HasSelectableInput bool
	// SelectableInputAriaLabel - Aria label to apply to the selectable input if one is rendered. Optional.
	SelectableInputAriaLabel string
	// OnSelectableInputChange - Callback that executes when the selectable input is changed. Optional.
	OnSelectableInputChange func(labelledBy string, event any // React.FormEvent )
}

type ContextProps struct {
	// CardId - 
	CardId string
	// RegisterTitleId - 
	RegisterTitleId func(id string)
	// IsExpanded - 
	IsExpanded bool
}

type AriaProps struct {
	// AriaLabel - Optional.
	AriaLabel string
	// AriaLabelledby - Optional.
	AriaLabelledby string
}
