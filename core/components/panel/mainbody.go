package panel

type MainBodyProps struct {
	// Children - Content rendered inside the panel main body div. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
}

// contents of react-core/src/components/Panel/PanelMainBody.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Panel/panel';
// import { css } from '@patternfly/react-styles';
// 
// export interface PanelMainBodyProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the panel main body div */
//   children?: React.ReactNode;
//   /** Class to add to outer div */
//   className?: string;
// }
// 
// export const PanelMainBody: React.FunctionComponent<PanelMainBodyProps> = ({
//   className,
//   children,
//   ...props
// }: PanelMainBodyProps) => (
//   <div className={css(styles.panelMainBody, className)} {...props}>
//     {children}
//   </div>
// );
// 
// PanelMainBody.displayName = 'PanelMainBody';
