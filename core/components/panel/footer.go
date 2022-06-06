package panel

type FooterProps struct {
	// Children - Content rendered inside the panel footer. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
}

// contents of react-core/src/components/Panel/PanelFooter.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Panel/panel';
// import { css } from '@patternfly/react-styles';
// 
// export interface PanelFooterProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the panel footer */
//   children?: React.ReactNode;
//   /** Class to add to outer div */
//   className?: string;
// }
// 
// export const PanelFooter: React.FunctionComponent<PanelFooterProps> = ({
//   className,
//   children,
//   ...props
// }: PanelFooterProps) => (
//   <div className={css(styles.panelFooter, className)} {...props}>
//     {children}
//   </div>
// );
// 
// PanelFooter.displayName = 'PanelFooter';
