package codeblock

type CodeProps struct {
	// Children - Code rendered inside the code block. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes passed to the code block pre wrapper. Optional.
	ClassName string
	// CodeClassName - Additional classes passed to the code block code. Optional.
	CodeClassName string
}

// contents of react-core/src/components/CodeBlock/CodeBlockCode.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/CodeBlock/code-block';
// import { css } from '@patternfly/react-styles';
// 
// export interface CodeBlockCodeProps extends React.HTMLProps<HTMLPreElement> {
//   /** Code rendered inside the code block */
//   children?: React.ReactNode;
//   /** Additional classes passed to the code block pre wrapper */
//   className?: string;
//   /** Additional classes passed to the code block code */
//   codeClassName?: string;
// }
// 
// export const CodeBlockCode: React.FunctionComponent<CodeBlockCodeProps> = ({
//   children = null,
//   className,
//   codeClassName,
//   ...props
// }: CodeBlockCodeProps) => (
//   <pre className={css(styles.codeBlockPre, className)} {...props}>
//     <code className={css(styles.codeBlockCode, codeClassName)}>{children}</code>
//   </pre>
// );
// 
// CodeBlockCode.displayName = 'CodeBlockCode';
