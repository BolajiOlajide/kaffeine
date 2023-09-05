package main

import "github.com/charmbracelet/lipgloss"

var errorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF0000")).
	SetString("error:")

var infoStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#0869dA"))

var genericStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#808080"))

var successStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00"))
