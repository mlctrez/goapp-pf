package aboutmodal

// from file "react-core/src/components/AboutModal/AboutModal.tsx"

type Props struct {
	// Children - Content rendered inside the about modal.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the about modal. Optional.
	ClassName string
	// IsOpen - Flag to show the about modal. Optional.
	IsOpen bool
	// OnClose - A callback for when the close button is clicked. Optional.
	OnClose func()
	// ProductName - Product name. Optional.
	ProductName string
	// Trademark - Trademark information. Optional.
	Trademark string
	// BrandImageSrc - The URL of the image for the brand.
	BrandImageSrc string
	// BrandImageAlt - The alternate text of the brand image.
	BrandImageAlt string
	// BackgroundImageSrc - The URL of the image for the background. Optional.
	BackgroundImageSrc string
	// NoAboutModalBoxContentContainer - Prevents the about modal from rendering content inside a container;
	// allows for more flexible layouts. Optional.
	NoAboutModalBoxContentContainer bool
	// AppendTo - The parent container to append the modal to. Defaults to document.body. Optional.
	AppendTo any /* any // HTMLElement  | func() any // HTMLElement  */
	// CloseButtonAriaLabel - Set aria label to the close button. Optional.
	CloseButtonAriaLabel string
	// DisableFocusTrap - Flag to disable focus trap. Optional.
	DisableFocusTrap bool
}

type ModalState struct {
	// Container - 
	Container any // HTMLElement 
}
