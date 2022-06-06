package text

type ListProps struct {
	// Children - Content rendered within the TextList. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the TextList. Optional.
	ClassName string
	// Component - The text list component. Optional.
	Component any //  /* "ul" | "ol" | "dl" */
}

// contents of react-core/src/components/Text/TextList.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// 
// export enum TextListVariants {
//   ul = 'ul',
//   ol = 'ol',
//   dl = 'dl'
// }
// 
// export interface TextListProps extends React.HTMLProps<HTMLElement> {
//   /** Content rendered within the TextList */
//   children?: React.ReactNode;
//   /** Additional classes added to the TextList */
//   className?: string;
//   /** The text list component */
//   component?: 'ul' | 'ol' | 'dl';
// }
// 
// export const TextList: React.FunctionComponent<TextListProps> = ({
//   children = null,
//   className = '',
//   component = TextListVariants.ul,
//   ...props
// }: TextListProps) => {
//   const Component: any = component;
// 
//   return (
//     <Component {...props} data-pf-content className={css(className)}>
//       {children}
//     </Component>
//   );
// };
// TextList.displayName = 'TextList';
