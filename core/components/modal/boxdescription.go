package modal

type BoxDescriptionProps struct {
	// Children - Content rendered inside the description. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the description. Optional.
	ClassName string
	// Id - ID of the description. Optional.
	Id string
}

// contents of react-core/src/components/Modal/ModalBoxDescription.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/ModalBox/modal-box';
// 
// export interface ModalBoxDescriptionProps {
//   /** Content rendered inside the description */
//   children?: React.ReactNode;
//   /** Additional classes added to the description */
//   className?: string;
//   /** ID of the description */
//   id?: string;
// }
// 
// export const ModalBoxDescription: React.FunctionComponent<ModalBoxDescriptionProps> = ({
//   children = null,
//   className = '',
//   id = '',
//   ...props
// }: ModalBoxDescriptionProps) => (
//   <div {...props} id={id} className={css(styles.modalBoxDescription, className)}>
//     {children}
//   </div>
// );
// ModalBoxDescription.displayName = 'ModalBoxDescription';
