package loginpage

// from file "react-core/src/components/LoginPage/LoginForm.tsx"

type LoginFormProps struct {
	// NoAutoFocus - Flag to indicate if the first dropdown item should not gain initial focus. Optional.
	NoAutoFocus bool
	// ClassName - Additional classes added to the Login Main Body's Form. Optional.
	ClassName string
	// ShowHelperText - Flag indicating the Helper Text is visible *. Optional.
	ShowHelperText bool
	// HelperText - Content displayed in the Helper Text component *. Optional.
	HelperText any // React.ReactNode 
	// HelperTextIcon - Icon displayed to the left in the Helper Text. Optional.
	HelperTextIcon any // React.ReactNode 
	// UsernameLabel - Label for the Username Input Field. Optional.
	UsernameLabel string
	// UsernameValue - Value for the Username. Optional.
	UsernameValue string
	// OnChangeUsername - Function that handles the onChange event for the Username. Optional.
	OnChangeUsername func(value string, event any // React.FormEvent )
	// IsValidUsername - Flag indicating if the Username is valid. Optional.
	IsValidUsername bool
	// PasswordLabel - Label for the Password Input Field. Optional.
	PasswordLabel string
	// PasswordValue - Value for the Password. Optional.
	PasswordValue string
	// OnChangePassword - Function that handles the onChange event for the Password. Optional.
	OnChangePassword func(value string, event any // React.FormEvent )
	// IsValidPassword - Flag indicating if the Password is valid. Optional.
	IsValidPassword bool
	// IsShowPasswordEnabled - Flag indicating if the user can toggle hiding the password. Optional.
	IsShowPasswordEnabled bool
	// ShowPasswordAriaLabel - Accessible label for the show password button. Optional.
	ShowPasswordAriaLabel string
	// HidePasswordAriaLabel - Accessible label for the hide password button. Optional.
	HidePasswordAriaLabel string
	// LoginButtonLabel - Label for the Log in Button Input. Optional.
	LoginButtonLabel string
	// IsLoginButtonDisabled - Flag indicating if the Login Button is disabled. Optional.
	IsLoginButtonDisabled bool
	// OnLoginButtonClick - Function that is called when the Login button is clicked. Optional.
	OnLoginButtonClick func(event any // React.MouseEvent )
	// RememberMeLabel - Label for the Remember Me Checkbox that indicates the user should be kept logged in.
	// If the label is not provided, the checkbox will not show. Optional.
	RememberMeLabel string
	// IsRememberMeChecked - Flag indicating if the remember me Checkbox is checked. Optional.
	IsRememberMeChecked bool
	// OnChangeRememberMe - Function that handles the onChange event for the Remember Me Checkbox. Optional.
	OnChangeRememberMe func(checked bool, event any // React.FormEvent )
}
