package clipboardcopy

type ActionProps struct {
	// Children - Content rendered inside the clipboard copy action. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the clipboard copy action. Optional.
	ClassName string
}

// contents of react-core/src/components/ClipboardCopy/ClipboardCopyAction.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/ClipboardCopy/clipboard-copy';
// import { css } from '@patternfly/react-styles';
// 
// export interface ClipboardCopyActionProps extends React.HTMLProps<HTMLLIElement> {
//   /** Content rendered inside the clipboard copy action. */
//   children?: React.ReactNode;
//   /** Additional classes added to the clipboard copy action. */
//   className?: string;
// }
// 
// export const ClipboardCopyAction: React.FunctionComponent<ClipboardCopyActionProps> = ({
//   children = null,
//   className = '',
//   ...props
// }: ClipboardCopyActionProps) => (
//   <span className={css(styles.clipboardCopyActionsItem, className)} {...props}>
//     {children}
//   </span>
// );
// ClipboardCopyAction.displayName = 'ClipboardCopyAction';
