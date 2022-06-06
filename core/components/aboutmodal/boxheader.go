package aboutmodal

type BoxHeaderProps struct {
	// ClassName - additional classes added to the button. Optional.
	ClassName string
	// ProductName - Name of the Product. Optional.
	ProductName string
	// Id - id to used for Modal Box header.
	Id string
}

// contents of react-core/src/components/AboutModal/AboutModalBoxHeader.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/AboutModalBox/about-modal-box';
// import { Title } from '../Title';
// 
// export interface AboutModalBoxHeaderProps extends React.HTMLProps<HTMLDivElement> {
//   /** additional classes added to the button  */
//   className?: string;
//   /** Name of the Product  */
//   productName?: string;
//   /** id to used for Modal Box header  */
//   id: string;
// }
// 
// export const AboutModalBoxHeader: React.FunctionComponent<AboutModalBoxHeaderProps> = ({
//   className = '',
//   productName = '',
//   id,
//   ...props
// }: AboutModalBoxHeaderProps) => (
//   <div className={css(styles.aboutModalBoxHeader, className)} {...props}>
//     <Title headingLevel="h1" size="4xl" id={id}>
//       {productName}
//     </Title>
//   </div>
// );
// AboutModalBoxHeader.displayName = 'AboutModalBoxHeader';
