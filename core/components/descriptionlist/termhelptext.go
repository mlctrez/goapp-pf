package descriptionlist

type TermHelpTextProps struct {
	// Children - Anything that can be rendered inside of list term.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the DescriptionListTermHelpText. Optional.
	ClassName string
}

// contents of react-core/src/components/DescriptionList/DescriptionListTermHelpText.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/DescriptionList/description-list';
// import { css } from '@patternfly/react-styles';
// 
// export interface DescriptionListTermHelpTextProps extends React.HTMLProps<HTMLElement> {
//   /** Anything that can be rendered inside of list term */
//   children: React.ReactNode;
//   /** Additional classes added to the DescriptionListTermHelpText */
//   className?: string;
// }
// 
// export const DescriptionListTermHelpText: React.FunctionComponent<DescriptionListTermHelpTextProps> = ({
//   children,
//   className,
//   ...props
// }: DescriptionListTermHelpTextProps) => (
//   <dt className={css(styles.descriptionListTerm, className)} {...props}>
//     {children}
//   </dt>
// );
// DescriptionListTermHelpText.displayName = 'DescriptionListTermHelpText';
