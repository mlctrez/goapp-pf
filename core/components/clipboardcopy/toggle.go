package clipboardcopy

type ToggleProps struct {
	// OnClick - 
	OnClick func(event any // React.MouseEvent )
	// Id - 
	Id string
	// TextId - 
	TextId string
	// ContentId - 
	ContentId string
	// IsExpanded - Optional.
	IsExpanded bool
	// ClassName - Optional.
	ClassName string
}

// contents of react-core/src/components/ClipboardCopy/ClipboardCopyToggle.tsx
// import * as React from 'react';
// import AngleRightIcon from '@patternfly/react-icons/dist/esm/icons/angle-right-icon';
// import AngleDownIcon from '@patternfly/react-icons/dist/esm/icons/angle-down-icon';
// import { Button } from '../Button';
// 
// export interface ClipboardCopyToggleProps
//   extends Omit<React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>, 'ref'> {
//   onClick: (event: React.MouseEvent) => void;
//   id: string;
//   textId: string;
//   contentId: string;
//   isExpanded?: boolean;
//   className?: string;
// }
// 
// export const ClipboardCopyToggle: React.FunctionComponent<ClipboardCopyToggleProps> = ({
//   onClick,
//   id,
//   textId,
//   contentId,
//   isExpanded = false,
//   ...props
// }: ClipboardCopyToggleProps) => (
//   <Button
//     type="button"
//     variant="control"
//     onClick={onClick}
//     id={id}
//     aria-labelledby={`${id} ${textId}`}
//     aria-controls={`${id} ${contentId}`}
//     aria-expanded={isExpanded}
//     {...props}
//   >
//     {isExpanded ? <AngleDownIcon aria-hidden="true" /> : <AngleRightIcon aria-hidden="true" />}
//   </Button>
// );
// ClipboardCopyToggle.displayName = 'ClipboardCopyToggle';
