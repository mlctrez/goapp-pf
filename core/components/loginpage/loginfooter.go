package loginpage

type LoginFooterProps struct {
	// Children - Content rendered inside the footer of the login layout. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional props are spread to the container <footer>. Optional.
	ClassName string
}

// contents of react-core/src/components/LoginPage/LoginFooter.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Login/login';
// import { css } from '@patternfly/react-styles';
// 
// export interface LoginFooterProps extends React.HTMLProps<HTMLElement> {
//   /** Content rendered inside the footer of the login layout */
//   children?: React.ReactNode;
//   /** Additional props are spread to the container <footer> */
//   className?: string;
// }
// 
// export const LoginFooter: React.FunctionComponent<LoginFooterProps> = ({
//   className = '',
//   children = null,
//   ...props
// }: LoginFooterProps) => (
//   <footer className={css(styles.loginFooter, className)} {...props}>
//     {children}
//   </footer>
// );
// LoginFooter.displayName = 'LoginFooter';
