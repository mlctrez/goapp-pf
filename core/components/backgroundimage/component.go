package backgroundimage

// from file "react-core/src/components/BackgroundImage/BackgroundImage.tsx"

type SrcMap struct {
	// Xs - 
	Xs string
	// Xs2x - 
	Xs2x string
	// Sm - 
	Sm string
	// Sm2x - 
	Sm2x string
	// Lg - 
	Lg string
}

type Props struct {
	// ClassName - Additional classes added to the background. Optional.
	ClassName string
	// Filter - Override svg filter to use. Optional.
	Filter any // React.ReactElement 
	// Src - Override image styles using a string or BackgroundImageSrc.
	Src any /* string | any // BackgroundImageSrcMap  */
}
