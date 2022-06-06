package form

type SectionProps struct {
	// Children - Content rendered inside the section. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the section. Optional.
	ClassName string
	// Title - Title for the section. Optional.
	Title any //  // React.ReactNode 
	// TitleElement - Element to wrap the section title. Optional.
	TitleElement any //  /* "div" | "h1" | "h2" | "h3" | "h4" | "h5" | "h6" */
}

// contents of react-core/src/components/Form/FormSection.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Form/form';
// import { css } from '@patternfly/react-styles';
// 
// export interface FormSectionProps extends Omit<React.HTMLProps<HTMLDivElement>, 'title'> {
//   /** Content rendered inside the section */
//   children?: React.ReactNode;
//   /** Additional classes added to the section */
//   className?: string;
//   /** Title for the section */
//   title?: React.ReactNode;
//   /** Element to wrap the section title*/
//   titleElement?: 'div' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';
// }
// 
// export const FormSection: React.FunctionComponent<FormSectionProps> = ({
//   className = '',
//   children,
//   title = '',
//   titleElement: TitleElement = 'div',
//   ...props
// }: FormSectionProps) => (
//   <section {...props} className={css(styles.formSection, className)}>
//     {title && <TitleElement className={css(styles.formSectionTitle, className)}>{title}</TitleElement>}
//     {children}
//   </section>
// );
// FormSection.displayName = 'FormSection';
