package aboutmodal

// from file "react-core/src/components/AboutModal/AboutModalBoxContent.tsx"

type BoxContentProps struct {
	// Children - content rendered inside the AboutModalBoxContent.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the AboutModalBoxContent. Optional.
	ClassName string
	// Id - id to use for About Modal Box aria described by.
	Id string
	// Trademark - The Trademark info for the product.
	Trademark string
	// NoAboutModalBoxContentContainer - Prevents the about modal from rendering content inside a container;
	// allows for more flexible layouts. Optional.
	NoAboutModalBoxContentContainer bool
}
