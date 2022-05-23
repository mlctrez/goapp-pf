package aboutmodal

// from file "react-core/src/components/AboutModal/AboutModalContainer.tsx"

type ContainerProps struct {
	// Children - content rendered inside the About Modal Box Content.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the About Modal Box. Optional.
	ClassName string
	// IsOpen - Flag to show the About Modal. Optional.
	IsOpen bool
	// OnClose - A callback for when the close button is clicked. Optional.
	OnClose func()
	// ProductName - Product Name. Optional.
	ProductName string
	// Trademark - Trademark information. Optional.
	Trademark string
	// BrandImageSrc - the URL of the image for the Brand.
	BrandImageSrc string
	// BrandImageAlt - the alternate text of the Brand image.
	BrandImageAlt string
	// BackgroundImageSrc - the URL of the image for the background. Optional.
	BackgroundImageSrc string
	// AboutModalBoxHeaderId - id to use for About Modal Box aria labeled by.
	AboutModalBoxHeaderId string
	// AboutModalBoxContentId - id to use for About Modal Box aria described by.
	AboutModalBoxContentId string
	// CloseButtonAriaLabel - Set close button aria label. Optional.
	CloseButtonAriaLabel string
	// DisableFocusTrap - Flag to disable focus trap. Optional.
	DisableFocusTrap bool
}
