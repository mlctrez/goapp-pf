package loginpage

type LoginMainHeaderProps struct {
	// Children - Content rendered inside the Login Main Header. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the Login Main Header. Optional.
	ClassName string
	// Title - Title for the Login Main Header. Optional.
	Title string
	// Subtitle - Subtitle that contains the Text, URL, and URL Text for the Login Main Header. Optional.
	Subtitle string
}

// contents of react-core/src/components/LoginPage/LoginMainHeader.tsx
// import * as React from 'react';
// import { Title, TitleSizes } from '../Title';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Login/login';
// 
// export interface LoginMainHeaderProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the Login Main Header */
//   children?: React.ReactNode;
//   /** Additional classes added to the Login Main Header */
//   className?: string;
//   /** Title for the Login Main Header */
//   title?: string;
//   /** Subtitle that contains the Text, URL, and URL Text for the Login Main Header */
//   subtitle?: string;
// }
// 
// export const LoginMainHeader: React.FunctionComponent<LoginMainHeaderProps> = ({
//   children = null,
//   className = '',
//   title = '',
//   subtitle = '',
//   ...props
// }: LoginMainHeaderProps) => (
//   <header className={css(styles.loginMainHeader, className)} {...props}>
//     {title && (
//       <Title headingLevel="h2" size={TitleSizes['3xl']}>
//         {title}
//       </Title>
//     )}
//     {subtitle && <p className={css(styles.loginMainHeaderDesc)}>{subtitle}</p>}
//     {children}
//   </header>
// );
// LoginMainHeader.displayName = 'LoginMainHeader';
