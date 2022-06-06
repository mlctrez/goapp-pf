package form

type FiledGroupHeaderTitleTextObject struct {
	// Text - Title text.
	Text any //  // React.ReactNode 
	// Id - The applied to the title div for accessibility.
	Id string
}

type FieldGroupHeaderTitleTextObject struct {
}

type FieldGroupHeaderProps struct {
	// ClassName - Additional classes added to the section. Optional.
	ClassName string
	// TitleText - Field group header title text. Optional.
	TitleText any //  // FormFieldGroupHeaderTitleTextObject 
	// TitleDescription - Field group header title description. Optional.
	TitleDescription any //  // React.ReactNode 
	// Actions - Field group header actions. Optional.
	Actions any //  // React.ReactNode 
}

// contents of react-core/src/components/Form/FormFieldGroupHeader.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Form/form';
// import { css } from '@patternfly/react-styles';
// 
// // typo - rename to FormFieldGroupHeaderTitleTextObject during breaking change release
// export interface FormFiledGroupHeaderTitleTextObject {
//   /** Title text. */
//   text: React.ReactNode;
//   /** The applied to the title div for accessibility */
//   id: string;
// }
// 
// export interface FormFieldGroupHeaderTitleTextObject extends FormFiledGroupHeaderTitleTextObject {}
// 
// export interface FormFieldGroupHeaderProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes added to the section */
//   className?: string;
//   /** Field group header title text */
//   titleText?: FormFieldGroupHeaderTitleTextObject;
//   /** Field group header title description */
//   titleDescription?: React.ReactNode;
//   /** Field group header actions */
//   actions?: React.ReactNode;
// }
// 
// export const FormFieldGroupHeader: React.FunctionComponent<FormFieldGroupHeaderProps> = ({
//   className,
//   titleText,
//   titleDescription,
//   actions,
//   ...props
// }: FormFieldGroupHeaderProps) => (
//   <div className={css(styles.formFieldGroupHeader, className)} {...props}>
//     <div className={css(styles.formFieldGroupHeaderMain)}>
//       {titleText && (
//         <div className={css(styles.formFieldGroupHeaderTitle)}>
//           <div className={css(styles.formFieldGroupHeaderTitleText)} id={titleText.id}>
//             {titleText.text}
//           </div>
//         </div>
//       )}
//       {titleDescription && <div className={css(styles.formFieldGroupHeaderDescription)}>{titleDescription}</div>}
//     </div>
//     <div className={css(styles.formFieldGroupHeaderActions)}>{actions && actions}</div>
//   </div>
// );
// FormFieldGroupHeader.displayName = 'FormFieldGroupHeader';
