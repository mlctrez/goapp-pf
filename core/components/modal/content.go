package modal

// from file "react-core/src/components/Modal/ModalContent.tsx"

type ContentProps struct {
	// Children - Content rendered inside the Modal.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the button. Optional.
	ClassName string
	// Variant - Variant of the modal. Optional.
	Variant any /* "small" | "medium" | "large" | "default" */
	// Position - Alternate position of the modal. Optional.
	Position "top"
	// PositionOffset - Offset from alternate position. Can be any valid CSS length/percentage. Optional.
	PositionOffset string
	// IsOpen - Flag to show the modal. Optional.
	IsOpen bool
	// Header - Complex header (more than just text), supersedes title for header content. Optional.
	Header any // React.ReactNode 
	// Help - Optional help section for the Modal Header. Optional.
	Help any // React.ReactNode 
	// Description - Description of the modal. Optional.
	Description any // React.ReactNode 
	// Title - Simple text content of the Modal Header, also used for aria-label on the body. Optional.
	Title string
	// TitleIconVariant - Optional alert icon (or other) to show before the title of the Modal Header When the
	// predefined alert types are used the default styling will be automatically applied. Optional.
	TitleIconVariant any /* "success" | "danger" | "warning" | "info" | "default" | any // React.ComponentType  */
	// TitleLabel - Optional title label text for screen readers. Optional.
	TitleLabel string
	// AriaLabelledby - Id of Modal Box label. Optional.
	AriaLabelledby any /* string | "" */
	// AriaLabel - Accessible descriptor of modal. Optional.
	AriaLabel string
	// AriaDescribedby - Id of Modal Box description. Optional.
	AriaDescribedby string
	// BodyAriaLabel - Accessible label applied to the modal box body. This should be used to communicate important
	// information about the modal box body div if needed, such as that it is scrollable. Optional.
	BodyAriaLabel string
	// BodyAriaRole - Accessible role applied to the modal box body. This will default to region if a body aria
	// label is applied. Set to a more appropriate role as applicable based on the modal content and context.
	// Optional.
	BodyAriaRole string
	// ShowClose - Flag to show the close button in the header area of the modal. Optional.
	ShowClose bool
	// Width - Default width of the content. Optional.
	Width any /* int | string */
	// Footer - Custom footer. Optional.
	Footer any // React.ReactNode 
	// Actions - Action buttons to add to the standard Modal Footer, ignored if `footer` is given. Optional.
	Actions any
	// OnClose - A callback for when the close button is clicked. Optional.
	OnClose func()
	// BoxId - Id of the ModalBox container.
	BoxId string
	// LabelId - Id of the ModalBox title.
	LabelId string
	// DescriptorId - Id of the ModalBoxBody.
	DescriptorId string
	// DisableFocusTrap - Flag to disable focus trap. Optional.
	DisableFocusTrap bool
	// HasNoBodyWrapper - Flag indicating if modal content should be placed in a modal box body wrapper. Optional.
	HasNoBodyWrapper bool
}
