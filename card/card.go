package card

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type CardVariant string

const (
	CardVariantDefault     CardVariant = "default"
	CardVariantPrimary     CardVariant = "primary"
	CardVariantMuted       CardVariant = "muted"
	CardVariantDestructive CardVariant = "destructive"
)

type CardSize string

const (
	CardSizeSm CardSize = "sm"
	CardSizeMd CardSize = "md"
	CardSizeLg CardSize = "lg"
)

type CardState string

const (
	CardStateNormal   CardState = "normal"
	CardStateFocused  CardState = "focused"
	CardStateDisabled CardState = "disabled"
)

var (
	baseCardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2)

	baseHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			MarginBottom(1)

	baseContentStyle = lipgloss.NewStyle()

	baseFooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			MarginTop(1)
)

var cardVariantStyles = map[CardVariant]lipgloss.Style{
	CardVariantDefault:     baseCardStyle.BorderForeground(lipgloss.Color("240")),
	CardVariantMuted:       baseCardStyle.BorderForeground(lipgloss.Color("238")).Background(lipgloss.Color("234")),
	CardVariantPrimary:     baseCardStyle.BorderForeground(lipgloss.Color("34")).Background(lipgloss.Color("235")),
	CardVariantDestructive: baseCardStyle.BorderForeground(lipgloss.Color("196")).Background(lipgloss.Color("235")),
}

var cardSizeStyles = map[CardSize]func(style lipgloss.Style) lipgloss.Style{
	CardSizeSm: func(style lipgloss.Style) lipgloss.Style {
		return style.Padding(0, 1)
	},
	CardSizeMd: func(style lipgloss.Style) lipgloss.Style {
		return style.Padding(1, 2)
	},
	CardSizeLg: func(style lipgloss.Style) lipgloss.Style {
		return style.Padding(2, 4)
	},
}

var cardStateStyles = map[CardState]func(style lipgloss.Style) lipgloss.Style{
	CardStateNormal: func(style lipgloss.Style) lipgloss.Style {
		return style
	},

	CardStateFocused: func(style lipgloss.Style) lipgloss.Style {
		return style.BorderForeground(lipgloss.Color("46"))
	},

	CardStateDisabled: func(style lipgloss.Style) lipgloss.Style {
		return style.BorderForeground(lipgloss.Color("238")).
			Foreground(lipgloss.Color("243"))
	},
}

type Card struct {
	header  string
	content string
	footer  string

	width int

	variant CardVariant
	size    CardSize
	state   CardState
}

func (c Card) View() string {
	headerStyle := baseHeaderStyle
	contentStyle := baseContentStyle
	footerStyle := baseFooterStyle

	sections := []string{}

	if c.header != "" {
		sections = append(sections, headerStyle.Render(c.header))
	}

	if c.content != "" {
		sections = append(sections, contentStyle.Render(c.content))
	}

	if c.footer != "" {
		sections = append(sections, footerStyle.Render(c.footer))
	}

	style, ok := cardVariantStyles[c.variant]
	if !ok {
		style = cardVariantStyles[CardVariantDefault]
	}

	if applySize, ok := cardSizeStyles[c.size]; ok {
		style = applySize(style)
	}

	if applyState, ok := cardStateStyles[c.state]; ok {
		style = applyState(style)
	}

	if c.width > 0 {
		style = style.Width(c.width)
	}

	return style.Render(strings.Join(sections, "\n"))
}

type CardOption func(*Card)

func Header(header string) CardOption {
	return func(c *Card) {
		c.header = header
	}
}

func Footer(footer string) CardOption {
	return func(c *Card) {
		c.footer = footer
	}
}

func Width(width int) CardOption {
	return func(c *Card) {
		c.width = width
	}
}

func Size(size CardSize) CardOption {
	return func(c *Card) {
		c.size = size
	}
}

func Focused() CardOption {
	return func(c *Card) {
		c.state = CardStateFocused
	}
}

func Disabled() CardOption {
	return func(c *Card) {
		c.state = CardStateDisabled
	}
}

func newCard(card Card, options ...CardOption) Card {
	card.size = CardSizeMd
	card.state = CardStateNormal

	for _, option := range options {
		option(&card)
	}

	return card
}

func Default(content string, options ...CardOption) Card {
	card := Card{
		content: content,
		variant: CardVariantDefault,
	}

	return newCard(card, options...)
}

func Primary(content string, options ...CardOption) Card {
	card := Card{
		content: content,
		variant: CardVariantPrimary,
	}

	return newCard(card, options...)
}

func Muted(content string, options ...CardOption) Card {
	card := Card{
		content: content,
		variant: CardVariantMuted,
	}

	return newCard(card, options...)
}

func Destructive(content string, options ...CardOption) Card {
	card := Card{
		content: content,
		variant: CardVariantDestructive,
	}

	return newCard(card, options...)
}
