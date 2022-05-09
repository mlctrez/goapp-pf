package avatar

type BorderType string

const (
	BorderLight BorderType = "light"
	BorderDark  BorderType = "dark"
)

func (b BorderType) check() BorderType {
	switch b {
	case BorderLight, BorderDark:
		return b
	default:
		return BorderLight
	}
}

type SizeType string

const (
	SizeSmall      SizeType = "sm"
	SizeMedium     SizeType = "md"
	SizeLarge      SizeType = "lg"
	SizeExtraLarge SizeType = "xl"
)

func Sizes() []SizeType {
	return []SizeType{SizeSmall, SizeMedium, SizeLarge, SizeExtraLarge}
}

func (s SizeType) check() SizeType {
	switch s {
	case SizeSmall, SizeMedium, SizeLarge, SizeExtraLarge:
		return s
	default:
		return SizeMedium
	}
}

func (s SizeType) name() string {
	switch s {
	case SizeSmall:
		return "Small"
	case SizeMedium:
		return "Medium"
	case SizeLarge:
		return "Large"
	case SizeExtraLarge:
		return "Extra Large"
	default:
		return ""
	}
}
