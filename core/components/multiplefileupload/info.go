package multiplefileupload

type InfoProps struct {
	// Children - Content rendered inside multiple file upload info. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
}

// contents of react-core/src/components/MultipleFileUpload/MultipleFileUploadInfo.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/MultipleFileUpload/multiple-file-upload';
// import { css } from '@patternfly/react-styles';
// 
// export interface MultipleFileUploadInfoProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside multiple file upload info */
//   children?: React.ReactNode;
//   /** Class to add to outer div */
//   className?: string;
// }
// 
// export const MultipleFileUploadInfo: React.FunctionComponent<MultipleFileUploadInfoProps> = ({
//   className,
//   children,
//   ...props
// }: MultipleFileUploadInfoProps) => (
//   <div className={css(styles.multipleFileUploadInfo, className)} {...props}>
//     {children}
//   </div>
// );
// 
// MultipleFileUploadInfo.displayName = 'MultipleFileUploadInfo';
