package teakit_examples

import tea "github.com/charmbracelet/bubbletea"

type Example1Model struct{}

func (m *Example1Model) Init() tea.Cmd {
	return nil
}

func (m *Example1Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Example1Model) View() string {
	return "This is Example1"
}

func NewExample1Model() *Example1Model {
	return &Example1Model{}
}
