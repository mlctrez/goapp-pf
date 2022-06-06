package loginpage

type LoginMainFooterProps struct {
	// ClassName - Additional classes added to the Login Main Footer. Optional.
	ClassName string
	// Children - Content rendered inside the Login Main Footer. Optional.
	Children any //  // React.ReactNode 
	// SocialMediaLoginContent - Content rendered inside the Login Main Footer as Social Media Links*. Optional.
	SocialMediaLoginContent any //  // React.ReactNode 
	// SignUpForAccountMessage - Content rendered inside of Login Main Footer Band to display a sign up for account
	// message. Optional.
	SignUpForAccountMessage any //  // React.ReactNode 
	// ForgotCredentials - Content rendered inside of Login Main Footer Band do display a forgot credentials
	// link*. Optional.
	ForgotCredentials any //  // React.ReactNode 
}

// contents of react-core/src/components/LoginPage/LoginMainFooter.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Login/login';
// 
// export interface LoginMainFooterProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes added to the Login Main Footer */
//   className?: string;
//   /** Content rendered inside the Login Main Footer */
//   children?: React.ReactNode;
//   /** Content rendered inside the Login Main Footer as Social Media Links* */
//   socialMediaLoginContent?: React.ReactNode;
//   /** Content rendered inside of Login Main Footer Band to display a sign up for account message */
//   signUpForAccountMessage?: React.ReactNode;
//   /** Content rendered inside of Login Main Footer Band do display a forgot credentials link* */
//   forgotCredentials?: React.ReactNode;
// }
// 
// export const LoginMainFooter: React.FunctionComponent<LoginMainFooterProps> = ({
//   children = null,
//   socialMediaLoginContent = null,
//   signUpForAccountMessage = null,
//   forgotCredentials = null,
//   className = '',
//   ...props
// }: LoginMainFooterProps) => (
//   <div className={css(styles.loginMainFooter, className)} {...props}>
//     {children}
//     {socialMediaLoginContent && <ul className={css(styles.loginMainFooterLinks)}>{socialMediaLoginContent}</ul>}
//     {(signUpForAccountMessage || forgotCredentials) && (
//       <div className={css(styles.loginMainFooterBand)}>
//         {signUpForAccountMessage}
//         {forgotCredentials}
//       </div>
//     )}
//   </div>
// );
// LoginMainFooter.displayName = 'LoginMainFooter';
