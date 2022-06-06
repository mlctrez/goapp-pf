package aboutmodal

type BoxContentProps struct {
	// Children - content rendered inside the AboutModalBoxContent.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the AboutModalBoxContent. Optional.
	ClassName string
	// Id - id to use for About Modal Box aria described by.
	Id string
	// Trademark - The Trademark info for the product.
	Trademark string
	// NoAboutModalBoxContentContainer - Prevents the about modal from rendering content inside a container;
	// allows for more flexible layouts. Optional.
	NoAboutModalBoxContentContainer bool
}

// contents of react-core/src/components/AboutModal/AboutModalBoxContent.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/AboutModalBox/about-modal-box';
// import contentStyles from '@patternfly/react-styles/css/components/Content/content';
// 
// export interface AboutModalBoxContentProps extends React.HTMLProps<HTMLDivElement> {
//   /** content rendered inside the AboutModalBoxContent  */
//   children: React.ReactNode;
//   /** additional classes added to the AboutModalBoxContent  */
//   className?: string;
//   /** id to use for About Modal Box aria described by  */
//   id: string;
//   /** The Trademark info for the product  */
//   trademark: string;
//   /** Prevents the about modal from rendering content inside a container; allows for more flexible layouts */
//   noAboutModalBoxContentContainer?: boolean;
// }
// 
// export const AboutModalBoxContent: React.FunctionComponent<AboutModalBoxContentProps> = ({
//   children,
//   className = '',
//   trademark,
//   id,
//   noAboutModalBoxContentContainer = false,
//   ...props
// }: AboutModalBoxContentProps) => (
//   <div className={css(styles.aboutModalBoxContent, className)} id={id} {...props}>
//     <div className={css('pf-c-about-modal-box__body')}>
//       {noAboutModalBoxContentContainer ? children : <div className={css(contentStyles.content)}>{children}</div>}
//     </div>
//     <p className={css(styles.aboutModalBoxStrapline)}>{trademark}</p>
//   </div>
// );
// AboutModalBoxContent.displayName = 'AboutModalBoxContent';
