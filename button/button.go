package button

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ButtonVariant string

const (
	ButtonVariantDefault     ButtonVariant = "default"
	ButtonVariantPrimary     ButtonVariant = "primary"
	ButtonVariantSecondary   ButtonVariant = "secondary"
	ButtonVariantDestructive ButtonVariant = "destructive"
	ButtonVariantGhost       ButtonVariant = "ghost"
)

type ButtonSize string

const (
	ButtonSizeSm ButtonSize = "sm"
	ButtonSizeMd ButtonSize = "md"
	ButtonSizeLg ButtonSize = "lg"
)

type ButtonState string

const (
	ButtonStateNormal   ButtonState = "normal"
	ButtonStateFocused  ButtonState = "focused"
	ButtonStateDisabled ButtonState = "disabled"
)

var (
	baseButtonStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 2)
)

var buttonVariantStyles = map[ButtonVariant]lipgloss.Style{
	ButtonVariantDefault:     baseButtonStyle.Foreground(lipgloss.Color("250")).Background(lipgloss.Color("235")).BorderForeground(lipgloss.Color("240")),
	ButtonVariantPrimary:     baseButtonStyle.Foreground(lipgloss.Color("15")).Background(lipgloss.Color("34")).BorderForeground(lipgloss.Color("34")),
	ButtonVariantSecondary:   baseButtonStyle.Foreground(lipgloss.Color("250")).Background(lipgloss.Color("236")).BorderForeground(lipgloss.Color("238")),
	ButtonVariantDestructive: baseButtonStyle.Foreground(lipgloss.Color("15")).Background(lipgloss.Color("196")).BorderForeground(lipgloss.Color("196")),
	ButtonVariantGhost:       lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("250")),
}

var buttonSizeStyles = map[ButtonSize]func(style lipgloss.Style) lipgloss.Style{
	ButtonSizeSm: func(style lipgloss.Style) lipgloss.Style {
		return style.Padding(0, 1)
	},
	ButtonSizeMd: func(style lipgloss.Style) lipgloss.Style {
		return style.Padding(0, 2)
	},
	ButtonSizeLg: func(style lipgloss.Style) lipgloss.Style {
		return style.Padding(1, 4)
	},
}

var buttonStateStyles = map[ButtonState]func(style lipgloss.Style) lipgloss.Style{
	ButtonStateNormal: func(style lipgloss.Style) lipgloss.Style {
		return style
	},

	ButtonStateFocused: func(style lipgloss.Style) lipgloss.Style {
		return style.
			BorderForeground(lipgloss.Color("40")).
			Foreground(lipgloss.Color("15"))
	},

	ButtonStateDisabled: func(style lipgloss.Style) lipgloss.Style {
		return style.
			Foreground(lipgloss.Color("243")).
			Background(lipgloss.Color("236")).
			BorderForeground(lipgloss.Color("238"))
	},
}

type Button struct {
	label string

	width int

	variant ButtonVariant
	size    ButtonSize
	state   ButtonState

	onPress func() tea.Msg
}

func (b Button) Update(msg tea.Msg) (Button, tea.Cmd) {
	if b.state == ButtonStateDisabled {
		return b, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if b.state == ButtonStateFocused && b.onPress != nil {
				return b, func() tea.Msg {
					return b.onPress()
				}
			}
		}
	}

	return b, nil
}

func (b Button) View() string {
	style, ok := buttonVariantStyles[b.variant]
	if !ok {
		style = buttonVariantStyles[ButtonVariantDefault]
	}

	if applySize, ok := buttonSizeStyles[b.size]; ok {
		style = applySize(style)
	}

	if applyState, ok := buttonStateStyles[b.state]; ok {
		style = applyState(style)
	}

	if b.width > 0 {
		style = style.Width(b.width)
	}

	return style.Render(b.label) + "\n"
}

type ButtonOption func(*Button)

func Width(width int) ButtonOption {
	return func(b *Button) {
		b.width = width
	}
}

func OnPress(onPress func() tea.Msg) ButtonOption {
	return func(b *Button) {
		b.onPress = onPress
	}
}

func Size(size ButtonSize) ButtonOption {
	return func(b *Button) {
		b.size = size
	}
}

func Focused() ButtonOption {
	return func(b *Button) {
		b.state = ButtonStateFocused
	}
}

func Disabled() ButtonOption {
	return func(b *Button) {
		b.state = ButtonStateDisabled
	}
}

func newButton(button Button, options ...ButtonOption) Button {
	button.size = ButtonSizeMd
	button.state = ButtonStateNormal

	for _, option := range options {
		option(&button)
	}

	return button
}

func Default(label string, options ...ButtonOption) Button {
	button := Button{
		label:   label,
		variant: ButtonVariantDefault,
	}

	return newButton(button, options...)
}

func Primary(label string, options ...ButtonOption) Button {
	button := Button{
		label:   label,
		variant: ButtonVariantPrimary,
	}

	return newButton(button, options...)
}

func Secondary(label string, options ...ButtonOption) Button {
	button := Button{
		label:   label,
		variant: ButtonVariantSecondary,
	}

	return newButton(button, options...)
}

func Destructive(label string, options ...ButtonOption) Button {
	button := Button{
		label:   label,
		variant: ButtonVariantDestructive,
	}

	return newButton(button, options...)
}

func Ghost(label string, options ...ButtonOption) Button {
	button := Button{
		label:   label,
		variant: ButtonVariantGhost,
	}

	return newButton(button, options...)
}
