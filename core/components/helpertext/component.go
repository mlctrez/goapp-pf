package helpertext

// from file "react-core/src/components/HelperText/HelperText.tsx"

type Props struct {
	// Children - Content rendered inside the helper text container. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes applied to the helper text container. Optional.
	ClassName string
	// Component - Component type of the helper text container. Optional.
	Component any /* "div" | "ul" */
	// Id - ID for the helper text container. The value of this prop can be passed into a form component's aria-describedby
	// prop when you intend for all helper text items to be announced to assistive technologies. Optional.
	Id string
	// IsLiveRegion - Flag for indicating whether the helper text container is a live region. Use this prop when
	// you expect or intend for any helper text items within the container to be dynamically updated. Optional.
	IsLiveRegion bool
}
