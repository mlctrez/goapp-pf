package popover

// from file "react-core/src/components/Popover/PopoverHeader.tsx"

type HeaderProps struct {
	// Children - Content of the popover header.
	Children any // React.ReactNode 
	// Icon - Indicates the header contains an icon. Optional.
	Icon any // React.ReactNode 
	// ClassName - Class to be applied to the header. Optional.
	ClassName string
	// TitleHeadingLevel - Heading level of the header title. Optional.
	TitleHeadingLevel any /* "h1" | "h2" | "h3" | "h4" | "h5" | "h6" */
	// AlertSeverityVariant - Severity variants for an alert popover. This modifies the color of the header to
	// match the severity. Optional.
	AlertSeverityVariant any /* "default" | "info" | "warning" | "success" | "danger" */
	// Id - Id of the header. Optional.
	Id string
	// AlertSeverityScreenReaderText - Text announced by screen reader when alert severity variant is set to
	// indicate severity level. Optional.
	AlertSeverityScreenReaderText string
}
