package menu

type BreadcrumbProps struct {
	// Children - Items within breadcrumb menu container. Optional.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/components/Menu/MenuBreadcrumb.tsx
// import React from 'react';
// import styles from '@patternfly/react-styles/css/components/Menu/menu';
// import { css } from '@patternfly/react-styles';
// 
// export interface MenuBreadcrumbProps extends Omit<React.HTMLAttributes<HTMLDivElement>, 'ref' | 'onSelect'> {
//   /** Items within breadcrumb menu container */
//   children?: React.ReactNode;
// }
// 
// export const MenuBreadcrumb: React.FunctionComponent<MenuBreadcrumbProps> = ({
//   children,
//   ...props
// }: MenuBreadcrumbProps) => (
//   <div className={css(styles.menuBreadcrumb)} {...props}>
//     {children}
//   </div>
// );
// 
// MenuBreadcrumb.displayName = 'MenuBreadcrumb';
