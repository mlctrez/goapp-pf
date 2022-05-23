package modal

// from file "react-core/src/components/Modal/ModalBoxTitle.tsx"

type BoxTitleProps struct {
	// Title - Content rendered inside the modal box header title.
	Title any // React.ReactNode 
	// TitleIconVariant - Optional alert icon (or other) to show before the title of the Modal Header When the
	// predefined alert types are used the default styling will be automatically applied. Optional.
	TitleIconVariant any /* "success" | "danger" | "warning" | "info" | "default" | any // React.ComponentType  */
	// TitleLabel - Optional title label text for screen readers. Optional.
	TitleLabel string
	// ClassName - Additional classes added to the modal box header title. Optional.
	ClassName string
	// Id - id of the modal box header title.
	Id string
}
