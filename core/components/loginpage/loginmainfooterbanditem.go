package loginpage

type LoginMainFooterBandItemProps struct {
	// Children - Content rendered inside the footer Link Item. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the Footer Link Item. Optional.
	ClassName string
}

// contents of react-core/src/components/LoginPage/LoginMainFooterBandItem.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Login/login';
// import { css } from '@patternfly/react-styles';
// 
// export interface LoginMainFooterBandItemProps extends React.HTMLProps<HTMLParagraphElement> {
//   /** Content rendered inside the footer Link Item */
//   children?: React.ReactNode;
//   /** Additional classes added to the Footer Link Item  */
//   className?: string;
// }
// 
// export const LoginMainFooterBandItem: React.FunctionComponent<LoginMainFooterBandItemProps> = ({
//   children = null,
//   className = '',
//   ...props
// }: LoginMainFooterBandItemProps) => (
//   <p className={css(`${styles.loginMainFooterBand}-item`, className)} {...props}>
//     {children}
//   </p>
// );
// LoginMainFooterBandItem.displayName = 'LoginMainFooterBandItem';
