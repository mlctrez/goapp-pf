package loginpage

type LoginFooterItemProps struct {
	// Children - Content rendered inside the footer Link Item. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the Footer Link Item. Optional.
	ClassName string
	// Href - The URL of the Footer Link Item. Optional.
	Href string
	// Target - Specifies where to open the linked document. Optional.
	Target string
}

// contents of react-core/src/components/LoginPage/LoginFooterItem.tsx
// import * as React from 'react';
// 
// export interface LoginFooterItemProps extends React.HTMLProps<HTMLAnchorElement> {
//   /** Content rendered inside the footer Link Item */
//   children?: React.ReactNode;
//   /** Additional classes added to the Footer Link Item  */
//   className?: string;
//   /** The URL of the Footer Link Item */
//   href?: string;
//   /** Specifies where to open the linked document */
//   target?: string;
// }
// 
// export const LoginFooterItem: React.FunctionComponent<LoginFooterItemProps> = ({
//   children = null,
//   href = '#',
//   target = '_blank',
//   ...props
// }: LoginFooterItemProps) =>
//   React.isValidElement(children) ? (
//     children
//   ) : (
//     <a target={target} href={href} {...props}>
//       {children}
//     </a>
//   );
// LoginFooterItem.displayName = 'LoginFooterItem';
