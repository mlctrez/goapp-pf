package codeblock

type Props struct {
	// Children - Content rendered inside the code block. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes passed to the code block wrapper. Optional.
	ClassName string
	// Actions - Actions in the code block header. Should be wrapped with CodeBlockAction. Optional.
	Actions any //  // React.ReactNode 
}

// contents of react-core/src/components/CodeBlock/CodeBlock.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/CodeBlock/code-block';
// import { css } from '@patternfly/react-styles';
// 
// export interface CodeBlockProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the code block */
//   children?: React.ReactNode;
//   /** Additional classes passed to the code block wrapper */
//   className?: string;
//   /** Actions in the code block header. Should be wrapped with CodeBlockAction. */
//   actions?: React.ReactNode;
// }
// 
// export const CodeBlock: React.FunctionComponent<CodeBlockProps> = ({
//   children = null,
//   className,
//   actions = null,
//   ...props
// }: CodeBlockProps) => (
//   <div className={css(styles.codeBlock, className)} {...props}>
//     <div className={css(styles.codeBlockHeader)}>
//       <div className={css(styles.codeBlockActions)}>{actions && actions}</div>
//     </div>
//     <div className={css(styles.codeBlockContent)}>{children}</div>
//   </div>
// );
// 
// CodeBlock.displayName = 'CodeBlock';
