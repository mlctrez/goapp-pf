package form

// from file "react-core/src/components/Form/FormFieldGroupHeader.tsx"

type FiledGroupHeaderTitleTextObject struct {
	// Text - Title text.
	Text any // React.ReactNode 
	// Id - The applied to the title div for accessibility.
	Id string
}

type FieldGroupHeaderTitleTextObject struct {
}

type FieldGroupHeaderProps struct {
	// ClassName - Additional classes added to the section. Optional.
	ClassName string
	// TitleText - Field group header title text. Optional.
	TitleText any // FormFieldGroupHeaderTitleTextObject 
	// TitleDescription - Field group header title description. Optional.
	TitleDescription any // React.ReactNode 
	// Actions - Field group header actions. Optional.
	Actions any // React.ReactNode 
}
