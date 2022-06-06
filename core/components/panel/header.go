package panel

type HeaderProps struct {
	// Children - Content rendered inside the panel header. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
}

// contents of react-core/src/components/Panel/PanelHeader.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Panel/panel';
// import { css } from '@patternfly/react-styles';
// 
// export interface PanelHeaderProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the panel header */
//   children?: React.ReactNode;
//   /** Class to add to outer div */
//   className?: string;
// }
// 
// export const PanelHeader: React.FunctionComponent<PanelHeaderProps> = ({
//   className,
//   children,
//   ...props
// }: PanelHeaderProps) => (
//   <div className={css(styles.panelHeader, className)} {...props}>
//     {children}
//   </div>
// );
// 
// PanelHeader.displayName = 'PanelHeader';
