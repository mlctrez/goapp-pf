package menu

type FooterProps struct {
	// Children - Content rendered inside the footer. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the footer. Optional.
	ClassName string
}

// contents of react-core/src/components/Menu/MenuFooter.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Menu/menu';
// import { css } from '@patternfly/react-styles';
// 
// export interface MenuFooterProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the footer */
//   children?: React.ReactNode;
//   /** Additional classes added to the footer */
//   className?: string;
// }
// 
// export const MenuFooter: React.FunctionComponent<MenuFooterProps> = ({
//   children,
//   className = '',
//   ...props
// }: MenuFooterProps) => (
//   <div {...props} className={css(styles.menuFooter, className)}>
//     {children}
//   </div>
// );
// 
// MenuFooter.displayName = 'MenuFooter';
