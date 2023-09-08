package prompter

import (
	"io"
	"regexp"
	"strings"
)

type PrompterOpt func(p *Prompter)

func WithOutputTo(w io.Writer) PrompterOpt {
	return func(p *Prompter) {
		p.out = w
	}
}

// Prompt simple prompting
func Prompt(message, defaultAnswer string, opts ...PrompterOpt) string {
	p := &Prompter{
		Message: message,
		Default: defaultAnswer,
	}
	for _, o := range opts {
		o(p)
	}
	return p.Prompt()
}

// YN y/n choice
func YN(message string, defaultToYes bool, opts ...PrompterOpt) bool {
	defaultChoice := "n"
	if defaultToYes {
		defaultChoice = "y"
	}
	p := &Prompter{
		Message:    message,
		Choices:    []string{"y", "n"},
		IgnoreCase: true,
		Default:    defaultChoice,
	}
	for _, o := range opts {
		o(p)
	}
	input := p.Prompt()

	return strings.ToLower(input) == "y"
}

// YesNo yes/no choice
func YesNo(message string, defaultToYes bool, opts ...PrompterOpt) bool {
	defaultChoice := "no"
	if defaultToYes {
		defaultChoice = "yes"
	}
	p := &Prompter{
		Message:    message,
		Choices:    []string{"yes", "no"},
		IgnoreCase: true,
		Default:    defaultChoice,
	}
	for _, o := range opts {
		o(p)
	}
	input := p.Prompt()

	return strings.ToLower(input) == "yes"
}

// Password asks password
func Password(message string, opts ...PrompterOpt) string {
	p := &Prompter{
		Message: message,
		NoEcho:  true,
	}
	for _, o := range opts {
		o(p)
	}

	return p.Prompt()
}

// Choose make a choice
func Choose(message string, choices []string, defaultChoice string, opts ...PrompterOpt) string {
	p := &Prompter{
		Message: message,
		Choices: choices,
		Default: defaultChoice,
	}
	for _, o := range opts {
		o(p)
	}

	return p.Prompt()
}

// Regexp checks the answer by regexp
func Regexp(message string, reg *regexp.Regexp, defaultAnswer string, opts ...PrompterOpt) string {
	p := &Prompter{
		Message: message,
		Regexp:  reg,
		Default: defaultAnswer,
	}
	for _, o := range opts {
		o(p)
	}

	return p.Prompt()
}
