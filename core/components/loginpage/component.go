package loginpage

// from file "react-core/src/components/LoginPage/LoginPage.tsx"

type Props struct {
	// Children - Anything that can be rendered inside of the LoginPage (e.g. <LoginPageForm>). Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the LoginPage. Optional.
	ClassName string
	// BrandImgSrc - Attribute that specifies the URL of the brand image for the LoginPage. Optional.
	BrandImgSrc string
	// BrandImgAlt - Attribute that specifies the alt text of the brand image for the LoginPage. Optional.
	BrandImgAlt string
	// BackgroundImgSrc - Attribute that specifies the URL of the background image for the LoginPage. Optional.
	BackgroundImgSrc any /* string | any // BackgroundImageSrcMap  */
	// BackgroundImgAlt - Attribute that specifies the alt text of the background image for the LoginPage. Optional.
	BackgroundImgAlt string
	// TextContent - Content rendered inside of the Text Component of the LoginPage. Optional.
	TextContent string
	// FooterListItems - Items rendered inside of the Footer List Component of the LoginPage. Optional.
	FooterListItems any // React.ReactNode 
	// FooterListVariants - Adds list variant styles for the Footer List component of the LoginPage. The only
	// current value is'inline'. Optional.
	FooterListVariants any // ListVariant.inline 
	// LoginTitle - Title for the Login Main Body Header of the LoginPage.
	LoginTitle string
	// LoginSubtitle - Subtitle for the Login Main Body Header of the LoginPage. Optional.
	LoginSubtitle string
	// SignUpForAccountMessage - Content rendered inside of Login Main Footer Band to display a sign up for account
	// message. Optional.
	SignUpForAccountMessage any // React.ReactNode 
	// ForgotCredentials - Content rendered inside of Login Main Footer Band to display a forgot credentials
	// link*. Optional.
	ForgotCredentials any // React.ReactNode 
	// SocialMediaLoginContent - Content rendered inside of Social Media Login footer section . Optional.
	SocialMediaLoginContent any // React.ReactNode 
}
