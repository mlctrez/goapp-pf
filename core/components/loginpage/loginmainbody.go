package loginpage

type LoginMainBodyProps struct {
	// Children - Content rendered inside the Login Main Body. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the Login Main Body. Optional.
	ClassName string
}

// contents of react-core/src/components/LoginPage/LoginMainBody.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Login/login';
// 
// export interface LoginMainBodyProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the Login Main Body */
//   children?: React.ReactNode;
//   /** Additional classes added to the Login Main Body */
//   className?: string;
// }
// 
// export const LoginMainBody: React.FunctionComponent<LoginMainBodyProps> = ({
//   children = null,
//   className = '',
//   ...props
// }: LoginMainBodyProps) => (
//   <div className={css(styles.loginMainBody, className)} {...props}>
//     {children}
//   </div>
// );
// LoginMainBody.displayName = 'LoginMainBody';
