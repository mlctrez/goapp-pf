package text

type ContentProps struct {
	// Children - Content rendered within the TextContent. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the TextContent. Optional.
	ClassName string
	// IsVisited - Flag to indicate the all links in a the content block have visited styles applied if the browser
	// determines the link has been visited. Optional.
	IsVisited bool
}

// contents of react-core/src/components/Text/TextContent.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Content/content';
// import { css } from '@patternfly/react-styles';
// 
// export interface TextContentProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered within the TextContent */
//   children?: React.ReactNode;
//   /** Additional classes added to the TextContent */
//   className?: string;
//   /** Flag to indicate the all links in a the content block have visited styles applied if the browser determines the link has been visited */
//   isVisited?: boolean;
// }
// 
// export const TextContent: React.FunctionComponent<TextContentProps> = ({
//   children = null,
//   className = '',
//   isVisited = false,
//   ...props
// }: TextContentProps) => (
//   <div {...props} className={css(styles.content, isVisited && styles.modifiers.visited, className)}>
//     {children}
//   </div>
// );
// TextContent.displayName = 'TextContent';
