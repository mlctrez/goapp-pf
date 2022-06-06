package drawer

type MainProps struct {
	// ClassName - Additional classes added to the drawer main wrapper. Optional.
	ClassName string
	// Children - Content to be rendered in the drawer main wrapper. Optional.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/components/Drawer/DrawerMain.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Drawer/drawer';
// import { css } from '@patternfly/react-styles';
// 
// export interface DrawerMainProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes added to the drawer main wrapper. */
//   className?: string;
//   /** Content to be rendered in the drawer main wrapper*/
//   children?: React.ReactNode;
// }
// 
// export const DrawerMain: React.FunctionComponent<DrawerMainProps> = ({
//   // eslint-disable-next-line @typescript-eslint/no-unused-vars
//   className = '',
//   children,
//   ...props
// }: DrawerMainProps) => (
//   <div className={css(styles.drawerMain, className)} {...props}>
//     {children}
//   </div>
// );
// DrawerMain.displayName = 'DrawerMain';
