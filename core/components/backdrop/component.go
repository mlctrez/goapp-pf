package backdrop

type Props struct {
	// Children - content rendered inside the backdrop. Optional.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the button. Optional.
	ClassName string
}

// contents of react-core/src/components/Backdrop/Backdrop.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Backdrop/backdrop';
// 
// export interface BackdropProps extends React.HTMLProps<HTMLDivElement> {
//   /** content rendered inside the backdrop */
//   children?: React.ReactNode;
//   /** additional classes added to the button */
//   className?: string;
// }
// 
// export const Backdrop: React.FunctionComponent<BackdropProps> = ({
//   children = null,
//   className = '',
//   ...props
// }: BackdropProps) => (
//   <div {...props} className={css(styles.backdrop, className)}>
//     {children}
//   </div>
// );
// Backdrop.displayName = 'Backdrop';
